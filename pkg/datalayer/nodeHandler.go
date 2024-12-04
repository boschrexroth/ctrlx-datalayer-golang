// MIT License
//
// Copyright (c) 2021-2024 Bosch Rexroth AG
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

package datalayer

/*
#include <stdio.h>
#include <stdbool.h>

#include <variant.h>
#include <provider_node.h>

typedef DLR_PROVIDER_NODE_CALLBACK TYPE_CB;
typedef DLR_PROVIDER_NODE_CALLBACKDATA TYPE_CBDATA;

typedef enum NODE_ACTION {
	NODE_ACTION_ON_CREATE,
	NODE_ACTION_ON_REMOVE,
	NODE_ACTION_ON_BROWSE,
	NODE_ACTION_ON_READ,
	NODE_ACTION_ON_WRITE,
	NODE_ACTION_ON_METADATA,
	NODE_ACTION_ON_SUBSCRIPTION,
	NODE_ACTION_ON_UNSUBSCRIPTION
}NODE_ACTION;

extern void callCallbackC(TYPE_CB cb, uint64_t cbdata, DLR_RESULT result, DLR_VARIANT data);

extern void nodeCallbackOnCreate(void* userdata, char* address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata);
extern void nodeCallbackOnRemove(void* userdata, char* address, TYPE_CB cb, TYPE_CBDATA cbdata);
extern void nodeCallbackOnBrowse(void* userdata, char* address, TYPE_CB cb, TYPE_CBDATA cbdata);
extern void nodeCallbackOnRead(void* userdata, char* address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata);
extern void nodeCallbackOnWrite(void* userdata, char* address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata);
extern void nodeCallbackOnMetadata(void* userdata, char* address, TYPE_CB cb, TYPE_CBDATA cbdata);
extern void nodeCallbackSubscription(void *userdata, DLR_SUBSCRIPTION subscription, char *address);
extern void nodeCallbackUnsubscription(void *userdata, DLR_SUBSCRIPTION subscription, char *address);

*/
import "C"
import (
	"sync"
	"unsafe"
)

// nodeAction enum
type nodeAction C.enum_NODE_ACTION

// nodeAction enum definition
const (
	nodeActionOnCreate        nodeAction = C.NODE_ACTION_ON_CREATE
	nodeActionOnRemove        nodeAction = C.NODE_ACTION_ON_REMOVE
	nodeActionOnBrowse        nodeAction = C.NODE_ACTION_ON_BROWSE
	nodeActionOnRead          nodeAction = C.NODE_ACTION_ON_READ
	nodeActionOnWrite         nodeAction = C.NODE_ACTION_ON_WRITE
	nodeActionOnMetadata      nodeAction = C.NODE_ACTION_ON_METADATA
	nodeActionOnSubsciption   nodeAction = C.NODE_ACTION_ON_SUBSCRIPTION
	nodeActionOnUnsubsciption nodeAction = C.NODE_ACTION_ON_UNSUBSCRIPTION
)

// nodeCallbackData type
type nodeUserData struct {
	channels ProviderNodeChannels
}

//export nodeCallbackGo
func nodeCallbackGo(cuserdata unsafe.Pointer, caddress *C.char, cdata C.DLR_VARIANT, cb C.TYPE_CB, cbdata C.ulong, action C.int) {
	var i int = *(*int)(cuserdata)
	var userdata *nodeUserData = nodeLookup(i)
	address := C.GoString(caddress)
	switch nodeAction(action) {
	case nodeActionOnCreate:
		userdata.channels.OnCreate <- ProviderNodeEventData{Address: address, Data: &Variant{this: cdata}, Callback: createCallbackC(cb, cbdata)}
	case nodeActionOnRemove:
		userdata.channels.OnRemove <- ProviderNodeEvent{Address: address, Callback: createCallbackC(cb, cbdata)}
	case nodeActionOnBrowse:
		userdata.channels.OnBrowse <- ProviderNodeEvent{Address: address, Callback: createCallbackC(cb, cbdata)}
	case nodeActionOnRead:
		userdata.channels.OnRead <- ProviderNodeEventData{Address: address, Data: &Variant{this: cdata}, Callback: createCallbackC(cb, cbdata)}
	case nodeActionOnWrite:
		userdata.channels.OnWrite <- ProviderNodeEventData{Address: address, Data: &Variant{this: cdata}, Callback: createCallbackC(cb, cbdata)}
	case nodeActionOnMetadata:
		userdata.channels.OnMetadata <- ProviderNodeEvent{Address: address, Callback: createCallbackC(cb, cbdata)}
	default:
		panic("Unknown action type in callback")
	}
}

