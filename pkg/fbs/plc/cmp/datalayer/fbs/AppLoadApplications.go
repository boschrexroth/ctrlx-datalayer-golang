// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AppLoadApplicationsT struct {
	Path string `json:"path"`
	Files []string `json:"files"`
}

func (t *AppLoadApplicationsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	pathOffset := flatbuffers.UOffsetT(0)
	if t.Path != "" {
		pathOffset = builder.CreateString(t.Path)
	}
	filesOffset := flatbuffers.UOffsetT(0)
	if t.Files != nil {
		filesLength := len(t.Files)
		filesOffsets := make([]flatbuffers.UOffsetT, filesLength)
		for j := 0; j < filesLength; j++ {
			filesOffsets[j] = builder.CreateString(t.Files[j])
		}
		AppLoadApplicationsStartFilesVector(builder, filesLength)
		for j := filesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(filesOffsets[j])
		}
		filesOffset = builder.EndVector(filesLength)
	}
	AppLoadApplicationsStart(builder)
	AppLoadApplicationsAddPath(builder, pathOffset)
	AppLoadApplicationsAddFiles(builder, filesOffset)
	return AppLoadApplicationsEnd(builder)
}

func (rcv *AppLoadApplications) UnPackTo(t *AppLoadApplicationsT) {
	t.Path = string(rcv.Path())
	filesLength := rcv.FilesLength()
	t.Files = make([]string, filesLength)
	for j := 0; j < filesLength; j++ {
		t.Files[j] = string(rcv.Files(j))
	}
}

func (rcv *AppLoadApplications) UnPack() *AppLoadApplicationsT {
	if rcv == nil { return nil }
	t := &AppLoadApplicationsT{}
	rcv.UnPackTo(t)
	return t
}

type AppLoadApplications struct {
	_tab flatbuffers.Table
}

func GetRootAsAppLoadApplications(buf []byte, offset flatbuffers.UOffsetT) *AppLoadApplications {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AppLoadApplications{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAppLoadApplications(buf []byte, offset flatbuffers.UOffsetT) *AppLoadApplications {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AppLoadApplications{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AppLoadApplications) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AppLoadApplications) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AppLoadApplications) Path() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AppLoadApplications) Files(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *AppLoadApplications) FilesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func AppLoadApplicationsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AppLoadApplicationsAddPath(builder *flatbuffers.Builder, path flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(path), 0)
}
func AppLoadApplicationsAddFiles(builder *flatbuffers.Builder, files flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(files), 0)
}
func AppLoadApplicationsStartFilesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func AppLoadApplicationsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
