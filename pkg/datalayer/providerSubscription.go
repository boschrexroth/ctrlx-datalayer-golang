// MIT License
//
// Copyright (c) 2021-2024 Bosch Rexroth AG
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
//extern unsigned long long CreateProviderSubscriptionsNotifyItems(size_t size);
//extern void DeleteProviderSubscriptionsNotifyItems(unsigned long long p);
//extern void SetProviderSubscriptionsNotifyItem(unsigned long long p, uint32_t index, DLR_VARIANT data, DLR_VARIANT info);
//extern DLR_NOTIFY_ITEM* GetProviderSubscriptionsNotifyItems(unsigned long long p);
import "C"
import (
	"time"

	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
)

// ProviderSubscription interface to manage subscription.
type ProviderSubscription struct {
	this C.DLR_SUBSCRIPTION
}

// GetUniqueId unambiguous identification of the subscription
func (p *ProviderSubscription) GetUniqueId() uint64 {
	return uint64(uintptr(p.this))
}

// GetTimestamp is a wrapper function for C.DLR_SubscriptionGetTimestamp
func (p *ProviderSubscription) GetTimestamp() time.Time {
	time := C.DLR_SubscriptionGetTimestamp(p.this)
	return getTime(int64(time))
}

// GetNodes gets the list of address nodes of a subscription
func (p *ProviderSubscription) GetNodes() []string {
	n := NewVariant()
	defer DeleteVariant(n)
	C.DLR_SubscriptionGetNodes(p.this, n.this)
	return n.GetArrayString()
}

// GetProps gets the properties of a subscription
func (p *ProviderSubscription) GetProps() *fbs.SubscriptionProperties {
	v := C.DLR_SubscriptionGetProps(p.this)
	vf := &Variant{this: v}
	props := fbs.GetRootAsSubscriptionProperties(vf.GetFlatbuffers(), 0)
	return props
}

// Publish publishes new data for this subscription
// Parameter status of notification. On failure subscription is canceled for all items.
// Parameter items list of the notification items.
func (p *ProviderSubscription) Publish(status Result, items []*NotifyItemPublish) Result {
	if len(items) == 0 {
		return ResultOk
	}
	handle := C.CreateProviderSubscriptionsNotifyItems(C.size_t(len(items)))
	for i, n := range items {
		C.SetProviderSubscriptionsNotifyItem(handle, C.uint32_t(i), n.GetData().this, n.GetInfo().this)
	}
	r := Result(C.DLR_SubscriptionPublish(p.this, C.DLR_RESULT(status), C.GetProviderSubscriptionsNotifyItems(handle), C.size_t(len(items))))
	C.DeleteProviderSubscriptionsNotifyItems(handle)
	return r
}
