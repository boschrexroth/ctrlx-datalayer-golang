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
#cgo linux pkg-config: libsystemd libzmq ctrlx-datalayer
#cgo CFLAGS: -I./headers -I/usr/include/comm/datalayer/c
#cgo linux LDFLAGS: -lcomm_datalayer -lzmq
#cgo windows,amd64 LDFLAGS: -lcomm_datalayer
#include <stdbool.h>
#include <stdlib.h>
#include "client.h"
#include "wrappers.h"
*/
import "C"
import (
	"errors"
	"unsafe"

	// import of c-headers
	_ "github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer/headers"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/pkg/fbs/comm/datalayer"
)

// TimeoutSetting gets the settings of the different timeout values.
type TimeoutSetting C.enum_DLR_TIMEOUT_SETTING

// TimeoutSetting enum definition
const (
	TimeoutSettingIdle = C.DLR_TIMEOUT_SETTING_IDLE
	TimeoutSettingPing = C.DLR_TIMEOUT_SETTING_PING
)

// ResponseCallback function type
type ResponseCallback = func(result Result, v *Variant)

type notifyResponseCallback = func(result Result, notifyItems []NotifyItem)

// Client is an interface for the accessing of data from the system.
type Client struct {
	this      C.DLR_CLIENT
	converter *subscriptionPropertiesConverter
}

// DeleteClient destructs the client.
// It destroys the client.
func DeleteClient(c *Client) {
	if c == nil {
		return
	}
	if c.this == nil {
		return
	}
	C.DLR_clientDelete(c.this)
	c.this = nil
}

// PingAsync pings the next hop. This function is asynchronous. It will return immediately. Callback will be called if function call is finished.
// Parameter callback is a callback to call when function is finished.
// Parameter userdata will be returned in callback as a user data. You can use this userdata to identify your request.
// It returns the status of hte function call.
func (c *Client) PingAsync(onResponse ResponseCallback) Result {

	return Result(C.ClientPingASync(c.this, responseRegister(onResponse)))
}

// CreateAsync creates an object. This function is asynchronous. It will return immediately. Callback will be called if function call is finished. Result data may be provided in callback function.
// Parameter address is an address of the node to create object in.
// Parameter data is a data of the object.
// Parameter callback is a callback to call when function is finished.
// Parameter userdata will be returned in callback as a user data. You can use this userdata to identify your request.
// It returns the status of function call.
func (c *Client) CreateAsync(address string, data *Variant, onResponse ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.ClientCreateASync(c.this, caddress, data.this, nil, responseRegister(onResponse)))
}

// RemoveAsync removes an object. This function is asynchronous. It will return immediately. Callback will be called if function call is finished. Result data may be provided in callback function.
// Parameters address is an address of the node to remove.
// Parameters callback is a callback to call when function is finished.
// Parameters userdata will be returned in callback as a user data. You can use this userdata to identify your request.
// It returns the status of function call.
func (c *Client) RemoveAsync(address string, onResponse ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.ClientRemoveASync(c.this, caddress, nil, responseRegister(onResponse)))
}

// BrowseAsync searches an object. This function is asynchronous. It will return immediately. Callback will be called if function call is finished. Result data may be provided in callback function.
// Parameter address is an address of the node to browse.
// Parameter callback is a callback to call when function is finished.
// Parameter userdata will be returned in callback as a user data. You can use this userdata to identify your request.
// It returns the status of function call.
func (c *Client) BrowseAsync(address string, onResponse ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.ClientBrowseASync(c.this, caddress, nil, responseRegister(onResponse)))
}

// ReadAsync reads an object. This function is asynchronous. It will return immediately. Callback will be called if function call is finished. Result data may be provided in callback function.
// Parameter address is an address of the node to read.
// Parameter callback is a callback to call when function is finished.
// Parameter userdata will be returned in callback as a user data. You can use this userdata to identify your request.
// It returns the status of function call.
func (c *Client) ReadAsync(address string, data *Variant, onResponse ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.ClientReadASync(c.this, caddress, data.this, nil, responseRegister(onResponse)))
}

// WriteAsync writes an object. This function is synchronous: It will wait for the answer.
// Parameter address is an address of the node to write.
// Parameter data is a data of the object.
// Parameter callback is a callback to call when function is finished.
// Parameter userdata will be returned in callback as a user data. You can use this userdata to identify your request.
// It returns the status of function call.
func (c *Client) WriteAsync(address string, data *Variant, onResponse ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.ClientWriteASync(c.this, caddress, data.this, nil, responseRegister(onResponse)))
}

