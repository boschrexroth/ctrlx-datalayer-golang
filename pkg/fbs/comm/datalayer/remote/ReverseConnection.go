// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package remote

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

type ReverseConnectionT struct {
	ConnectionId string `json:"connectionId"`
	ConnectionAddress string `json:"connectionAddress"`
	Name string `json:"name"`
}

func (t *ReverseConnectionT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	connectionIdOffset := flatbuffers.UOffsetT(0)
	if t.ConnectionId != "" {
		connectionIdOffset = builder.CreateString(t.ConnectionId)
	}
	connectionAddressOffset := flatbuffers.UOffsetT(0)
	if t.ConnectionAddress != "" {
		connectionAddressOffset = builder.CreateString(t.ConnectionAddress)
	}
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	ReverseConnectionStart(builder)
	ReverseConnectionAddConnectionId(builder, connectionIdOffset)
	ReverseConnectionAddConnectionAddress(builder, connectionAddressOffset)
	ReverseConnectionAddName(builder, nameOffset)
	return ReverseConnectionEnd(builder)
}

func (rcv *ReverseConnection) UnPackTo(t *ReverseConnectionT) {
	t.ConnectionId = string(rcv.ConnectionId())
	t.ConnectionAddress = string(rcv.ConnectionAddress())
	t.Name = string(rcv.Name())
}

func (rcv *ReverseConnection) UnPack() *ReverseConnectionT {
	if rcv == nil { return nil }
	t := &ReverseConnectionT{}
	rcv.UnPackTo(t)
	return t
}

type ReverseConnection struct {
	_tab flatbuffers.Table
}

func GetRootAsReverseConnection(buf []byte, offset flatbuffers.UOffsetT) *ReverseConnection {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ReverseConnection{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsReverseConnection(buf []byte, offset flatbuffers.UOffsetT) *ReverseConnection {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ReverseConnection{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ReverseConnection) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ReverseConnection) Table() flatbuffers.Table {
	return rcv._tab
}

/// local id of reverse connection
func (rcv *ReverseConnection) ConnectionId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// local id of reverse connection
func ReverseConnectionKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &ReverseConnection{}
	obj2 := &ReverseConnection{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.ConnectionId()) < string(obj2.ConnectionId())
}

func (rcv *ReverseConnection) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &ReverseConnection{}
		obj.Init(buf, tableOffset)
		comp := bytes.Compare(obj.ConnectionId(), bKey)
		if comp > 0 {
			span = middle
		} else if comp < 0 {
			middle += 1
			start += middle
			span -= middle
		} else {
			rcv.Init(buf, tableOffset)
			return true
		}
	}
	return false
}

/// address to remote control to connect to (connect string)
func (rcv *ReverseConnection) ConnectionAddress() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// address to remote control to connect to (connect string)
/// name the datalayer appears on remote server
func (rcv *ReverseConnection) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name the datalayer appears on remote server
func ReverseConnectionStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ReverseConnectionAddConnectionId(builder *flatbuffers.Builder, connectionId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(connectionId), 0)
}
func ReverseConnectionAddConnectionAddress(builder *flatbuffers.Builder, connectionAddress flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(connectionAddress), 0)
}
func ReverseConnectionAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(name), 0)
}
func ReverseConnectionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
