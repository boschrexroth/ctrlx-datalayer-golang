package main

import (
	"flag"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	"github.com/loov/hrtime/hrplot"
	"github.com/loov/hrtime/hrtesting"
)

var nodes = []string{
	"gobenchmark/int32",
	"scheduler/admin/properties/priority",
	"datalayer/curvemq/publickey",
	"framework/bundles/com_boschrexroth_comm_datalayer/location",
}

type BenchmarkFunc func(b *testing.B)

var benchmarks = []BenchmarkFunc{
	BenchmarkReadSync,
	BenchmarkReadASync,
}
var address string
var lastValue unsafe.Pointer
var lastSize uint64

var (
	connectionClient   = flag.String("client", "tcp://boschrexroth:boschrexroth@192.168.1.1:2069", "connection string for client")
	connectionProvider = flag.String("provider", "tcp://boschrexroth:boschrexroth@192.168.1.1:2070", "connection string for provider")

	client *datalayer.Client
)

func main() {
	testing.Init()

	flag.Parse()
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	system.Start(false)
	client = system.Factory().CreateClient(*connectionClient)
	defer datalayer.DeleteClient(client)

	provider := NewBechmarkProvider()
	if err := provider.Connect(system, *connectionProvider); err != nil {
		panic(err)
	}
	defer provider.Close()

	fmt.Printf("Connection: %s\n", *connectionClient)
	for _, a := range nodes {
		address = a
		fmt.Printf("---\nAddress: %s\n", address)
		for _, b := range benchmarks {
			name := runtime.FuncForPC(reflect.ValueOf(b).Pointer()).Name()
			fmt.Printf("Benchmark %s: %s\n", name, testing.Benchmark(b))
		}
	}
}

func BenchmarkReadSync(b *testing.B) {
	bench := hrtesting.NewBenchmark(b)
	defer bench.Report()
	defer hrplot.All(fmt.Sprintf("BenchmarkReadSync_%s.svg", strings.ReplaceAll(address, "/", "_")), bench)
	v := datalayer.NewVariant()
	for bench.Next() {
		res := client.ReadSync(address, v)
		if res != 0 {
			fmt.Printf("Failed to read: %x\n", res)
			b.Fatal()
		}
		lastValue = v.GetData()
		lastSize = v.GetSize()
	}
	datalayer.DeleteVariant(v)
}

func BenchmarkReadASync(b *testing.B) {
	bench := hrtesting.NewBenchmark(b)
	defer bench.Report()
	defer hrplot.All(fmt.Sprintf("BenchmarkReadASync_%s.svg", strings.ReplaceAll(address, "/", "_")), bench)

	done := make(chan struct{})
	rc := func(res datalayer.Result, v *datalayer.Variant) {
		if res != 0 {
			b.Fatal("Failed async")
		}
		lastValue = v.GetData()
		lastSize = v.GetSize()
		done <- struct{}{}
	}
	for bench.Next() {
		v := datalayer.NewVariant()
		res := client.ReadAsync(address, v, rc)
		if res != 0 {
			b.Fatal("Failed to read")
		}
		datalayer.DeleteVariant(v)
		<-done
	}
}
