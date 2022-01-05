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
//#include <factory.h>
import "C"
import "unsafe"

// Factory class
type Factory struct {
	this   C.DLR_FACTORY
	system *System
}

// CreateClient creates a client for accessing data of the system.
// Parameter remote is an address of the ctrlX Data Layer.
// It returns the client.
func (f *Factory) CreateClient(remote string) *Client {
	cremote := C.CString(remote)
	defer C.free(unsafe.Pointer(cremote))
	ptr := C.DLR_factoryCreateClient(f.this, cremote)
	return &Client{
		this:      ptr,
		converter: newSubscriptionPropertiesConverter(f.system),
	}
}

// CreateProvider creates a provider to provide data to the ctrlX Data Layer.
// Parameter remote is an address of the ctrlX Data Layer.
// It returns the provider.
func (f *Factory) CreateProvider(remote string) *Provider {
	cremote := C.CString(remote)
	defer C.free(unsafe.Pointer(cremote))
	ptr := C.DLR_factoryCreateProvider(f.this, cremote)
	return &Provider{this: ptr}
}
