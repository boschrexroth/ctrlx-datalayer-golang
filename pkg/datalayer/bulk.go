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
//#include <system.h>
//#include <bulk.h>
//extern DLR_RESULT ClientReadBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey);
//extern DLR_RESULT ClientWriteBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey);
//extern DLR_RESULT ClientCreateBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey);
//extern DLR_RESULT ClientDeleteBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey);
//extern DLR_RESULT ClientBrowseBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey);
//extern DLR_RESULT ClientMetadataBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey);
//extern unsigned long long CreateBulkRequest(uint32_t size);
//extern void DeleteBulkRequest(unsigned long long p);
//extern void SetBulkRequestAddress(unsigned long long p, uint32_t index, const char *address);
//extern void SetBulkRequestData(unsigned long long p, uint32_t index, DLR_VARIANT data);
import "C"

import (
	"time"
	"unsafe"
)

// Structure for bulk response
type Response struct {
	Address string
	Data    *Variant
	Time    time.Time
	Result  Result
}

// Structure for bulk request
type request struct {
	address *C.char
	data    *Variant
}

// Structure for bulk
// Function calls are combined into a single call
type Bulk struct {
	client       *Client
	requests     []request
	Response     []Response
	asyncCreator *asyncCreator
}

// ResponseBulkCallback function type
type ResponseBulkCallback = func([]Response)

// Struct for BulkReadArg
// Address for the request
// Argument of the request
type BulkReadArg struct {
	Address  string
	Argument *Variant
}

// Struct for BulkWriteArg
// Address for the request
// Data of the request
type BulkWriteArg struct {
	Address string
	Data    *Variant
}

// Struct for BulkCreateArg
// Address for the request
// Data of the request
type BulkCreateArg struct {
	Address string
	Data    *Variant
}

// Structure for bulkCreater helper class
type bulkCreator struct {
	bulk C.DLR_BULK
}

// Structure for asyncCreator helper class
type asyncCreator struct {
	asyncHandle C.ulonglong
}

// newBulk creates a ctrlX Data Layer bulk
// Parameter c is a Client
// It returns a ctrlX Data Layer bulk.
func newBulk(c *Client) *Bulk {
	return &Bulk{
		client:       c,
		requests:     make([]request, 0),
		Response:     make([]Response, 0),
		asyncCreator: nil,
	}
}

// DeleteBulk destructs the bulk.
// It destroys the bulk.
func DeleteBulk(b *Bulk) {
	if b == nil {
		return
	}
	b.close()
}

// close remove the internal bulk datas
func (b *Bulk) close() {
	b.closeAsync()
	b.closeResponses()
	b.closeRequests()
}

// add a bulk request
// Parameter address is an address of the node.
// Parameter data read arguments for read operation or
// write argument for write operation, 'nil' when not used
func (b *Bulk) add(address string, data *Variant) {
	req := request{address: C.CString(address), data: data}
	b.requests = append(b.requests, req)
}

// Read read values using a bulk operation. This function is synchronous.
// Parameter args is a list of read arguments
// It returns the status of function call
func (b *Bulk) Read(args ...BulkReadArg) Result {
	b.closeResponses()
	b.closeRequests()
	for _, arg := range args {
		b.add(arg.Address, arg.Argument)
	}
	result, bulkcreator := newBulkCreator(b.requests)
	defer deleteBulkCreator(bulkcreator)
	if result != ResultOk {
		return result
	}
	result = b.clientBulkReadSync(bulkcreator.bulk)
	if result != ResultOk {
		return result
	}
	result = b.createResponse(bulkcreator.bulk)
	return result
}

// ReadAsync reads values using a bulk operation. This function is asynchronous.
// Parameter cb is the callback of response
// Parameter args is a list of read arguments
// It returns the status of function call
func (b *Bulk) ReadAsync(cb ResponseBulkCallback, args ...BulkReadArg) Result {
	b.closeAsync()
	b.closeRequests()
	for _, arg := range args {
		b.add(arg.Address, arg.Argument)
	}
	b.asyncCreator = newAsyncCreator(b.requests)
	result := Result(C.ClientReadBulkASync(b.client.this, b.asyncCreator.asyncHandle, nil, responseBulkRegister(cb)))
	return result
}

