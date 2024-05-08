// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// version information (uses semantic versioning) of a single motion component
type ComponentVersionInfoT struct {
	Name string `json:"name"`
	MajorVersion uint32 `json:"majorVersion"`
	MinorVersion uint32 `json:"minorVersion"`
	PatchVersion uint32 `json:"patchVersion"`
	BuildNumber uint32 `json:"buildNumber"`
	Branch string `json:"branch"`
	VersionString string `json:"versionString"`
	Commit string `json:"commit"`
	Type MotionComponentType `json:"type"`
	AddInfo []string `json:"addInfo"`
}

func (t *ComponentVersionInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	branchOffset := flatbuffers.UOffsetT(0)
	if t.Branch != "" {
		branchOffset = builder.CreateString(t.Branch)
	}
	versionStringOffset := flatbuffers.UOffsetT(0)
	if t.VersionString != "" {
		versionStringOffset = builder.CreateString(t.VersionString)
	}
	commitOffset := flatbuffers.UOffsetT(0)
	if t.Commit != "" {
		commitOffset = builder.CreateString(t.Commit)
	}
	addInfoOffset := flatbuffers.UOffsetT(0)
	if t.AddInfo != nil {
		addInfoLength := len(t.AddInfo)
		addInfoOffsets := make([]flatbuffers.UOffsetT, addInfoLength)
		for j := 0; j < addInfoLength; j++ {
			addInfoOffsets[j] = builder.CreateString(t.AddInfo[j])
		}
		ComponentVersionInfoStartAddInfoVector(builder, addInfoLength)
		for j := addInfoLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(addInfoOffsets[j])
		}
		addInfoOffset = builder.EndVector(addInfoLength)
	}
	ComponentVersionInfoStart(builder)
	ComponentVersionInfoAddName(builder, nameOffset)
	ComponentVersionInfoAddMajorVersion(builder, t.MajorVersion)
	ComponentVersionInfoAddMinorVersion(builder, t.MinorVersion)
	ComponentVersionInfoAddPatchVersion(builder, t.PatchVersion)
	ComponentVersionInfoAddBuildNumber(builder, t.BuildNumber)
	ComponentVersionInfoAddBranch(builder, branchOffset)
	ComponentVersionInfoAddVersionString(builder, versionStringOffset)
	ComponentVersionInfoAddCommit(builder, commitOffset)
	ComponentVersionInfoAddType(builder, t.Type)
	ComponentVersionInfoAddAddInfo(builder, addInfoOffset)
	return ComponentVersionInfoEnd(builder)
}

func (rcv *ComponentVersionInfo) UnPackTo(t *ComponentVersionInfoT) {
	t.Name = string(rcv.Name())
	t.MajorVersion = rcv.MajorVersion()
	t.MinorVersion = rcv.MinorVersion()
	t.PatchVersion = rcv.PatchVersion()
	t.BuildNumber = rcv.BuildNumber()
	t.Branch = string(rcv.Branch())
	t.VersionString = string(rcv.VersionString())
	t.Commit = string(rcv.Commit())
	t.Type = rcv.Type()
	addInfoLength := rcv.AddInfoLength()
	t.AddInfo = make([]string, addInfoLength)
	for j := 0; j < addInfoLength; j++ {
		t.AddInfo[j] = string(rcv.AddInfo(j))
	}
}

func (rcv *ComponentVersionInfo) UnPack() *ComponentVersionInfoT {
	if rcv == nil { return nil }
	t := &ComponentVersionInfoT{}
	rcv.UnPackTo(t)
	return t
}

type ComponentVersionInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsComponentVersionInfo(buf []byte, offset flatbuffers.UOffsetT) *ComponentVersionInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ComponentVersionInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsComponentVersionInfo(buf []byte, offset flatbuffers.UOffsetT) *ComponentVersionInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ComponentVersionInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ComponentVersionInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ComponentVersionInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the motion component
func (rcv *ComponentVersionInfo) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the motion component
/// major version number of the motion component
func (rcv *ComponentVersionInfo) MajorVersion() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// major version number of the motion component
func (rcv *ComponentVersionInfo) MutateMajorVersion(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// minor version number of the motion component
func (rcv *ComponentVersionInfo) MinorVersion() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// minor version number of the motion component
func (rcv *ComponentVersionInfo) MutateMinorVersion(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

/// patch version number of the motion component
func (rcv *ComponentVersionInfo) PatchVersion() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// patch version number of the motion component
func (rcv *ComponentVersionInfo) MutatePatchVersion(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

/// build number of the motion component
func (rcv *ComponentVersionInfo) BuildNumber() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// build number of the motion component
func (rcv *ComponentVersionInfo) MutateBuildNumber(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

/// name of the branch, in which this motion component was built
func (rcv *ComponentVersionInfo) Branch() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the branch, in which this motion component was built
/// complete version name string
func (rcv *ComponentVersionInfo) VersionString() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// complete version name string
/// ID of the commit, that was built for the component
func (rcv *ComponentVersionInfo) Commit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// ID of the commit, that was built for the component
/// type of the motion component
func (rcv *ComponentVersionInfo) Type() MotionComponentType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return MotionComponentType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// type of the motion component
func (rcv *ComponentVersionInfo) MutateType(n MotionComponentType) bool {
	return rcv._tab.MutateInt8Slot(20, int8(n))
}

/// additional informations (not yet used)
func (rcv *ComponentVersionInfo) AddInfo(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *ComponentVersionInfo) AddInfoLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// additional informations (not yet used)
func ComponentVersionInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func ComponentVersionInfoAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ComponentVersionInfoAddMajorVersion(builder *flatbuffers.Builder, majorVersion uint32) {
	builder.PrependUint32Slot(1, majorVersion, 0)
}
func ComponentVersionInfoAddMinorVersion(builder *flatbuffers.Builder, minorVersion uint32) {
	builder.PrependUint32Slot(2, minorVersion, 0)
}
func ComponentVersionInfoAddPatchVersion(builder *flatbuffers.Builder, patchVersion uint32) {
	builder.PrependUint32Slot(3, patchVersion, 0)
}
func ComponentVersionInfoAddBuildNumber(builder *flatbuffers.Builder, buildNumber uint32) {
	builder.PrependUint32Slot(4, buildNumber, 0)
}
func ComponentVersionInfoAddBranch(builder *flatbuffers.Builder, branch flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(branch), 0)
}
func ComponentVersionInfoAddVersionString(builder *flatbuffers.Builder, versionString flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(versionString), 0)
}
func ComponentVersionInfoAddCommit(builder *flatbuffers.Builder, commit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(commit), 0)
}
func ComponentVersionInfoAddType(builder *flatbuffers.Builder, type_ MotionComponentType) {
	builder.PrependInt8Slot(8, int8(type_), 0)
}
func ComponentVersionInfoAddAddInfo(builder *flatbuffers.Builder, addInfo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(addInfo), 0)
}
func ComponentVersionInfoStartAddInfoVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ComponentVersionInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}