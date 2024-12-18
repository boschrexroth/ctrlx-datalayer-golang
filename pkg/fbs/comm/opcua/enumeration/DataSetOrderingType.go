// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package enumeration

import "strconv"

type DataSetOrderingType int32

const (
	DataSetOrderingTypeUndefined               DataSetOrderingType = 0
	DataSetOrderingTypeAscendingWriterId       DataSetOrderingType = 1
	DataSetOrderingTypeAscendingWriterIdSingle DataSetOrderingType = 2
)

var EnumNamesDataSetOrderingType = map[DataSetOrderingType]string{
	DataSetOrderingTypeUndefined:               "Undefined",
	DataSetOrderingTypeAscendingWriterId:       "AscendingWriterId",
	DataSetOrderingTypeAscendingWriterIdSingle: "AscendingWriterIdSingle",
}

var EnumValuesDataSetOrderingType = map[string]DataSetOrderingType{
	"Undefined":               DataSetOrderingTypeUndefined,
	"AscendingWriterId":       DataSetOrderingTypeAscendingWriterId,
	"AscendingWriterIdSingle": DataSetOrderingTypeAscendingWriterIdSingle,
}

func (v DataSetOrderingType) String() string {
	if s, ok := EnumNamesDataSetOrderingType[v]; ok {
		return s
	}
	return "DataSetOrderingType(" + strconv.FormatInt(int64(v), 10) + ")"
}
