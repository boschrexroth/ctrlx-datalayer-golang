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

//#include <stdbool.h>
//#include <stdlib.h>
//#include <provider.h>
import "C"
import "unsafe"

// Provider interface to manage provider nodes.
type Provider struct {
	this C.DLR_PROVIDER
}

// DeleteProvider deletes the provider instance.
func DeleteProvider(p *Provider) {
	if p == nil {
		return
	}
	if p.this == nil {
		return
	}
	p.Stop()
	C.DLR_providerDelete(p.this)
	p.this = nil
}

// RegisterType registers a type to the ctrlX Data Layer.
// Parameter address is an address of the node to register (no wildcards allowed).
// Parameter pathname is a path to flatbuffer bfbs.
// It returns the status of function call.
func (p *Provider) RegisterType(address string, pathname string) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	cpath := C.CString(pathname)
	defer C.free(unsafe.Pointer(cpath))
	return Result(C.DLR_providerRegisterType(p.this, caddress, cpath))
}

// RegisterTypeVariant registers a type to the ctrlX Data Layer.
// Parameter address is an address of the node to register (no wildcards allowed).
// Parameter data is a variant which contains the flatbuffer bfbs.
// It returns the status of function call.
func (p *Provider) RegisterTypeVariant(address string, data *Variant) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_providerRegisterTypeVariant(p.this, caddress, data.this))
}

// UnregisterType unregisters a node to the ctrlX Data Layer.
// Parameter address is an address of the node to unregister (wildcards allowed).
// It returns the status of function call.
func (p *Provider) UnregisterType(address string) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_providerUnregisterType(p.this, caddress))
}

// RegisterNode registers a node to the ctrlX Data Layer.
// Parameter address is an address of the node to register (wildcards allowed).
// Parameter node ia a node to register.
// It returns the status of function call.
func (p *Provider) RegisterNode(address string, node *ProviderNode) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_providerRegisterNode(p.this, caddress, node.this))
}

// UnregisterNode unregisters a node from the ctrlX Data Layer.
// Parameter address is an address of the node to unregister (wildcards allowed).
// Parameter node ia a node to register.
// It returns the status of function call.
func (p *Provider) UnregisterNode(address string) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_providerUnregisterNode(p.this, caddress))
}

// SetTimeoutNode sets timeout for a node for asynchron requests (default value is 1000ms).
// Parameter node is a node to set timeout.
// Parameter timeoutMS is a timeout in milliseconds for this node.
// It returns the status of function call.
func (p *Provider) SetTimeoutNode(node *ProviderNode, timeoutMS uint) Result {
	return Result(C.DLR_providerSetTimeoutNode(p.this, node.this, C.uint(timeoutMS)))
}

// Start starts the provider.
// It returns the status of function call.
func (p *Provider) Start() Result {
	return Result(C.DLR_providerStart(p.this))
}

// Stop stops the provider.
// It returns the status of function call.
func (p *Provider) Stop() Result {
	return Result(C.DLR_providerStop(p.this))
}

// IsConnected returns whether provider is connected or not.
// It returns the status of the connection.
func (p *Provider) IsConnected() bool {
	return bool(C.DLR_providerIsConnected(p.this))
}

// GetToken returns the current token of the current request.You can call this function during your onRead, onWrite, ... methods of your ProviderNodes. If there is no current request the method return an empty token.
// It returns the current token.
func (p *Provider) GetToken() *Variant {
	ptr := C.DLR_providerGetToken(p.this)
	return &Variant{this: ptr}
}
