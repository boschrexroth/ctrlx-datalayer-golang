package datalayer

//#include <stdbool.h>
//#include <stdlib.h>
//#include "provider.h"
import "C"
import "unsafe"

// Provider class
type Provider struct {
	this C.DLR_PROVIDER
}

// DeleteProvider destructor
func DeleteProvider(p *Provider) {
	p.Stop()
	C.DLR_providerDelete(p.this)
}

// RegisterType provider
func (p *Provider) RegisterType(address string, pathname string) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	cpath := C.CString(pathname)
	defer C.free(unsafe.Pointer(cpath))
	return Result(C.DLR_providerRegisterType(p.this, caddress, cpath))
}

// UnregisterType provider
func (p *Provider) UnregisterType(address string) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_providerUnregisterType(p.this, caddress))
}

// RegisterNode provider
func (p *Provider) RegisterNode(address string, node *ProviderNode) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_providerRegisterNode(p.this, caddress, node.this))
}

// UnregisterNode provider
func (p *Provider) UnregisterNode(address string) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_providerUnregisterNode(p.this, caddress))
}

// SetTimeoutNode provider
func (p *Provider) SetTimeoutNode(node *ProviderNode, timeoutMS uint) Result {
	return Result(C.DLR_providerSetTimeoutNode(p.this, node.this, C.uint(timeoutMS)))
}

// Start provider
func (p *Provider) Start() Result {
	return Result(C.DLR_providerStart(p.this))
}

// Stop provider
func (p *Provider) Stop() Result {
	return Result(C.DLR_providerStop(p.this))
}

// IsConnected provider
func (p *Provider) IsConnected() bool {
	return bool(C.DLR_providerIsConnected(p.this))
}

// GetToken provider
func (p *Provider) GetToken() *Variant {
	ptr := C.DLR_providerGetToken(p.this)
	return &Variant{this: ptr}
}
