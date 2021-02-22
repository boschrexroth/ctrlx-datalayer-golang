package datalayer

/*
#include <stdio.h>
#include <stdbool.h>

#include "variant.h"

extern void responseCallbackC(DLR_RESULT result, DLR_VARIANT data, void* userdata);
*/
import "C"

import (
	"sync"
	"unsafe"
)

// responseUserData type
type responseUserData struct {
	f ResponseCallback
}

//export responseCallbackGo
func responseCallbackGo(result Result, cdata C.DLR_VARIANT, cuserdata unsafe.Pointer) {
	var i int = *(*int)(cuserdata)
	var userdata *responseUserData = responseLookup(i)
	if cdata == nil {
		userdata.f(result, nil)
	} else {
		userdata.f(result, &Variant{this: cdata})
	}
	responseUnregister(i)
}

// getResponseCallbackC get callback function in C
func getResponseCallbackC() *[0]byte {
	return (*[0]byte)(C.responseCallbackC)
}

// getResponseUserdata get userdata for C callback
func getResponseUserdata(f ResponseCallback) unsafe.Pointer {
	userdata := &responseUserData{f: f}
	i := responseRegister(userdata)
	return unsafe.Pointer(&i)
}

// --------------------------------------------------------------------------------------------------------
// see https://github.com/golang/go/wiki/cgo
var responseMu sync.Mutex
var responseIndex int
var responseFns = make(map[int]*responseUserData)

func responseRegister(userdata *responseUserData) int {
	responseMu.Lock()
	defer responseMu.Unlock()
	responseIndex++
	for responseFns[responseIndex] != nil {
		responseIndex++
	}
	responseFns[responseIndex] = userdata
	return responseIndex
}

func responseLookup(i int) *responseUserData {
	responseMu.Lock()
	defer responseMu.Unlock()
	return responseFns[i]
}

func responseUnregister(i int) {
	responseMu.Lock()
	defer responseMu.Unlock()
	delete(responseFns, i)
}
