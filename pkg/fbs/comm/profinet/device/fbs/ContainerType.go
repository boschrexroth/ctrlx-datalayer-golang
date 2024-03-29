// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

type ContainerType uint16

const (
	ContainerTypeTYPE_UNDEFINED      ContainerType = 0
	ContainerTypeTYPE_BYTE           ContainerType = 1
	ContainerTypeTYPE_ARRAY_OF_BYTE  ContainerType = 2
	ContainerTypeTYPE_WORD           ContainerType = 3
	ContainerTypeTYPE_ARRAY_OF_WORD  ContainerType = 4
	ContainerTypeTYPE_DWORD          ContainerType = 5
	ContainerTypeTYPE_ARRAY_OF_DWORD ContainerType = 6
	ContainerTypeTYPE_LWORD          ContainerType = 7
	ContainerTypeTYPE_ARRAY_OF_LWORD ContainerType = 8
	ContainerTypeTYPE_SINT           ContainerType = 9
	ContainerTypeTYPE_ARRAY_OF_SINT  ContainerType = 10
	ContainerTypeTYPE_USINT          ContainerType = 11
	ContainerTypeTYPE_ARRAY_OF_USINT ContainerType = 12
	ContainerTypeTYPE_INT            ContainerType = 13
	ContainerTypeTYPE_ARRAY_OF_INT   ContainerType = 14
	ContainerTypeTYPE_UINT           ContainerType = 15
	ContainerTypeTYPE_ARRAY_OF_UINT  ContainerType = 16
	ContainerTypeTYPE_DINT           ContainerType = 17
	ContainerTypeTYPE_ARRAY_OF_DINT  ContainerType = 18
	ContainerTypeTYPE_UDINT          ContainerType = 19
	ContainerTypeTYPE_ARRAY_OF_UDINT ContainerType = 20
	ContainerTypeTYPE_LINT           ContainerType = 21
	ContainerTypeTYPE_ARRAY_OF_LINT  ContainerType = 22
	ContainerTypeTYPE_ULINT          ContainerType = 23
	ContainerTypeTYPE_ARRAY_OF_ULINT ContainerType = 24
	ContainerTypeTYPE_REAL           ContainerType = 25
	ContainerTypeTYPE_ARRAY_OF_REAL  ContainerType = 26
	ContainerTypeTYPE_LREAL          ContainerType = 27
	ContainerTypeTYPE_ARRAY_OF_LREAL ContainerType = 28
)

var EnumNamesContainerType = map[ContainerType]string{
	ContainerTypeTYPE_UNDEFINED:      "TYPE_UNDEFINED",
	ContainerTypeTYPE_BYTE:           "TYPE_BYTE",
	ContainerTypeTYPE_ARRAY_OF_BYTE:  "TYPE_ARRAY_OF_BYTE",
	ContainerTypeTYPE_WORD:           "TYPE_WORD",
	ContainerTypeTYPE_ARRAY_OF_WORD:  "TYPE_ARRAY_OF_WORD",
	ContainerTypeTYPE_DWORD:          "TYPE_DWORD",
	ContainerTypeTYPE_ARRAY_OF_DWORD: "TYPE_ARRAY_OF_DWORD",
	ContainerTypeTYPE_LWORD:          "TYPE_LWORD",
	ContainerTypeTYPE_ARRAY_OF_LWORD: "TYPE_ARRAY_OF_LWORD",
	ContainerTypeTYPE_SINT:           "TYPE_SINT",
	ContainerTypeTYPE_ARRAY_OF_SINT:  "TYPE_ARRAY_OF_SINT",
	ContainerTypeTYPE_USINT:          "TYPE_USINT",
	ContainerTypeTYPE_ARRAY_OF_USINT: "TYPE_ARRAY_OF_USINT",
	ContainerTypeTYPE_INT:            "TYPE_INT",
	ContainerTypeTYPE_ARRAY_OF_INT:   "TYPE_ARRAY_OF_INT",
	ContainerTypeTYPE_UINT:           "TYPE_UINT",
	ContainerTypeTYPE_ARRAY_OF_UINT:  "TYPE_ARRAY_OF_UINT",
	ContainerTypeTYPE_DINT:           "TYPE_DINT",
	ContainerTypeTYPE_ARRAY_OF_DINT:  "TYPE_ARRAY_OF_DINT",
	ContainerTypeTYPE_UDINT:          "TYPE_UDINT",
	ContainerTypeTYPE_ARRAY_OF_UDINT: "TYPE_ARRAY_OF_UDINT",
	ContainerTypeTYPE_LINT:           "TYPE_LINT",
	ContainerTypeTYPE_ARRAY_OF_LINT:  "TYPE_ARRAY_OF_LINT",
	ContainerTypeTYPE_ULINT:          "TYPE_ULINT",
	ContainerTypeTYPE_ARRAY_OF_ULINT: "TYPE_ARRAY_OF_ULINT",
	ContainerTypeTYPE_REAL:           "TYPE_REAL",
	ContainerTypeTYPE_ARRAY_OF_REAL:  "TYPE_ARRAY_OF_REAL",
	ContainerTypeTYPE_LREAL:          "TYPE_LREAL",
	ContainerTypeTYPE_ARRAY_OF_LREAL: "TYPE_ARRAY_OF_LREAL",
}