//export nodeCallbackSubGo
func nodeCallbackSubGo(cuserdata unsafe.Pointer, sub C.DLR_SUBSCRIPTION, caddress *C.char, action C.int) {
	var i int = *(*int)(cuserdata)
	var userdata *nodeUserData = nodeLookup(i)
	address := C.GoString(caddress)
	switch nodeAction(action) {
	case nodeActionOnSubsciption:
		userdata.channels.OnSubscribe <- ProviderNodeSubscription{Address: address, Subsciption: &ProviderSubscription{this: sub}}
	case nodeActionOnUnsubsciption:
		userdata.channels.OnUnsubscribe <- ProviderNodeSubscription{Address: address, Subsciption: &ProviderSubscription{this: sub}}
	default:
		panic("Unknown action type in callback")
	}
}

// getNodeCallbacksC get struct with subset of the C callbacks
func getNodeCallbacksC(userdata unsafe.Pointer) C.DLR_PROVIDER_NODE_CALLBACKS {
	return C.DLR_PROVIDER_NODE_CALLBACKS{
		userData:   userdata,
		onCreate:   (*[0]byte)(C.nodeCallbackOnCreate),
		onRemove:   (*[0]byte)(C.nodeCallbackOnRemove),
		onBrowse:   (*[0]byte)(C.nodeCallbackOnBrowse),
		onRead:     (*[0]byte)(C.nodeCallbackOnRead),
		onWrite:    (*[0]byte)(C.nodeCallbackOnWrite),
		onMetadata: (*[0]byte)(C.nodeCallbackOnMetadata),
	}
}

// getNodeSubCallbacksC get struct with all C callbacks
func getNodeSubCallbacksC(userdata unsafe.Pointer) C.DLR_PROVIDER_NODE_CALLBACKS {
	return C.DLR_PROVIDER_NODE_CALLBACKS{
		userData:      userdata,
		onCreate:      (*[0]byte)(C.nodeCallbackOnCreate),
		onRemove:      (*[0]byte)(C.nodeCallbackOnRemove),
		onBrowse:      (*[0]byte)(C.nodeCallbackOnBrowse),
		onRead:        (*[0]byte)(C.nodeCallbackOnRead),
		onWrite:       (*[0]byte)(C.nodeCallbackOnWrite),
		onMetadata:    (*[0]byte)(C.nodeCallbackOnMetadata),
		onSubscribe:   (*[0]byte)(C.nodeCallbackSubscription),
		onUnsubscribe: (*[0]byte)(C.nodeCallbackUnsubscription),
	}
}

// getNodeUserdata store Go channels in userdata of C callback
func getNodeUserdata(channels ProviderNodeChannels) unsafe.Pointer {
	userdata := &nodeUserData{channels: channels}
	i := nodeRegister(userdata)
	return unsafe.Pointer(&i)
}

// createCallbackC closure to hide C types
func createCallbackC(cb C.TYPE_CB, cbdata C.ulong) ProviderNodeCallback {
	return func(result Result, data *Variant) {
		if data == nil {
			C.callCallbackC(cb, cbdata, C.DLR_RESULT(result), nil)
		} else {
			C.callCallbackC(cb, cbdata, C.DLR_RESULT(result), data.this)
		}
	}
}

// --------------------------------------------------------------------------------------------------------
// see https://github.com/golang/go/wiki/cgo
var nodeMu sync.Mutex
var nodeIndex int
var nodeFns = make(map[int]*nodeUserData)

func nodeRegister(userdata *nodeUserData) int {
	nodeMu.Lock()
	defer nodeMu.Unlock()
	nodeIndex++
	for nodeFns[nodeIndex] != nil {
		nodeIndex++
	}
	nodeFns[nodeIndex] = userdata
	return nodeIndex
}

func nodeLookup(i int) *nodeUserData {
	nodeMu.Lock()
	defer nodeMu.Unlock()
	return nodeFns[i]
}

func nodeUnregister(i int) {
	nodeMu.Lock()
	defer nodeMu.Unlock()
	delete(nodeFns, i)
}
