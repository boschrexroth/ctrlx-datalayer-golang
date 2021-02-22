// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import "strconv"

type NotifyType int32

const (
	NotifyTypeData     NotifyType = 0
	NotifyTypeBrowse   NotifyType = 1
	NotifyTypeMetadata NotifyType = 2
)

var EnumNamesNotifyType = map[NotifyType]string{
	NotifyTypeData:     "Data",
	NotifyTypeBrowse:   "Browse",
	NotifyTypeMetadata: "Metadata",
}

var EnumValuesNotifyType = map[string]NotifyType{
	"Data":     NotifyTypeData,
	"Browse":   NotifyTypeBrowse,
	"Metadata": NotifyTypeMetadata,
}

func (v NotifyType) String() string {
	if s, ok := EnumNamesNotifyType[v]; ok {
		return s
	}
	return "NotifyType(" + strconv.FormatInt(int64(v), 10) + ")"
}