// Delete deletes nodes using a bulk operation. This function is synchronous.
// Parameter addresses is a list of node addresses
// It returns the status of function call
func (b *Bulk) Delete(addresses ...string) Result {
	b.closeResponses()
	b.closeRequests()
	for _, arg := range addresses {
		b.add(arg, nil)
	}
	result, bulkcreator := newBulkCreator(b.requests)
	defer deleteBulkCreator(bulkcreator)
	if result != ResultOk {
		return result
	}
	result = b.clientBulkDeleteSync(bulkcreator.bulk)
	if result != ResultOk {
		return result
	}
	result = b.createResponse(bulkcreator.bulk)
	return result
}

// DeleteAsync deletes nodes using a bulk operation. This function is asynchronous.
// Parameter cb is the callback of response
// Parameter addresses is a list of node addresses
// It returns the status of function call
func (b *Bulk) DeleteAsync(cb ResponseBulkCallback, addresses ...string) Result {
	b.closeAsync()
	b.closeRequests()
	for _, arg := range addresses {
		b.add(arg, nil)
	}
	b.asyncCreator = newAsyncCreator(b.requests)
	result := Result(C.ClientDeleteBulkASync(b.client.this, b.asyncCreator.asyncHandle, nil, responseBulkRegister(cb)))
	return result
}

// Write writes nodes using a bulk operation. This function is synchronous.
// Parameter args is a list of write arguments
// It returns the status of function call
func (b *Bulk) Write(args ...BulkWriteArg) Result {
	b.closeResponses()
	b.closeRequests()
	for _, arg := range args {
		b.add(arg.Address, arg.Data)
	}
	result, bulkcreator := newBulkCreator(b.requests)
	defer deleteBulkCreator(bulkcreator)
	if result != ResultOk {
		return result
	}
	result = b.clientBulkWriteSync(bulkcreator.bulk)
	if result != ResultOk {
		return result
	}
	result = b.createResponse(bulkcreator.bulk)
	return result
}

// WriteAsync deletes nodes using a bulk operation. This function is asynchronous.
// Parameter cb is the callback of response
// Parameter args is a list of write arguments
// It returns the status of function call
func (b *Bulk) WriteAsync(cb ResponseBulkCallback, args ...BulkWriteArg) Result {
	b.closeAsync()
	b.closeRequests()
	for _, arg := range args {
		b.add(arg.Address, arg.Data)
	}
	b.asyncCreator = newAsyncCreator(b.requests)
	result := Result(C.ClientWriteBulkASync(b.client.this, b.asyncCreator.asyncHandle, nil, responseBulkRegister(cb)))
	return result
}

// Create creates nodes using a bulk operation. This function is synchronous.
// Parameter args is a list of create arguments
// It returns the status of function call
func (b *Bulk) Create(args ...BulkCreateArg) Result {
	b.closeResponses()
	b.closeRequests()
	for _, arg := range args {
		b.add(arg.Address, arg.Data)
	}
	result, bulkcreator := newBulkCreator(b.requests)
	defer deleteBulkCreator(bulkcreator)
	if result != ResultOk {
		return result
	}
	result = b.clientBulkCreateSync(bulkcreator.bulk)
	if result != ResultOk {
		return result
	}
	result = b.createResponse(bulkcreator.bulk)
	return result
}

// CreateAsync creates nodes using a bulk operation. This function is asynchronous.
// Parameter cb is the callback of response
// Parameter args is a list of create arguments
// It returns the status of function call
func (b *Bulk) CreateAsync(cb ResponseBulkCallback, args ...BulkCreateArg) Result {
	b.closeAsync()
	b.closeRequests()
	for _, arg := range args {
		b.add(arg.Address, arg.Data)
	}
	b.asyncCreator = newAsyncCreator(b.requests)
	result := Result(C.ClientCreateBulkASync(b.client.this, b.asyncCreator.asyncHandle, nil, responseBulkRegister(cb)))
	return result
}

// Browse browses nodes using a bulk operation. This function is synchronous.
// Parameter addresses is a list of node addresses
// It returns the status of function call
func (b *Bulk) Browse(addresses ...string) Result {
	b.closeResponses()
	b.closeRequests()
	for _, arg := range addresses {
		b.add(arg, nil)
	}
	result, bulkcreator := newBulkCreator(b.requests)
	defer deleteBulkCreator(bulkcreator)
	if result != ResultOk {
		return result
	}
	result = b.clientBulkBrowseSync(bulkcreator.bulk)
	if result != ResultOk {
		return result
	}
	result = b.createResponse(bulkcreator.bulk)
	return result
}

