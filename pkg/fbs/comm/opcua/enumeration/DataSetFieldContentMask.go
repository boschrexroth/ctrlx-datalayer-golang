// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package enumeration

import "strconv"

type DataSetFieldContentMask uint32

const (
	DataSetFieldContentMaskStatusCode        DataSetFieldContentMask = 1
	DataSetFieldContentMaskSourceTimestamp   DataSetFieldContentMask = 2
	DataSetFieldContentMaskServerTimestamp   DataSetFieldContentMask = 4
	DataSetFieldContentMaskSourcePicoSeconds DataSetFieldContentMask = 8
	DataSetFieldContentMaskServerPicoSeconds DataSetFieldContentMask = 16
	DataSetFieldContentMaskRawData           DataSetFieldContentMask = 32
)

var EnumNamesDataSetFieldContentMask = map[DataSetFieldContentMask]string{
	DataSetFieldContentMaskStatusCode:        "StatusCode",
	DataSetFieldContentMaskSourceTimestamp:   "SourceTimestamp",
	DataSetFieldContentMaskServerTimestamp:   "ServerTimestamp",
	DataSetFieldContentMaskSourcePicoSeconds: "SourcePicoSeconds",
	DataSetFieldContentMaskServerPicoSeconds: "ServerPicoSeconds",
	DataSetFieldContentMaskRawData:           "RawData",
}

var EnumValuesDataSetFieldContentMask = map[string]DataSetFieldContentMask{
	"StatusCode":        DataSetFieldContentMaskStatusCode,
	"SourceTimestamp":   DataSetFieldContentMaskSourceTimestamp,
	"ServerTimestamp":   DataSetFieldContentMaskServerTimestamp,
	"SourcePicoSeconds": DataSetFieldContentMaskSourcePicoSeconds,
	"ServerPicoSeconds": DataSetFieldContentMaskServerPicoSeconds,
	"RawData":           DataSetFieldContentMaskRawData,
}

func (v DataSetFieldContentMask) String() string {
	if s, ok := EnumNamesDataSetFieldContentMask[v]; ok {
		return s
	}
	return "DataSetFieldContentMask(" + strconv.FormatInt(int64(v), 10) + ")"
}
