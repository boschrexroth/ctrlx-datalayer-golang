/**
 * MIT License
 *
 * Copyright (c) 2021-2022 Bosch Rexroth AG
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package datalayer_test

import (
	"fmt"
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	a "github.com/stretchr/testify/assert"
)

var providerAddress string = ""

func init() {
	providerAddress = ctrlxAddress()
}

func TestProviderNode(t *testing.T) {
	a.NotPanics(t, func() { datalayer.DeleteProviderNode(nil) })
}

func TestDeleteProviderNode(t *testing.T) {
	if providerAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	n := datalayer.NewProviderNode()
	a.NotNil(t, n)
	a.NotPanics(t, func() { datalayer.DeleteProviderNode(n) })
}

func TestDeleteProviderNodeDeadlock(t *testing.T) {
	if providerAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	n := datalayer.NewProviderNode()
	a.NotNil(t, n)
	a.NotPanics(t, func() { datalayer.DeleteProviderNode(n) })
	a.NotPanics(t, func() { datalayer.DeleteProviderNode(n) })
}

func TestDeleteProvider(t *testing.T) {
	a.NotPanics(t, func() { datalayer.DeleteProvider(nil) })
}

func TestProviderInit(t *testing.T) {
	if providerAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)
	a.NotPanics(t, func() { system.Start(false) })

	provider := system.Factory().CreateProvider(providerAddress)
	defer datalayer.DeleteProvider(provider)
	a.NotNil(t, provider)

	a.Equal(t, datalayer.Result(0), provider.Start())
	a.True(t, provider.IsConnected())

	v := provider.GetToken()
	a.NotNil(t, v)
	defer datalayer.DeleteVariant(v)

	a.Equal(t, datalayer.Result(0), provider.Stop())
	a.False(t, provider.IsConnected())
}

func initProvider() (*datalayer.System, *datalayer.Provider) {
	system := datalayer.NewSystem("")
	system.Start(false)

	provider := system.Factory().CreateProvider(providerAddress)
	return system, provider
}

func TestProviderRegisterNode(t *testing.T) {
	if providerAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, provider := initProvider()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteProvider(provider)

	node := datalayer.NewProviderNode()
	defer datalayer.DeleteProviderNode(node)
	a.NotNil(t, node)

	a.NotPanics(t, func() { go startTestNode(node) })

	a.Equal(t, datalayer.Result(0), provider.RegisterNode("godata/test", node))
	a.Equal(t, datalayer.Result(0), provider.Start())
	a.Equal(t, true, provider.IsConnected())

	a.Equal(t, datalayer.Result(0), provider.SetTimeoutNode(node, 1000))

	a.Equal(t, datalayer.Result(0), provider.UnregisterNode("godata/test"))
}

func TestProviderRegisterType(t *testing.T) {
	if providerAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, provider := initProvider()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteProvider(provider)

	a.Equal(t, datalayer.Result(0), provider.Start())
	a.Equal(t, true, provider.IsConnected())

	a.Equal(t, datalayer.Result(0), provider.RegisterType("types/godata/test", "../fbs/sampleSchema.bfbs"))

	a.Equal(t, datalayer.Result(0), provider.UnregisterType("types/godata/test"))
}

var goFloat float32 = 0.369

func startTestNode(node *datalayer.ProviderNode) {
	for {
		select {
		case event := <-node.Channels().OnCreate:
			fmt.Println("event: oncreate")
			event.Callback(datalayer.Result(0), nil)

		case event := <-node.Channels().OnRemove:
			fmt.Println("event: OnRemove")
			event.Callback(datalayer.Result(0), nil)

		case event := <-node.Channels().OnBrowse:
			newData := datalayer.NewVariant()
			defer datalayer.DeleteVariant(newData)
			newData.SetArrayString([]string{"d1", "d2"})
			fmt.Println("event: OnBrowse")
			event.Callback(datalayer.Result(0), newData)

		case event := <-node.Channels().OnRead:
			newData := datalayer.NewVariant()
			defer datalayer.DeleteVariant(newData)
			newData.SetFloat32(goFloat)
			fmt.Println("event: OnRead")
			event.Callback(datalayer.Result(0), newData)

		case event := <-node.Channels().OnWrite:
			goFloat = event.Data.GetFloat32()
			fmt.Println("event: OnWrite")
			event.Callback(datalayer.Result(0), event.Data)

		case event := <-node.Channels().OnMetadata:
			fmt.Println("event: OnMetadata")
			event.Callback(datalayer.Result(0), nil)

		case <-node.Channels().Done:
			fmt.Println("event: Done")
			return
		}
	}
}
