// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CfgGlobalT struct {
	InitScript *CfgInitScriptT
}

func (t *CfgGlobalT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	initScriptOffset := t.InitScript.Pack(builder)
	CfgGlobalStart(builder)
	CfgGlobalAddInitScript(builder, initScriptOffset)
	return CfgGlobalEnd(builder)
}

func (rcv *CfgGlobal) UnPackTo(t *CfgGlobalT) {
	t.InitScript = rcv.InitScript(nil).UnPack()
}

func (rcv *CfgGlobal) UnPack() *CfgGlobalT {
	if rcv == nil { return nil }
	t := &CfgGlobalT{}
	rcv.UnPackTo(t)
	return t
}

type CfgGlobal struct {
	_tab flatbuffers.Table
}

func GetRootAsCfgGlobal(buf []byte, offset flatbuffers.UOffsetT) *CfgGlobal {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CfgGlobal{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCfgGlobal(buf []byte, offset flatbuffers.UOffsetT) *CfgGlobal {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CfgGlobal{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CfgGlobal) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CfgGlobal) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CfgGlobal) InitScript(obj *CfgInitScript) *CfgInitScript {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(CfgInitScript)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func CfgGlobalStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func CfgGlobalAddInitScript(builder *flatbuffers.Builder, initScript flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(initScript), 0)
}
func CfgGlobalEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
