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

import (
	"time"

	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
	flatbuffers "github.com/google/flatbuffers/go"
)

// NotifyItemPublish interface for publish subscription items.
type NotifyItemPublish struct {
	data       *Variant
	info       *Variant
	notifyInfo *fbs.NotifyInfoT
}

// NewNotifyItemPublish creates a NotifyItemPublish
// Parameter address is the address of the provider node.
func NewNotifyItemPublish(address string) *NotifyItemPublish {
	i := &NotifyItemPublish{
		data: NewVariant(),
		info: NewVariant(),
		notifyInfo: &fbs.NotifyInfoT{
			Node:           address,
			Timestamp:      toFiletime(time.Now()),
			NotifyType:     fbs.NotifyTypeData,
			SequenceNumber: 0,
		},
	}
	return i
}

// deleteNotifyItemPublish removes a NotifyItemPublish.
// Parameter n is an instance of the NotifyItemPublish.
func deleteNotifyItemPublish(n *NotifyItemPublish) {
	if n == nil {
		return
	}
	DeleteVariant(n.data)
	DeleteVariant(n.info)
}

// DeleteNotifyItemsPublish removes a lit of NotifyItemPublish.
// Parameter s is list of the NotifyItemPublish.
func DeleteNotifyItemsPublish(l []*NotifyItemPublish) {
	for _, n := range l {
		deleteNotifyItemPublish(n)
	}
}

// GetData gets the data of the notify item.
// It returns the variant.
func (i *NotifyItemPublish) GetData() *Variant {
	return i.data
}

// GetInfo gets the info of the notify item.
// containing notify_info.fbs (address, timestamp, type, ...).
// It returns the variant
func (i *NotifyItemPublish) GetInfo() *Variant {
	builder := flatbuffers.NewBuilder(4096)
	niPack := i.notifyInfo.Pack(builder)
	builder.Finish(niPack)
	i.info.SetFlatbuffers(builder.FinishedBytes())
	return i.info
}

// SetNotifyType sets the NotifyType of the notify_info.fbs.
func (i *NotifyItemPublish) SetNotifyType(n fbs.NotifyType) {
	i.notifyInfo.NotifyType = n
}

// SetTimestamp sets a time.Time value since January 1, 1970 UTC of the notify_info.fbs.
func (i *NotifyItemPublish) SetTimestamp(t time.Time) {
	i.notifyInfo.Timestamp = toFiletime(t)
}

// SetEventType sets eventtype of the notify_info.fbs.
func (i *NotifyItemPublish) SetEventType(e string) {
	i.notifyInfo.EventType = e
}

// SetSequenceNumber sets sequence number of the notify_info.fbs.
func (i *NotifyItemPublish) SetSequenceNumber(n uint64) {
	i.notifyInfo.SequenceNumber = n
}

// SetSequenceNumber sets source name of the notify_info.fbs.
func (i *NotifyItemPublish) SetSourceName(s string) {
	i.notifyInfo.SourceName = s
}
