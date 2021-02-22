package datalayer

//#include <stdbool.h>
//#include <stdlib.h>
//#include "provider_node.h"
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

// ProviderNodeChannels struct
type ProviderNodeChannels struct {
	OnCreate   chan ProviderNodeEventData
	OnRemove   chan ProviderNodeEvent
	OnBrowse   chan ProviderNodeEvent
	OnRead     chan ProviderNodeEventData
	OnWrite    chan ProviderNodeEventData
	OnMetadata chan ProviderNodeEvent
	Done       chan bool
}

// ProviderNode class
type ProviderNode struct {
	this     C.DLR_PROVIDER_NODE
	userdata unsafe.Pointer
	channels ProviderNodeChannels
}

// Channels providernode
func (n *ProviderNode) Channels() *ProviderNodeChannels {
	return &n.channels
}

// NewProviderNode constructor
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

// DeleteProviderNode destructor
func DeleteProviderNode(n *ProviderNode) {
	n.channels.Done <- true
	var i int = *(*int)(n.userdata)
	nodeUnregister(i)
	C.DLR_providerNodeDelete(n.this)
}
