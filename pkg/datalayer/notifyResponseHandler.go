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
#include <stdio.h>
#include <stdbool.h>
#include <system.h>

*/
import "C"
import (
	"sync"
	"unsafe"
)

//export notifyResponseCallbackGo
func notifyResponseCallbackGo(result Result, cNotifyItems unsafe.Pointer, count uint32, responseKey C.ulonglong) {
	onNotify := notifyResponseLookup(responseKey)
	if onNotify == nil {
		return
	}
	if cNotifyItems == nil {
		onNotify(result, nil)
	} else {
		nis := (*[1 << 30]C.NotifyItem)(cNotifyItems)[:count:count]
		notifyItems := make([]NotifyItem, 0, count)
		for _, ni := range nis {
			notifyItems = append(notifyItems, NotifyItem{
				Data: Variant{this: ni.data},
				Info: Variant{this: ni.info},
			})
		}
		onNotify(result, notifyItems)
	}
}

// --------------------------------------------------------------------------------------------------------
// see https://github.com/golang/go/wiki/cgo
var notifyResponseMu sync.Mutex
var notifyResponseIndex C.ulonglong
var notifyResponseFns = make(map[C.ulonglong]notifyResponseCallback)

func notifyResponseRegister(onNotify notifyResponseCallback) C.ulonglong {
	notifyResponseMu.Lock()
	defer notifyResponseMu.Unlock()
	notifyResponseIndex++
	for notifyResponseFns[notifyResponseIndex] != nil {
		notifyResponseIndex++
	}
	notifyResponseFns[notifyResponseIndex] = onNotify
	return notifyResponseIndex
}

func notifyResponseLookup(i C.ulonglong) notifyResponseCallback {
	notifyResponseMu.Lock()
	defer notifyResponseMu.Unlock()
	return notifyResponseFns[i]
}

func notifyResponseUnregister(i C.ulonglong) {
	notifyResponseMu.Lock()
	defer notifyResponseMu.Unlock()
	delete(notifyResponseFns, i)
}
