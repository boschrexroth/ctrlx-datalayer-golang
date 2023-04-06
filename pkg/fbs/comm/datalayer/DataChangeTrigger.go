// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import "strconv"

type DataChangeTrigger int32

const (
	/// Inform if STATUS changes
	DataChangeTriggerStatus               DataChangeTrigger = 0
	/// Inform if Value or Status changed
	DataChangeTriggerStatusValue          DataChangeTrigger = 1
	/// Inform if Status Or Status Or Timestamp changes (inform always)
	DataChangeTriggerStatusValueTimestamp DataChangeTrigger = 2
)

var EnumNamesDataChangeTrigger = map[DataChangeTrigger]string{
	DataChangeTriggerStatus:               "Status",
	DataChangeTriggerStatusValue:          "StatusValue",
	DataChangeTriggerStatusValueTimestamp: "StatusValueTimestamp",
}

var EnumValuesDataChangeTrigger = map[string]DataChangeTrigger{
	"Status":               DataChangeTriggerStatus,
	"StatusValue":          DataChangeTriggerStatusValue,
	"StatusValueTimestamp": DataChangeTriggerStatusValueTimestamp,
}

func (v DataChangeTrigger) String() string {
	if s, ok := EnumNamesDataChangeTrigger[v]; ok {
		return s
	}
	return "DataChangeTrigger(" + strconv.FormatInt(int64(v), 10) + ")"
}