// CreateSubscription creates a subscription as configured in the subscriptionProperties, which returns value to the onSubscription callback.
// It returns the status of function call.
func (c *Client) CreateSubscription(id string, subscriptionProperties SubscriptionProperties, onSubscription OnSubscription) (*Subscription, Result) {
	data := NewVariant()
	defer DeleteVariant(data)
	subscriptionProperties.Id = id
	jsonBytes, err := subscriptionProperties.BuildJson()
	if err != nil {
		return nil, C.DL_INVALID_VALUE
	}
	propBytes, err := c.converter.JsonToFlatBuffer(jsonBytes)
	if err != nil {
		return nil, C.DL_INVALID_VALUE
	}
	data.SetFlatbuffers(propBytes)
	notifyKey := notifyResponseRegister(func(result Result, notifyItems []NotifyItem) {
		items := make(map[string]Variant)
		for _, notifyItem := range notifyItems {
			info := fbs.GetRootAsNotifyInfo(notifyItem.Info.GetFlatbuffers(), 0)
			items[string(info.Node())] = notifyItem.Data
		}
		onSubscription(result, items)
	})
	r := C.ClientCreateSubscriptionSync(c.this, data.this, notifyKey, nil)
	if r != 0 {
		return nil, Result(r)
	}
	return &Subscription{
		id:        id,
		client:    c.this,
		notifyKey: notifyKey,
		addresses: make(map[string]struct{}),
	}, Result(r)
}

type OnSubscription func(result Result, items map[string]Variant)

// DeleteSubscription deletes a subscription.
func (c *Client) DeleteSubscription(subscription *Subscription) {
	cId := C.CString(subscription.id)
	defer C.free(unsafe.Pointer(cId))
	notifyResponseUnregister(subscription.notifyKey)
	C.DLR_clientUnsubscribeAllSync(c.this, cId)
}

// MetadataAsync reads metadata of an object. This function is asynchronous. It will return immediately. Callback will be called if function call is finished. Result data may be provided in callback function.
// Parameter address is an address of the node to read metadata.
// Paremeter callback is a callback to call when function is finished.
// Parmeter userdata will be returned in callback as a user data. You can use this userdata to identify your request.
// It returns the status of function call.
func (c *Client) MetadataAsync(address string, onResponse ResponseCallback) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.ClientMetadataASync(c.this, caddress, nil, responseRegister(onResponse)))
}

// PingSync pings the next hop. This function is synchronous.
// It returns the status of function call.
func (c *Client) PingSync() Result {
	return Result(C.DLR_clientPingSync(c.this))
}

// CreateSync creates an object. This function is synchronous.
// Parameter address is an address of the node to create object in.
// Parameter variant is a data of the object.
// It returns the status of function call or a variant result of write or a tuple (Result, Variant).
func (c *Client) CreateSync(address string, data *Variant) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientCreateSync(c.this, caddress, data.this, nil))
}

// RemoveSync removes an object. This function is synchronous.
// Parameter address is an address of the node to remove.
// It returns the status of function call.
func (c *Client) RemoveSync(address string) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientRemoveSync(c.this, caddress, nil))
}

// BrowseSync searches an object. This function is synchronous.
// Parameter address is an address of the node to browse.
// It returns the status of function call or a tuple (Result, Variant).
// It returns the children of the node. Data will be provided as Variant array of strings.
func (c *Client) BrowseSync(address string) (Result, *Variant) {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	data := NewVariant()
	r := Result(C.DLR_clientBrowseSync(c.this, caddress, data.this, nil))
	return r, data
}

// Browse returns all subnodes of address
func (c *Client) Browse(address string) (Result, []string) {
	r, d := c.BrowseSync(address)
	defer DeleteVariant(d)
	if r != Result(0) {
		return r, []string{}
	}
	v := make([]string, len(d.GetArrayString()))
	copy(v, d.GetArrayString())
	return r, v
}

// ReadSync reads an object. This function is synchronous.
// Paramater address is an address of the node to read.
// Parameter args reads arguments data of the node.
// It returns the status of function call or a data of the node or a tuple (Result, Variant).
func (c *Client) ReadSyncArgs(address string, args *Variant) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientReadSync(c.this, caddress, args.this, nil))
}

// ReadSync reads an object. This function is synchronous.
// Parameter address is an address of the node to read.
// It returns the status of function call or a data of the node or a tuple (Result, Variant).
func (c *Client) ReadSync(address string) (Result, *Variant) {
	data := NewVariant()
	r := c.ReadSyncArgs(address, data)
	return r, data
}

// WriteSync writes an object. This function is synchronous.
// Parameter address an address of the node to write.
// Parameter variant ia a new data of the node.
// It returns the status of function call or a result of write or a tuple Result, Variant).
func (c *Client) WriteSync(address string, data *Variant) Result {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return Result(C.DLR_clientWriteSync(c.this, caddress, data.this, nil))
}

