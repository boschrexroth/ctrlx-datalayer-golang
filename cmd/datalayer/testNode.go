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
	"fmt"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
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
			req := event.Data.GetFloat32()
			newData.SetFloat32(req + 2.0)
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
