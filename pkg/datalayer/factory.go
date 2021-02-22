package datalayer

//#include <stdbool.h>
//#include <stdlib.h>
//#include "factory.h"
import "C"
import "unsafe"

// Factory class
type Factory struct {
	this C.DLR_FACTORY
}

// CreateClient factory
func (f *Factory) CreateClient(remote string) *Client {
	cremote := C.CString(remote)
	defer C.free(unsafe.Pointer(cremote))
	ptr := C.DLR_factoryCreateClient(f.this, cremote)
	return &Client{this: ptr}
}

// CreateProvider factory
func (f *Factory) CreateProvider(remote string) *Provider {
	cremote := C.CString(remote)
	defer C.free(unsafe.Pointer(cremote))
	ptr := C.DLR_factoryCreateProvider(f.this, cremote)
	return &Provider{this: ptr}
}
