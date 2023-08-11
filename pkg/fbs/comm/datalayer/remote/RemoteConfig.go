// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package remote

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RemoteConfigT struct {
	Remotes []*ConfigItemT `json:"remotes"`
}

func (t *RemoteConfigT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	remotesOffset := flatbuffers.UOffsetT(0)
	if t.Remotes != nil {
		remotesLength := len(t.Remotes)
		remotesOffsets := make([]flatbuffers.UOffsetT, remotesLength)
		for j := 0; j < remotesLength; j++ {
			remotesOffsets[j] = t.Remotes[j].Pack(builder)
		}
		RemoteConfigStartRemotesVector(builder, remotesLength)
		for j := remotesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(remotesOffsets[j])
		}
		remotesOffset = builder.EndVector(remotesLength)
	}
	RemoteConfigStart(builder)
	RemoteConfigAddRemotes(builder, remotesOffset)
	return RemoteConfigEnd(builder)
}

func (rcv *RemoteConfig) UnPackTo(t *RemoteConfigT) {
	remotesLength := rcv.RemotesLength()
	t.Remotes = make([]*ConfigItemT, remotesLength)
	for j := 0; j < remotesLength; j++ {
		x := ConfigItem{}
		rcv.Remotes(&x, j)
		t.Remotes[j] = x.UnPack()
	}
}

func (rcv *RemoteConfig) UnPack() *RemoteConfigT {
	if rcv == nil { return nil }
	t := &RemoteConfigT{}
	rcv.UnPackTo(t)
	return t
}

type RemoteConfig struct {
	_tab flatbuffers.Table
}

func GetRootAsRemoteConfig(buf []byte, offset flatbuffers.UOffsetT) *RemoteConfig {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RemoteConfig{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRemoteConfig(buf []byte, offset flatbuffers.UOffsetT) *RemoteConfig {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RemoteConfig{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RemoteConfig) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RemoteConfig) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RemoteConfig) Remotes(obj *ConfigItem, j int) bool {
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

func (rcv *RemoteConfig) RemotesByKey(obj *ConfigItem, key string) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *RemoteConfig) RemotesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func RemoteConfigStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func RemoteConfigAddRemotes(builder *flatbuffers.Builder, remotes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(remotes), 0)
}
func RemoteConfigStartRemotesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RemoteConfigEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
