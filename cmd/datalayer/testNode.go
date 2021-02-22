package main

import (
	"fmt"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
)

var goFloat float32 = 0.369

func startTestNode(node *datalayer.ProviderNode) {
	fmt.Println("Test Node started")
	for {
		select {
		case event := <-node.Channels().OnCreate:
			fmt.Println("OnCreate event: ", event.Address, event.Data)
			goFloat = 0.369
			event.Callback(datalayer.Result(0), nil)
		case event := <-node.Channels().OnRemove:
			fmt.Println("OnRemove event: ", event.Address)
			goFloat = 0
			event.Callback(datalayer.Result(0), nil)
		case event := <-node.Channels().OnBrowse:
			fmt.Println("OnBrowse event: ", event.Address)
			newData := datalayer.NewVariant()
			defer datalayer.DeleteVariant(newData)
			newData.SetArrayString([]string{})
			event.Callback(datalayer.Result(0), newData)
		case event := <-node.Channels().OnRead:
			fmt.Println("OnRead event: ", event.Address, event.Data)
			newData := datalayer.NewVariant()
			defer datalayer.DeleteVariant(newData)
			newData.SetFloat32(goFloat)
			event.Callback(datalayer.Result(0), newData)
		case event := <-node.Channels().OnWrite:
			fmt.Println("OnWrite event: ", event.Address, event.Data)
			goFloat = event.Data.GetFloat32()
			event.Callback(datalayer.Result(0), event.Data)
		case event := <-node.Channels().OnMetadata:
			fmt.Println("OnMetadata event: ", event.Address)
			event.Callback(datalayer.Result(1), nil)
		case <-node.Channels().Done:
			fmt.Println("Test Node stopped")
			return
		}
	}
}