// BrowseAsync browses nodes using a bulk operation. This function is asynchronous.
// Parameter cb is the callback of response
// Parameter addresses is a list of node addresses
// It returns the status of function call
func (b *Bulk) BrowseAsync(cb ResponseBulkCallback, addresses ...string) Result {
	b.closeAsync()
	b.closeRequests()
	for _, arg := range addresses {
		b.add(arg, nil)
	}
	b.asyncCreator = newAsyncCreator(b.requests)
	result := Result(C.ClientBrowseBulkASync(b.client.this, b.asyncCreator.asyncHandle, nil, responseBulkRegister(cb)))
	return result
}

// Metadata reads metadata of nodes using a bulk operation. This function is synchronous.
// Parameter addresses is a list of node addresses
// It returns the status of function call
func (b *Bulk) Metadata(addresses ...string) Result {
	b.closeResponses()
	b.closeRequests()
	for _, arg := range addresses {
		b.add(arg, nil)
	}
	result, bulkcreator := newBulkCreator(b.requests)
	defer deleteBulkCreator(bulkcreator)
	if result != ResultOk {
		return result
	}
	result = b.clientBulkMetadataSync(bulkcreator.bulk)
	if result != ResultOk {
		return result
	}
	result = b.createResponse(bulkcreator.bulk)
	return result
}

// MetadataAsync reads metadata of nodes using a bulk operation. This function is asynchronous.
// Parameter cb is the callback of response
// Parameter addresses is a list of node addresses
// It returns the status of function call
func (b *Bulk) MetadataAsync(cb ResponseBulkCallback, addresses ...string) Result {
	b.closeAsync()
	b.closeRequests()
	for _, arg := range addresses {
		b.add(arg, nil)
	}
	b.asyncCreator = newAsyncCreator(b.requests)
	result := Result(C.ClientMetadataBulkASync(b.client.this, b.asyncCreator.asyncHandle, nil, responseBulkRegister(cb)))
	return result
}

// closeRequests the list of request
func (b *Bulk) closeRequests() {
	for i := 0; i < len(b.requests); i++ {
		C.free(unsafe.Pointer(b.requests[i].address))
	}
	b.requests = make([]request, 0)
}

// closeResponses the list of response
func (b *Bulk) closeResponses() {
	for i := 0; i < len(b.Response); i++ {
		DeleteVariant(b.Response[i].Data)
	}
	b.Response = make([]Response, 0)
}

// closeAsync the list of response
func (b *Bulk) closeAsync() {
	if b.asyncCreator == nil {
		return
	}
	deleteAsyncCreator(b.asyncCreator)
	b.asyncCreator = nil
}

// createResponse a list of response
func (b *Bulk) createResponse(bulk C.DLR_BULK) Result {
	for i := 0; i < len(b.requests); i++ {
		resp := Response{}
		resp.Address = b.bulkGetResponseAddress(bulk, i)
		resp.Result = b.bulkGetResponseResult(bulk, i)
		resp.Data = b.bulkGetResponseData(bulk, i)
		resp.Time = b.bulkGetResponseTimestamp(bulk, i)
		b.Response = append(b.Response, resp)
	}
	return ResultOk
}

// clientBulkReadSync is a wrapper function for C.DLR_clientBulkReadSync
func (b *Bulk) clientBulkReadSync(bulk C.DLR_BULK) Result {
	return Result(C.DLR_clientBulkReadSync(b.client.this, bulk, nil))
}

// clientBulkDeleteSync is a wrapper function for C.DLR_clientBulkDeleteSync
func (b *Bulk) clientBulkDeleteSync(bulk C.DLR_BULK) Result {
	return Result(C.DLR_clientBulkDeleteSync(b.client.this, bulk, nil))
}

// clientBulkWriteSync is a wrapper function for C.DLR_clientBulkWriteSync
func (b *Bulk) clientBulkWriteSync(bulk C.DLR_BULK) Result {
	return Result(C.DLR_clientBulkWriteSync(b.client.this, bulk, nil))
}

// clientBulkCreateSync is a wrapper function for C.DLR_clientBulkCreateSync
func (b *Bulk) clientBulkCreateSync(bulk C.DLR_BULK) Result {
	return Result(C.DLR_clientBulkCreateSync(b.client.this, bulk, nil))
}

// clientBulkBrowseSync is a wrapper function for C.DLR_clientBulkBrowseSync
func (b *Bulk) clientBulkBrowseSync(bulk C.DLR_BULK) Result {
	return Result(C.DLR_clientBulkBrowseSync(b.client.this, bulk, nil))
}

