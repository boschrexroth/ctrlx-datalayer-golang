// MIT License
//
// Copyright (c) 2021-2022 Bosch Rexroth AG
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Package main
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
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
	system.Start(true)

	// fmt.Println("-----")
	// fmt.Println("Start provider")
	// provider := system.Factory().CreateProvider(*connectionProvider)
	// defer datalayer.DeleteProvider(provider)

	// node := datalayer.NewProviderNode()
	// defer datalayer.DeleteProviderNode(node)
	// go startTestNode(node)

	// fmt.Println("Register provider: ", provider.RegisterNode("godata/test", node))
	// fmt.Println("Start provider: ", provider.Start())
	// time.Sleep(2 * time.Second)

	// fmt.Println("Provider connected: ", provider.IsConnected())
	// fmt.Println("Provider done")
	// fmt.Println("-----")

	// time.Sleep(2 * time.Second)

	var client *datalayer.Client = system.Factory().CreateClient(*connectionClient)
	defer datalayer.DeleteClient(client)
	fmt.Println("-----")
	fmt.Println("Start client")

	time.Sleep(2 * time.Second)

	client.SetTimeout(datalayer.TimeoutSettingPing, 2000)
	fmt.Printf("Ping: %x\n", client.PingSync())

	subscription, result := client.CreateSubscription("mySub", datalayer.DefaultSubscriptionProperties(), onSubscribe)
	fmt.Printf("Create Subscription %x\n", result)

	fmt.Println("Subscribe", subscription.Subscribe("scheduler/admin/info/counter"))
	lastTime = time.Now()
	<-time.After(5 * time.Second)
	fmt.Println("Unsubscribe", subscription.Unsubscribe("scheduler/admin/info/counter"))
	// rc := func(res datalayer.Result, v *datalayer.Variant) {
	// 	if v == nil {
	// 		fmt.Printf("Async callback in Go (no data): %x\n", res)
	// 	} else {
	// 		fmt.Println("Async callback in Go: ", res, v.GetFloat32())
	// 	}
	// }
	// fmt.Println("PingAsync: ", client.PingAsync(rc))

	r, v := client.BrowseSync("")
	defer datalayer.DeleteVariant(v)
	fmt.Println("Browse: ", r, v.GetArrayString())

	// v.SetFloat32(0.1234)
	// fmt.Println("Write: ", client.WriteSync("godata/test", v))
	// v.SetFloat32(1.0)
	// fmt.Println("Read: ", client.ReadSync("/framework/metrics/system/cpu-utilisation-percent", v), v.GetFloat32())

	// v.SetFloat32(0.2468)
	// fmt.Println("Async Write: ", client.WriteAsync("godata/test", v, rc))

	// fmt.Println("Read: ", client.ReadSync("/framework/metrics/system/cpu-utilisation-percent", v), v.GetFloat32())
	// fmt.Println("Async Read: ", client.ReadAsync("g/framework/metrics/system/cpu-utilisation-percent", v, rc))

	// v = datalayer.NewVariant()
	// defer datalayer.DeleteVariant(v)
	// v.SetString("Hello World")
	// fmt.Println("Test string: ", v.GetString())
	// v.SetArrayFloat64([]float64{2, 4})
	// fmt.Println("Test float64 array: ", v.GetArrayFloat64())
	// v.SetArrayBool8([]bool{true, false})
	// fmt.Println("Test bool array: ", v.GetArrayBool8())

	// w := datalayer.NewVariant()
	// defer datalayer.DeleteVariant(w)
	// w.SetArrayString([]string{"Hallo", "Welt"})
	// fmt.Println("Test string array: ", w.GetArrayString())

	//lastTime = time.Now()
	//for i:= 0; i < 30; i++ {
	//	client.ReadSync("/framework/metrics/system/cpu-utilisation-percent", v)
	//	fmt.Println("Read: ", time.Now().Sub(lastTime), v.GetFloat32())
	//	lastTime = time.Now()
	//	<-time.After(500*time.Millisecond)
	//}

	fmt.Println("Client done")
	fmt.Println("-----")

	fmt.Println("Wait...")
	<-sigs
}

func onSubscribe(result datalayer.Result, items map[string]datalayer.Variant) {
	fmt.Println("Notify ", time.Now().Sub(lastTime))
	for k, v := range items {
		fmt.Println(k, " ", v.GetUint32())
	}
	lastTime = time.Now()
}

var lastTime time.Time

func onNotify(result datalayer.Result, notifyItems []datalayer.NotifyItem) {
	fmt.Println("Notify ", time.Now().Sub(lastTime), result, notifyItems[0].Data.GetUint32())
	lastTime = time.Now()
}
