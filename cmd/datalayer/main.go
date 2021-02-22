package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
)

var (
	connectionClient   = flag.String("client", "tcp://boschrexroth:boschrexroth@192.168.1.1:2069", "connection string for client")
	connectionProvider = flag.String("provider", "tcp://boschrexroth:boschrexroth@192.168.1.1:2070", "connection string for provider")
)

func main() {
	flag.Parse()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Datalayer Test Go")
	fmt.Printf("Using ctrlX CORE as client: %s\n", *connectionClient)
	fmt.Printf("Using ctrlX CORE as provider: %s\n", *connectionProvider)

	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	system.Start(false)

	fmt.Println("-----")
	fmt.Println("Start provider")
	provider := system.Factory().CreateProvider(*connectionProvider)
	defer datalayer.DeleteProvider(provider)

	node := datalayer.NewProviderNode()
	defer datalayer.DeleteProviderNode(node)
	go startTestNode(node)

	fmt.Println("Register provider: ", provider.RegisterNode("godata/test", node))
	fmt.Println("Start provider: ", provider.Start())
	time.Sleep(2 * time.Second)

	fmt.Println("Provider connected: ", provider.IsConnected())
	fmt.Println("Provider done")
	fmt.Println("-----")

	time.Sleep(2 * time.Second)

	var client *datalayer.Client = system.Factory().CreateClient(*connectionClient)
	defer datalayer.DeleteClient(client)
	fmt.Println("-----")
	fmt.Println("Start client")

	time.Sleep(2 * time.Second)

	client.SetTimeout(datalayer.TimeoutSettingPing, 2000)
	fmt.Printf("Ping: %x\n", client.PingSync())

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		if v == nil {
			fmt.Printf("Async callback in Go (no data): %x\n", res)
		} else {
			fmt.Println("Async callback in Go: ", res, v.GetFloat32())
		}
	}
	fmt.Println("PingAsync: ", client.PingAsync(rc))

	v := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v)

	fmt.Println("Browse: ", client.BrowseSync("", v), v.GetArrayString())

	v.SetFloat32(0.1234)
	fmt.Println("Write: ", client.WriteSync("godata/test", v))
	fmt.Println("Read: ", client.ReadSync("godata/test", v), v.GetFloat32())

	v.SetFloat32(0.2468)
	fmt.Println("Async Write: ", client.WriteAsync("godata/test", v, rc))

	fmt.Println("Read: ", client.ReadSync("godata/test", v), v.GetFloat32())
	fmt.Println("Async Read: ", client.ReadAsync("godata/test", v, rc))

	v = datalayer.NewVariant()
	defer datalayer.DeleteVariant(v)
	v.SetString("Hello World")
	fmt.Println("Test string: ", v.GetString())
	v.SetArrayFloat64([]float64{2, 4})
	fmt.Println("Test float64 array: ", v.GetArrayFloat64())
	v.SetArrayBool8([]bool{true, false})
	fmt.Println("Test bool array: ", v.GetArrayBool8())

	w := datalayer.NewVariant()
	defer datalayer.DeleteVariant(w)
	w.SetArrayString([]string{"Hallo", "Welt"})
	fmt.Println("Test string array: ", w.GetArrayString())

	fmt.Println("Client done")
	fmt.Println("-----")

	fmt.Println("Wait...")

	<-sigs
}
