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

package datalayer

/*
#include <stdbool.h>
#include <stdlib.h>
#include <client.h>
*/
import "C"
import (
	"sort"
	"sync"
	"unsafe"
)

// Subscription holds a subscription.
type Subscription struct {
	id        string
	client    C.DLR_CLIENT
	notifyKey C.ulonglong
	mx        sync.Mutex
	addresses map[string]struct{}
}

// Id of the subscription.
// It returns the Id.
func (s *Subscription) Id() string {
	return s.id
}

// Addresses are all currently subscribed addresses.
// It returns all subscribed adresses.
func (s *Subscription) Addresses() []string {
	addresses := make([]string, 0, len(s.addresses))
	s.mx.Lock()
	for address := range s.addresses {
		addresses = append(addresses, address)
	}
	s.mx.Unlock()
	sort.Strings(addresses)
	return addresses
}

// Subscribe registers to all passed addresses.
// Parameter addresses is a string of addresses.
// It returns the result.
func (s *Subscription) Subscribe(addresses ...string) Result {
	cId := C.CString(s.id)
	size := C.size_t(unsafe.Sizeof((*C.char)(nil)))
	length := C.size_t(len(addresses))
	cAddresses := (**C.char)(C.malloc(length * size))
	cAddressView := (*[1 << 30]*C.char)(unsafe.Pointer(cAddresses))[0:len(addresses):len(addresses)]
	defer func() {
		C.free(unsafe.Pointer(cId))
	}()
	for i, address := range addresses {
		cAddressView[i] = C.CString(address)
	}
	result := C.DLR_clientSubscribeMultiSync(s.client, cId, cAddresses, C.uint(len(addresses)))
	if result == 0 {
		s.mx.Lock()
		for _, address := range addresses {
			s.addresses[address] = struct{}{}
		}
		s.mx.Unlock()
	}
	return Result(result)
}

// Unsubscribe unregisters all passed addresses.
// Parameter addresses is a string of addresses.
// It returns the result.
func (s *Subscription) Unsubscribe(addresses ...string) Result {
	cId := C.CString(s.id)
	size := C.size_t(unsafe.Sizeof((*C.char)(nil)))
	length := C.size_t(len(addresses))
	cAddresses := (**C.char)(C.malloc(length * size))
	cAddressView := (*[1 << 30]*C.char)(unsafe.Pointer(cAddresses))[0:len(addresses):len(addresses)]
	defer func() {
		C.free(unsafe.Pointer(cId))
	}()
	for i, address := range addresses {
		cAddressView[i] = C.CString(address)
	}
	result := C.DLR_clientUnsubscribeMultiSync(s.client, cId, cAddresses, C.uint(len(addresses)))
	if result == 0 {
		s.mx.Lock()
		for _, address := range addresses {
			delete(s.addresses, address)
		}
		s.mx.Unlock()
	}
	return Result(result)
}
