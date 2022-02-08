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

package datalayer

//#include <stdbool.h>
//#include <stdlib.h>
//#include <provider_node.h>
import "C"
import "unsafe"

// providerNodeChannelSize size of channels
const providerNodeChannelSize = 5

// ProviderNodeCallback function
type ProviderNodeCallback = func(result Result, data *Variant)

// ProviderNodeEvent event
type ProviderNodeEvent struct {
	Address  string
	Callback ProviderNodeCallback
}

// ProviderNodeEventData event
type ProviderNodeEventData struct {
	Address  string
	Data     *Variant
	Callback ProviderNodeCallback
}

// ProviderNodeChannels sets the struct.
type ProviderNodeChannels struct {
	OnCreate   chan ProviderNodeEventData
	OnRemove   chan ProviderNodeEvent
	OnBrowse   chan ProviderNodeEvent
	OnRead     chan ProviderNodeEventData
	OnWrite    chan ProviderNodeEventData
	OnMetadata chan ProviderNodeEvent
	Done       chan bool
}

// ProviderNode interface for providing data to the system.
type ProviderNode struct {
	this     C.DLR_PROVIDER_NODE
	userdata unsafe.Pointer
	channels ProviderNodeChannels
}

// Channels get all channels.
// It returns all channels.
func (n *ProviderNode) Channels() *ProviderNodeChannels {
	return &n.channels
}

// NewProviderNode initializes the provider node.
func NewProviderNode() *ProviderNode {
	channels := ProviderNodeChannels{
		OnCreate:   make(chan ProviderNodeEventData, providerNodeChannelSize),
		OnRemove:   make(chan ProviderNodeEvent, providerNodeChannelSize),
		OnBrowse:   make(chan ProviderNodeEvent, providerNodeChannelSize),
		OnRead:     make(chan ProviderNodeEventData, providerNodeChannelSize),
		OnWrite:    make(chan ProviderNodeEventData, providerNodeChannelSize),
		OnMetadata: make(chan ProviderNodeEvent, providerNodeChannelSize),
		Done:       make(chan bool),
	}
	userdata := getNodeUserdata(channels)
	ptr := C.DLR_providerNodeCreate(getNodeCallbacksC(userdata))
	return &ProviderNode{
		this:     ptr,
		userdata: userdata,
		channels: channels,
	}
}

// DeleteProviderNode destructs the provider node.
func DeleteProviderNode(n *ProviderNode) {
	if n == nil {
		return
	}
	if n.this == nil {
		return
	}
	close(n.channels.Done)
	var i int = *(*int)(n.userdata)
	nodeUnregister(i)
	C.DLR_providerNodeDelete(n.this)
	n.this = nil
}