var EnumValuesContainerType = map[string]ContainerType{
	"TYPE_UNDEFINED":      ContainerTypeTYPE_UNDEFINED,
	"TYPE_BYTE":           ContainerTypeTYPE_BYTE,
	"TYPE_ARRAY_OF_BYTE":  ContainerTypeTYPE_ARRAY_OF_BYTE,
	"TYPE_WORD":           ContainerTypeTYPE_WORD,
	"TYPE_ARRAY_OF_WORD":  ContainerTypeTYPE_ARRAY_OF_WORD,
	"TYPE_DWORD":          ContainerTypeTYPE_DWORD,
	"TYPE_ARRAY_OF_DWORD": ContainerTypeTYPE_ARRAY_OF_DWORD,
	"TYPE_LWORD":          ContainerTypeTYPE_LWORD,
	"TYPE_ARRAY_OF_LWORD": ContainerTypeTYPE_ARRAY_OF_LWORD,
	"TYPE_SINT":           ContainerTypeTYPE_SINT,
	"TYPE_ARRAY_OF_SINT":  ContainerTypeTYPE_ARRAY_OF_SINT,
	"TYPE_USINT":          ContainerTypeTYPE_USINT,
	"TYPE_ARRAY_OF_USINT": ContainerTypeTYPE_ARRAY_OF_USINT,
	"TYPE_INT":            ContainerTypeTYPE_INT,
	"TYPE_ARRAY_OF_INT":   ContainerTypeTYPE_ARRAY_OF_INT,
	"TYPE_UINT":           ContainerTypeTYPE_UINT,
	"TYPE_ARRAY_OF_UINT":  ContainerTypeTYPE_ARRAY_OF_UINT,
	"TYPE_DINT":           ContainerTypeTYPE_DINT,
	"TYPE_ARRAY_OF_DINT":  ContainerTypeTYPE_ARRAY_OF_DINT,
	"TYPE_UDINT":          ContainerTypeTYPE_UDINT,
	"TYPE_ARRAY_OF_UDINT": ContainerTypeTYPE_ARRAY_OF_UDINT,
	"TYPE_LINT":           ContainerTypeTYPE_LINT,
	"TYPE_ARRAY_OF_LINT":  ContainerTypeTYPE_ARRAY_OF_LINT,
	"TYPE_ULINT":          ContainerTypeTYPE_ULINT,
	"TYPE_ARRAY_OF_ULINT": ContainerTypeTYPE_ARRAY_OF_ULINT,
	"TYPE_REAL":           ContainerTypeTYPE_REAL,
	"TYPE_ARRAY_OF_REAL":  ContainerTypeTYPE_ARRAY_OF_REAL,
	"TYPE_LREAL":          ContainerTypeTYPE_LREAL,
	"TYPE_ARRAY_OF_LREAL": ContainerTypeTYPE_ARRAY_OF_LREAL,
}

func (v ContainerType) String() string {
	if s, ok := EnumNamesContainerType[v]; ok {
		return s
	}
	return "ContainerType(" + strconv.FormatInt(int64(v), 10) + ")"
}
