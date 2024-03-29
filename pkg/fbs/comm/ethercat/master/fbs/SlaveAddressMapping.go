// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///This node delivers the EtherCAT addresses of the configured slave devices from the EtherCAT Master.
///Find slave name by address or find address by slave name
///Leave request empty to get a response of all configured slaves
type SlaveAddressMappingT struct {
	Request *SlaveAddressMappingRequestT `json:"request"`
	Response *SlaveAddressMappingResponseT `json:"response"`
}

func (t *SlaveAddressMappingT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	responseOffset := t.Response.Pack(builder)
	SlaveAddressMappingStart(builder)
	SlaveAddressMappingAddRequest(builder, requestOffset)
	SlaveAddressMappingAddResponse(builder, responseOffset)
	return SlaveAddressMappingEnd(builder)
}

func (rcv *SlaveAddressMapping) UnPackTo(t *SlaveAddressMappingT) {
	t.Request = rcv.Request(nil).UnPack()
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *SlaveAddressMapping) UnPack() *SlaveAddressMappingT {
	if rcv == nil { return nil }
	t := &SlaveAddressMappingT{}
	rcv.UnPackTo(t)
	return t
}

type SlaveAddressMapping struct {
	_tab flatbuffers.Table
}

func GetRootAsSlaveAddressMapping(buf []byte, offset flatbuffers.UOffsetT) *SlaveAddressMapping {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SlaveAddressMapping{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSlaveAddressMapping(buf []byte, offset flatbuffers.UOffsetT) *SlaveAddressMapping {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SlaveAddressMapping{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SlaveAddressMapping) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SlaveAddressMapping) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SlaveAddressMapping) Request(obj *SlaveAddressMappingRequest) *SlaveAddressMappingRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SlaveAddressMappingRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SlaveAddressMapping) Response(obj *SlaveAddressMappingResponse) *SlaveAddressMappingResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SlaveAddressMappingResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func SlaveAddressMappingStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SlaveAddressMappingAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func SlaveAddressMappingAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(response), 0)
}
func SlaveAddressMappingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
