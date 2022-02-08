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

#include <variant.h>

extern void responseCallbackC(DLR_RESULT result, DLR_VARIANT data, void* userdata);
*/
import "C"

import (
	"sync"
)

//export responseCallbackGo
func responseCallbackGo(result Result, cdata C.DLR_VARIANT, responseKey C.ulonglong) {
	onResponse := responseLookup(responseKey)
	if cdata == nil {
		onResponse(result, nil)
	} else {
		onResponse(result, &Variant{this: cdata})
	}
	responseUnregister(responseKey)
}

// --------------------------------------------------------------------------------------------------------
// see https://github.com/golang/go/wiki/cgo
var responseMu sync.Mutex
var responseIndex C.ulonglong
var responseFns = make(map[C.ulonglong]ResponseCallback)

func responseRegister(onResponse ResponseCallback) C.ulonglong {
	responseMu.Lock()
	defer responseMu.Unlock()
	responseIndex++
	for responseFns[responseIndex] != nil {
		responseIndex++
	}
	responseFns[responseIndex] = onResponse
	return responseIndex
}

func responseLookup(i C.ulonglong) ResponseCallback {
	responseMu.Lock()
	defer responseMu.Unlock()
	return responseFns[i]
}

func responseUnregister(i C.ulonglong) {
	responseMu.Lock()
	defer responseMu.Unlock()
	delete(responseFns, i)
}
