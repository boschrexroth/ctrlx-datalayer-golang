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

#include <client.h>

extern void responseCallbackC(DLR_RESULT result, DLR_VARIANT data, void* userdata);
*/
import "C"

import (
	"sync"
	"unsafe"
)

//export responseBulkCallbackGo
func responseBulkCallbackGo(responses unsafe.Pointer, count uint32, responseKey C.ulonglong) {
	onResponse := responseBulkLookup(responseKey)
	if onResponse == nil {
		return
	}

	resps := make([]Response, count)
	nis := (*[1 << 30]C.DLR_BULK_RESPONSE)(responses)[:count:count]

	for i := 0; i < int(count); i++ {
		rep := Response{
			Address: C.GoString(nis[i].address),
			Data:    &Variant{this: nis[i].data},
			Time:    getTime(int64(nis[i].timestamp)),
			Result:  Result(nis[i].result)}
		resps[i] = rep
	}
	onResponse(resps)
	responseBulkUnregister(responseKey)
}

// --------------------------------------------------------------------------------------------------------
// see https://github.com/golang/go/wiki/cgo
var responseBulkMu sync.Mutex
var responseBulkIndex C.ulonglong
var responseBulkFns = make(map[C.ulonglong]ResponseBulkCallback)

func responseBulkRegister(onResponse ResponseBulkCallback) C.ulonglong {
	responseBulkMu.Lock()
	defer responseBulkMu.Unlock()
	responseBulkIndex++
	for responseBulkFns[responseBulkIndex] != nil {
		responseBulkIndex++
	}
	responseBulkFns[responseBulkIndex] = onResponse
	return responseBulkIndex
}

func responseBulkLookup(i C.ulonglong) ResponseBulkCallback {
	responseBulkMu.Lock()
	defer responseBulkMu.Unlock()
	return responseBulkFns[i]
}

func responseBulkUnregister(i C.ulonglong) {
	responseBulkMu.Lock()
	defer responseBulkMu.Unlock()
	delete(responseBulkFns, i)
}
