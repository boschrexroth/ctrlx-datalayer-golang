// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import "strconv"

type MemoryType int8

const (
	MemoryTypeUnknown      MemoryType = 0
	MemoryTypeInput        MemoryType = 1
	MemoryTypeOutput       MemoryType = 2
	MemoryTypeSharedRetain MemoryType = 3
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