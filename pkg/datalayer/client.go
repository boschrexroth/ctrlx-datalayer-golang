package datalayer

/*
#cgo linux pkg-config: libsystemd
#cgo linux,amd64 LDFLAGS: -lcomm_datalayer -lzmq
#cgo linux,arm64 LDFLAGS: -lcomm_datalayer -lzmq -lgcrypt -lgpg-error
#cgo windows,amd64 LDFLAGS: -lcomm_datalayer
#include <stdbool.h>
#include <stdlib.h>
#include "client.h"
*/
import "C"
import "unsafe"

// TimeoutSetting enum
type TimeoutSetting C.enum_DLR_TIMEOUT_SETTING

// TimeoutSetting enum definition
const (
	TimeoutSettingIdle TimeoutSetting = C.DLR_TIMEOUT_SETTING_IDLE
	TimeoutSettingPing TimeoutSetting = C.DLR_TIMEOUT_SETTING_PING
)

// ResponseCallback function type
type ResponseCallback = func(result Result, v *Variant)

// Client class
type Client struct {
	this C.DLR_CLIENT
}

// DeleteClient destructor
func DeleteClient(c *Client) {
	C.DLR_clientDelete(c.this)
}

// PingAsync client
func (c *Client) PingAsync(f ResponseCallback) Result {
	return Result(C.DLR_clientPingASync(c.this, getResponseCallbackC(), getResponseUserdata(f)))
}

// CreateAsync client
func (c *Client) CreateAsync(address string, data *Variant, f ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientCreateASync(c.this, caddress, data.this, nil, getResponseCallbackC(), getResponseUserdata(f)))
}

// RemoveAsync client
func (c *Client) RemoveAsync(address string, f ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientRemoveASync(c.this, caddress, nil, getResponseCallbackC(), getResponseUserdata(f)))
}

// BrowseAsync client
func (c *Client) BrowseAsync(address string, f ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientBrowseASync(c.this, caddress, nil, getResponseCallbackC(), getResponseUserdata(f)))
}

// ReadAsync client
func (c *Client) ReadAsync(address string, data *Variant, f ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientReadASync(c.this, caddress, data.this, nil, getResponseCallbackC(), getResponseUserdata(f)))
}

// WriteAsync client
func (c *Client) WriteAsync(address string, data *Variant, f ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientWriteASync(c.this, caddress, data.this, nil, getResponseCallbackC(), getResponseUserdata(f)))
}

// MetadataAsync client
func (c *Client) MetadataAsync(address string, f ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientMetadataASync(c.this, caddress, nil, getResponseCallbackC(), getResponseUserdata(f)))
}

// PingSync client
func (c *Client) PingSync() Result {
	return Result(C.DLR_clientPingSync(c.this))
}

// CreateSync client
func (c *Client) CreateSync(address string, data *Variant) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientCreateSync(c.this, caddress, data.this, nil))
}

// RemoveSync client
func (c *Client) RemoveSync(address string) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientRemoveSync(c.this, caddress, nil))
}

// BrowseSync client
func (c *Client) BrowseSync(address string, data *Variant) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientBrowseSync(c.this, caddress, data.this, nil))
}

// ReadSync client
func (c *Client) ReadSync(address string, data *Variant) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientReadSync(c.this, caddress, data.this, nil))
}

// WriteSync client
func (c *Client) WriteSync(address string, data *Variant) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientWriteSync(c.this, caddress, data.this, nil))
}

// MetadataSync client
func (c *Client) MetadataSync(address string, data *Variant) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientMetadataSync(c.this, caddress, data.this, nil))
}

// SetTimeout client
func (c *Client) SetTimeout(timeout TimeoutSetting, value uint) Result {
	return Result(C.DLR_clientSetTimeout(c.this, C.DLR_TIMEOUT_SETTING(timeout), C.uint(value)))
}

// IsConnected client
func (c *Client) IsConnected() bool {
	return bool(C.DLR_clientIsConnected(c.this))
}

// SetAuthToken client
func (c *Client) SetAuthToken(token string) {
	ctoken := C.CString(token)
	defer C.free(unsafe.Pointer(ctoken))
	C.DLR_clientSetAuthToken(c.this, ctoken)
}

// GetAuthToken client
func (c *Client) GetAuthToken() string {
	p := C.DLR_clientGetAuthToken(c.this)
	s := C.GoString(p)
	defer C.free(unsafe.Pointer(p))
	return s
}
