// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///This node reads/writes a parameter from/to an EtherCAT slave device. 
///The slave must support the mailbox protocol SoE (Servo drive over EtherCAT). 
///Note: the Slave must be in EtherCAT state PreOP, SafeOP or OP for mailbox communication.
type ParameterT struct {
	Request *ParameterRequestT `json:"request"`
	Response *ParameterResponseT `json:"response"`
}

func (t *ParameterT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	responseOffset := t.Response.Pack(builder)
	ParameterStart(builder)
	ParameterAddRequest(builder, requestOffset)
	ParameterAddResponse(builder, responseOffset)
	return ParameterEnd(builder)
}

func (rcv *Parameter) UnPackTo(t *ParameterT) {
	t.Request = rcv.Request(nil).UnPack()
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *Parameter) UnPack() *ParameterT {
	if rcv == nil { return nil }
	t := &ParameterT{}
	rcv.UnPackTo(t)
	return t
}

type Parameter struct {
	_tab flatbuffers.Table
}

func GetRootAsParameter(buf []byte, offset flatbuffers.UOffsetT) *Parameter {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Parameter{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsParameter(buf []byte, offset flatbuffers.UOffsetT) *Parameter {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Parameter{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Parameter) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Parameter) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Parameter) Request(obj *ParameterRequest) *ParameterRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ParameterRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Parameter) Response(obj *ParameterResponse) *ParameterResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ParameterResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ParameterStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ParameterAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func ParameterAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(response), 0)
}
func ParameterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
