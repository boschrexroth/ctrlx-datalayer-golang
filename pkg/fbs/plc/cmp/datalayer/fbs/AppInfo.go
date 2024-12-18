// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

/// informations of the application as set in plc engineering
type AppInfoT struct {
	Name string `json:"name"`
	Author string `json:"author"`
	Version string `json:"version"`
	Description string `json:"description"`
	Profile string `json:"profile"`
	Compiler string `json:"compiler"`
	DateAndTime *AppDateAndTimeT `json:"dateAndTime"`
}

func (t *AppInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	authorOffset := flatbuffers.UOffsetT(0)
	if t.Author != "" {
		authorOffset = builder.CreateString(t.Author)
	}
	versionOffset := flatbuffers.UOffsetT(0)
	if t.Version != "" {
		versionOffset = builder.CreateString(t.Version)
	}
	descriptionOffset := flatbuffers.UOffsetT(0)
	if t.Description != "" {
		descriptionOffset = builder.CreateString(t.Description)
	}
	profileOffset := flatbuffers.UOffsetT(0)
	if t.Profile != "" {
		profileOffset = builder.CreateString(t.Profile)
	}
	compilerOffset := flatbuffers.UOffsetT(0)
	if t.Compiler != "" {
		compilerOffset = builder.CreateString(t.Compiler)
	}
	dateAndTimeOffset := t.DateAndTime.Pack(builder)
	AppInfoStart(builder)
	AppInfoAddName(builder, nameOffset)
	AppInfoAddAuthor(builder, authorOffset)
	AppInfoAddVersion(builder, versionOffset)
	AppInfoAddDescription(builder, descriptionOffset)
	AppInfoAddProfile(builder, profileOffset)
	AppInfoAddCompiler(builder, compilerOffset)
	AppInfoAddDateAndTime(builder, dateAndTimeOffset)
	return AppInfoEnd(builder)
}

func (rcv *AppInfo) UnPackTo(t *AppInfoT) {
	t.Name = string(rcv.Name())
	t.Author = string(rcv.Author())
	t.Version = string(rcv.Version())
	t.Description = string(rcv.Description())
	t.Profile = string(rcv.Profile())
	t.Compiler = string(rcv.Compiler())
	t.DateAndTime = rcv.DateAndTime(nil).UnPack()
}

func (rcv *AppInfo) UnPack() *AppInfoT {
	if rcv == nil { return nil }
	t := &AppInfoT{}
	rcv.UnPackTo(t)
	return t
}

type AppInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsAppInfo(buf []byte, offset flatbuffers.UOffsetT) *AppInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AppInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAppInfo(buf []byte, offset flatbuffers.UOffsetT) *AppInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AppInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AppInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AppInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the plc project
func (rcv *AppInfo) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the plc project
func AppInfoKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &AppInfo{}
	obj2 := &AppInfo{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.Name()) < string(obj2.Name())
}

func (rcv *AppInfo) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &AppInfo{}
		obj.Init(buf, tableOffset)
		comp := bytes.Compare(obj.Name(), bKey)
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

/// author of the application
func (rcv *AppInfo) Author() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// author of the application
/// version of the application
func (rcv *AppInfo) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// version of the application
/// descpriotion of the application
func (rcv *AppInfo) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// descpriotion of the application
/// PLC engineering version
func (rcv *AppInfo) Profile() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// PLC engineering version
/// used compiler version for the loaded application
func (rcv *AppInfo) Compiler() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// used compiler version for the loaded application
/// last time this application was loaded
func (rcv *AppInfo) DateAndTime(obj *AppDateAndTime) *AppDateAndTime {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AppDateAndTime)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// last time this application was loaded
func AppInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func AppInfoAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func AppInfoAddAuthor(builder *flatbuffers.Builder, author flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(author), 0)
}
func AppInfoAddVersion(builder *flatbuffers.Builder, version flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(version), 0)
}
func AppInfoAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(description), 0)
}
func AppInfoAddProfile(builder *flatbuffers.Builder, profile flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(profile), 0)
}
func AppInfoAddCompiler(builder *flatbuffers.Builder, compiler flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(compiler), 0)
}
func AppInfoAddDateAndTime(builder *flatbuffers.Builder, dateAndTime flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(dateAndTime), 0)
}
func AppInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
