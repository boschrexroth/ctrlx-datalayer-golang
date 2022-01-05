// MIT License
//
// Copyright (c) 2021 Bosch Rexroth AG
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

//#include <stdbool.h>
//#include <stdlib.h>
//#include <system.h>
import "C"
import (
	"unsafe"
)

// System class
type System struct {
	this C.DLR_SYSTEM
}

// NewSystem creates a datalayer system.
// Parameter ipcPath is a path for interprocess communication - use null pointer for automatic detection.
// It returns a datalayer system.
func NewSystem(ipcPath string) *System {
	cpath := C.CString(ipcPath)
	defer C.free(unsafe.Pointer(cpath))
	ptr := C.DLR_systemCreate(cpath)
	return &System{this: ptr}
}

// DeleteSystem removes a datalayer system.
// Parameter d is a system.
func DeleteSystem(d *System) {
	if d == nil {
		return
	}
	d.Stop(true)
	C.DLR_systemDelete(d.this)
}

// Start runs a dalayer system.
// Parameter boStartBroker uses true to start a local broker. Use false to connect to an existing datalayer system e.g. the integrated datalayer system
// within the Automationcore.
func (d *System) Start(boStartBroker bool) {
	C.DLR_systemStart(d.this, C.bool(boStartBroker))
}

// Stop breaks a dalayer system.
// Parameter boForceProviderStop forces stop off all created providers for this datalayer system.
// It returns false if there is a client or provider active.
func (d *System) Stop(boForceProviderStop bool) {
	C.DLR_systemStop(d.this, C.bool(boForceProviderStop))
}

// Factory system.
// It returns the factory to create clients and provider for the datalayer.
func (d *System) Factory() *Factory {
	ptr := C.DLR_systemFactory(d.this)
	return &Factory{this: ptr, system: d}
}

// JSONConverter system.
// It returns a converter between JSON and a Variant.
func (d *System) JSONConverter() *Converter {
	ptr := C.DLR_systemJsonConverter(d.this)
	return &Converter{this: ptr}
}

// SetBfbsPath sets the base path to bfbs files.
// Parameter path is a base path to bfbs files.S
func (d *System) SetBfbsPath(path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.DLR_systemSetBfbsPath(d.this, cpath)
}