// MetadataSync reads a metadata of an object. This function is synchronous.
// Parameter address is an address of the node to read metadata of.
// It returns the status of function call or tuple (Result, Variant) or a metadata of the node. Data will be provided as Variant flatbuffers with metadata.fbs data type.
func (c *Client) MetadataSync(address string) (Result, *Variant) {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	data := NewVariant()
	r := Result(C.DLR_clientMetadataSync(c.this, caddress, data.this, nil))
	return r, data
}

// Metadata reads metadata of an object. This function is synchronous.
// Parameter address is an address of the node to read metadata of.
// It returns the status of function call or a tuple (Result, Variant) or a metadata of the node. Data will be provided as Variant flatbuffers with metadata.fbs data type.
func (c *Client) Metadata(address string) (Result, *fbs.Metadata) {
	r, v := c.MetadataSync(address)
	defer DeleteVariant(v)
	if r != Result(0) {
		return r, nil
	}
	f := make([]byte, len(v.GetFlatbuffers()))
	copy(f, v.GetFlatbuffers())
	m := fbs.GetRootAsMetadata(f, 0)
	return r, m
}

// SetTimeout sets a client timeout value.
// Parameter timeout is a timeout to set.
// Parameter value is a value to set.
// It returns the status of function call.
func (c *Client) SetTimeout(timeout TimeoutSetting, value uint) Result {
	return Result(C.DLR_clientSetTimeout(c.this, C.DLR_TIMEOUT_SETTING(timeout), C.uint(value)))
}

// IsConnected returns whether provider is connected.
// It returns the status of connection as <bool>.
func (c *Client) IsConnected() bool {
	return bool(C.DLR_clientIsConnected(c.this))
}

// SetAuthToken sets persistent security access token for authentication as JWT payload.
// Parameter token is a security access &token for authentication.
func (c *Client) SetAuthToken(token string) {
	ctoken := C.CString(token)
	defer C.free(unsafe.Pointer(ctoken))
	C.DLR_clientSetAuthToken(c.this, ctoken)
}

// GetAuthToken returns persistent security access token for authentication.
// It returns the security access token for authentication as <string>.
func (c *Client) GetAuthToken() string {
	p := C.DLR_clientGetAuthToken(c.this)
	s := C.GoString(p)
	return s
}

func (c *Client) readJsonSync(cv *Converter, address string, indentStep int, data *Variant) (Result, *Variant) {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	r := Result(C.DLR_clientReadJsonSync(c.this, cv.this, caddress, data.this, C.int(indentStep), nil))
	return r, data
}

// This function reads a values as a JSON string.
// Parameter converter is reference to the converter (see System json_converter()).
// Parameter address is an address of the node to read.
// Parameter indentStep is an indentation length for json string.
// It returns the status of function and generated JSON
func (c *Client) ReadJsonSync(cv *Converter, address string, indentStep int) (Result, []byte) {
	return c.ReadJsonSyncArgs(cv, address, indentStep, nil)
}

// This function reads a values as a JSON string.
// Parameter converter is reference to the converter (see System json_converter()).
// Parameter address is an address of the node to read.
// Parameter indentStep is an indentation length for json string.
// Parameter json generated JSON as array of bytes (string)
// It returns the status of function and generated JSON
func (c *Client) ReadJsonSyncArgs(cv *Converter, address string, indentStep int, data []byte) (Result, []byte) {
	d := NewVariant()
	defer DeleteVariant(d)
	if data != nil {
		d.SetString(string(data))
	}
	r, d := c.readJsonSync(cv, address, indentStep, d)

	if r != ResultOk {
		return r, nil
	}
	return r, []byte(d.GetString())
}

// This function writes a JSON value.
// Parameter converter is a reference to the converter (see System json_converter()).
// Parameter address is an address of the node to write.
// Parameter json is a JSON value to write.
// It returns the status of function and error Error object.
func (c *Client) WriteJsonSync(cv *Converter, address string, json []byte) (Result, error) {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	cjson := C.CString(string(json))
	defer C.free(unsafe.Pointer(cjson))
	err := NewVariant()
	defer DeleteVariant(err)
	r := Result(C.DLR_clientWriteJsonSync(c.this, cv.this, caddress, cjson, err.this, nil))
	if r != Result(0) {
		return r, errors.New(err.GetString())
	}
	return r, nil
}

// CreateBulk creates a ctrlX Data Layer bulk
func (c *Client) CreateBulk() *Bulk {
	return newBulk(c)
}
