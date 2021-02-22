package datalayer

//#include <stdbool.h>
//#include <stdlib.h>
//#include "system.h"
import "C"
import (
	"unsafe"
)

// System class
type System struct {
	this C.DLR_SYSTEM
}

// NewSystem constructor
func NewSystem(ipcPath string) *System {
	cpath := C.CString(ipcPath)
	defer C.free(unsafe.Pointer(cpath))
	ptr := C.DLR_systemCreate(cpath)
	return &System{this: ptr}
}

// DeleteSystem destructor
func DeleteSystem(d *System) {
	C.DLR_systemDelete(d.this)
}

// Start system
func (d *System) Start(boStartBroker bool) {
	C.DLR_systemStart(d.this, false)
}

// Stop system
func (d *System) Stop(boForceProviderStop bool) {
	C.DLR_systemStop(d.this, C.bool(boForceProviderStop))
}

// Factory system
func (d *System) Factory() *Factory {
	ptr := C.DLR_systemFactory(d.this)
	return &Factory{this: ptr}
}

// JSONConverter system
func (d *System) JSONConverter() *Converter {
	ptr := C.DLR_systemJsonConverter(d.this)
	return &Converter{this: ptr}
}

// SetBfbsPath system
func (d *System) SetBfbsPath(path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.DLR_systemSetBfbsPath(d.this, cpath)
}
