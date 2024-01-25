// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package drive

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

type ConnectionStatusT struct {
	Key string `json:"key"`
	Value Status `json:"value"`
}

func (t *ConnectionStatusT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	keyOffset := flatbuffers.UOffsetT(0)
	if t.Key != "" {
		keyOffset = builder.CreateString(t.Key)
	}
	ConnectionStatusStart(builder)
	ConnectionStatusAddKey(builder, keyOffset)
	ConnectionStatusAddValue(builder, t.Value)
	return ConnectionStatusEnd(builder)
}

func (rcv *ConnectionStatus) UnPackTo(t *ConnectionStatusT) {
	t.Key = string(rcv.Key())
	t.Value = rcv.Value()
}

func (rcv *ConnectionStatus) UnPack() *ConnectionStatusT {
	if rcv == nil { return nil }
	t := &ConnectionStatusT{}
	rcv.UnPackTo(t)
	return t
}

type ConnectionStatus struct {
	_tab flatbuffers.Table
}

func GetRootAsConnectionStatus(buf []byte, offset flatbuffers.UOffsetT) *ConnectionStatus {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ConnectionStatus{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsConnectionStatus(buf []byte, offset flatbuffers.UOffsetT) *ConnectionStatus {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ConnectionStatus{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ConnectionStatus) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ConnectionStatus) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ConnectionStatus) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ConnectionStatusKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &ConnectionStatus{}
	obj2 := &ConnectionStatus{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.Key()) < string(obj2.Key())
}

func (rcv *ConnectionStatus) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &ConnectionStatus{}
		obj.Init(buf, tableOffset)
		comp := bytes.Compare(obj.Key(), bKey)
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

/// corresponding value
func (rcv *ConnectionStatus) Value() Status {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return Status(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// corresponding value
func (rcv *ConnectionStatus) MutateValue(n Status) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func ConnectionStatusStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ConnectionStatusAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func ConnectionStatusAddValue(builder *flatbuffers.Builder, value Status) {
	builder.PrependInt8Slot(1, int8(value), 0)
}
func ConnectionStatusEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