// clientBulkMetadataSync is a wrapper function for C.DLR_clientBulkMetadataSync
func (b *Bulk) clientBulkMetadataSync(bulk C.DLR_BULK) Result {
	return Result(C.DLR_clientBulkMetadataSync(b.client.this, bulk, nil))
}

// bulkGetResponseAddress is a wrapper function for C.DLR_bulkGetResponseAddress
func (b *Bulk) bulkGetResponseAddress(bulk C.DLR_BULK, index int) string {
	p := C.DLR_bulkGetResponseAddress(bulk, C.size_t(index))
	s := C.GoString(p)
	return s
}

// bulkGetResponseResult is a wrapper function for C.DLR_bulkGetResponseResult
func (b *Bulk) bulkGetResponseResult(bulk C.DLR_BULK, index int) Result {
	return Result(C.DLR_bulkGetResponseResult(bulk, C.size_t(index)))
}

// bulkGetResponseData is a wrapper function for C.DLR_bulkGetResponseData
func (b *Bulk) bulkGetResponseData(bulk C.DLR_BULK, index int) *Variant {
	val := C.DLR_bulkGetResponseData(bulk, C.size_t(index))
	v := &Variant{this: val} //only reference
	data := NewVariant()
	v.Copy(data)
	return data
}

// bulkGetResponseTimestamp is a wrapper function for C.DLR_bulkGetResponseTimestamp
func (b *Bulk) bulkGetResponseTimestamp(bulk C.DLR_BULK, index int) time.Time {
	time := C.DLR_bulkGetResponseTimestamp(bulk, C.size_t(index))
	return getTime(int64(time))
}

// newBulkCreator create a bulk helper class
// Parameter requestes is a list of request
func newBulkCreator(requests []request) (Result, *bulkCreator) {
	c := &bulkCreator{bulk: nil}
	result := c.create(requests)
	return result, c
}

// deleteBulkCreator destructs the bulk.
// It destroys the bulk.
func deleteBulkCreator(c *bulkCreator) {
	c.bulkDelete()
}

// create a list of bulk request
// Parameter requestes is a list of request
func (c *bulkCreator) create(requests []request) Result {
	c.bulk = C.DLR_bulkCreate(C.size_t(len(requests)))
	for i, req := range requests {
		result := c.bulksetRequestAddress(i, req.address)
		if result != ResultOk {
			return result
		}
		if req.data != nil {
			result = c.bulksetRequestData(i, req.data)
			if result != ResultOk {
				return result
			}
		}
	}
	return ResultOk
}

// bulkCreate wrapper 'C' function
func (c *bulkCreator) bulkCreate(size int) {
	c.bulk = C.DLR_bulkCreate(C.size_t(size))
}

// bulkDelete wrapper 'C' function
func (c *bulkCreator) bulkDelete() {
	C.DLR_bulkDelete(c.bulk)
}

// bulksetRequestAddress wrapper 'C' function
func (c *bulkCreator) bulksetRequestAddress(index int, address *C.char) Result {
	return Result(C.DLR_bulkSetRequestAddress(c.bulk, C.size_t(index), address))
}

// bulksetRequestData wrapper 'C' function
func (c *bulkCreator) bulksetRequestData(index int, data *Variant) Result {
	return Result(C.DLR_bulkSetRequestData(c.bulk, C.size_t(index), data.this))
}

// newAsyncCreator create a bulk helper class
// Parameter requestes is a list of request
func newAsyncCreator(requests []request) *asyncCreator {
	a := &asyncCreator{asyncHandle: 0}
	a.create(requests)
	return a
}

// deleteAsyncCreator destructs the bulk.
// It destroys the bulk.
func deleteAsyncCreator(a *asyncCreator) {
	C.DeleteBulkRequest(a.asyncHandle)
	a.asyncHandle = 0
}

// create a list of bulk request
// Parameter requestes is a list of request
func (a *asyncCreator) create(requests []request) {
	a.asyncHandle = C.CreateBulkRequest(C.uint32_t(len(requests)))
	for i, req := range requests {
		C.SetBulkRequestAddress(a.asyncHandle, C.uint32_t(i), req.address)
		if req.data != nil {
			C.SetBulkRequestData(a.asyncHandle, C.uint32_t(i), req.data.this)
		}
	}
}
