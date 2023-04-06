// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import "strconv"

type MemoryType int8

const (
	MemoryTypeUnknown      MemoryType = 0
	/// data flow from owner to user
	MemoryTypeInput        MemoryType = 1
	/// data flow from user to owner
	MemoryTypeOutput       MemoryType = 2
	/// shared ram in retain area
	MemoryTypeSharedRetain MemoryType = 3
	/// shared ram
	MemoryTypeShared       MemoryType = 4
)

var EnumNamesMemoryType = map[MemoryType]string{
	MemoryTypeUnknown:      "Unknown",
	MemoryTypeInput:        "Input",
	MemoryTypeOutput:       "Output",
	MemoryTypeSharedRetain: "SharedRetain",
	MemoryTypeShared:       "Shared",
}

var EnumValuesMemoryType = map[string]MemoryType{
	"Unknown":      MemoryTypeUnknown,
	"Input":        MemoryTypeInput,
	"Output":       MemoryTypeOutput,
	"SharedRetain": MemoryTypeSharedRetain,
	"Shared":       MemoryTypeShared,
}

func (v MemoryType) String() string {
	if s, ok := EnumNamesMemoryType[v]; ok {
		return s
	}
	return "MemoryType(" + strconv.FormatInt(int64(v), 10) + ")"
}
