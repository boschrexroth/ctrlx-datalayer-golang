// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package sdk

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfObjectPoolDiagTypeT struct {
	ObjectPoolDiagArray []*ObjectPoolDiagTypeT `json:"objectPoolDiagArray"`
}

func (t *ArrayOfObjectPoolDiagTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	objectPoolDiagArrayOffset := flatbuffers.UOffsetT(0)
	if t.ObjectPoolDiagArray != nil {
		objectPoolDiagArrayLength := len(t.ObjectPoolDiagArray)
		objectPoolDiagArrayOffsets := make([]flatbuffers.UOffsetT, objectPoolDiagArrayLength)
		for j := 0; j < objectPoolDiagArrayLength; j++ {
			objectPoolDiagArrayOffsets[j] = t.ObjectPoolDiagArray[j].Pack(builder)
		}
		ArrayOfObjectPoolDiagTypeStartObjectPoolDiagArrayVector(builder, objectPoolDiagArrayLength)
		for j := objectPoolDiagArrayLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(objectPoolDiagArrayOffsets[j])
		}
		objectPoolDiagArrayOffset = builder.EndVector(objectPoolDiagArrayLength)
	}
	ArrayOfObjectPoolDiagTypeStart(builder)
	ArrayOfObjectPoolDiagTypeAddObjectPoolDiagArray(builder, objectPoolDiagArrayOffset)
	return ArrayOfObjectPoolDiagTypeEnd(builder)
}

func (rcv *ArrayOfObjectPoolDiagType) UnPackTo(t *ArrayOfObjectPoolDiagTypeT) {
	objectPoolDiagArrayLength := rcv.ObjectPoolDiagArrayLength()
	t.ObjectPoolDiagArray = make([]*ObjectPoolDiagTypeT, objectPoolDiagArrayLength)
	for j := 0; j < objectPoolDiagArrayLength; j++ {
		x := ObjectPoolDiagType{}
		rcv.ObjectPoolDiagArray(&x, j)
		t.ObjectPoolDiagArray[j] = x.UnPack()
	}
}

func (rcv *ArrayOfObjectPoolDiagType) UnPack() *ArrayOfObjectPoolDiagTypeT {
	if rcv == nil { return nil }
	t := &ArrayOfObjectPoolDiagTypeT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfObjectPoolDiagType struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfObjectPoolDiagType(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfObjectPoolDiagType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfObjectPoolDiagType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfObjectPoolDiagType(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfObjectPoolDiagType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfObjectPoolDiagType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfObjectPoolDiagType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfObjectPoolDiagType) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfObjectPoolDiagType) ObjectPoolDiagArray(obj *ObjectPoolDiagType, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ArrayOfObjectPoolDiagType) ObjectPoolDiagArrayLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ArrayOfObjectPoolDiagTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfObjectPoolDiagTypeAddObjectPoolDiagArray(builder *flatbuffers.Builder, objectPoolDiagArray flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(objectPoolDiagArray), 0)
}
func ArrayOfObjectPoolDiagTypeStartObjectPoolDiagArrayVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayOfObjectPoolDiagTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
