// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type SegmentTriggerCondition int8

const (
	SegmentTriggerConditionSEG_TRIGGER_EXTERNAL           SegmentTriggerCondition = 0
	SegmentTriggerConditionSEG_TRIGGER_SLAVE_REACHED_ABS  SegmentTriggerCondition = 1
	SegmentTriggerConditionSEG_TRIGGER_SLAVE_REACHED_REL  SegmentTriggerCondition = 2
	SegmentTriggerConditionSEG_TRIGGER_MASTER_REACHED_ABS SegmentTriggerCondition = 3
	SegmentTriggerConditionSEG_TRIGGER_MASTER_REACHED_REL SegmentTriggerCondition = 4
)

var EnumNamesSegmentTriggerCondition = map[SegmentTriggerCondition]string{
	SegmentTriggerConditionSEG_TRIGGER_EXTERNAL:           "SEG_TRIGGER_EXTERNAL",
	SegmentTriggerConditionSEG_TRIGGER_SLAVE_REACHED_ABS:  "SEG_TRIGGER_SLAVE_REACHED_ABS",
	SegmentTriggerConditionSEG_TRIGGER_SLAVE_REACHED_REL:  "SEG_TRIGGER_SLAVE_REACHED_REL",
	SegmentTriggerConditionSEG_TRIGGER_MASTER_REACHED_ABS: "SEG_TRIGGER_MASTER_REACHED_ABS",
	SegmentTriggerConditionSEG_TRIGGER_MASTER_REACHED_REL: "SEG_TRIGGER_MASTER_REACHED_REL",
}

var EnumValuesSegmentTriggerCondition = map[string]SegmentTriggerCondition{
	"SEG_TRIGGER_EXTERNAL":           SegmentTriggerConditionSEG_TRIGGER_EXTERNAL,
	"SEG_TRIGGER_SLAVE_REACHED_ABS":  SegmentTriggerConditionSEG_TRIGGER_SLAVE_REACHED_ABS,
	"SEG_TRIGGER_SLAVE_REACHED_REL":  SegmentTriggerConditionSEG_TRIGGER_SLAVE_REACHED_REL,
	"SEG_TRIGGER_MASTER_REACHED_ABS": SegmentTriggerConditionSEG_TRIGGER_MASTER_REACHED_ABS,
	"SEG_TRIGGER_MASTER_REACHED_REL": SegmentTriggerConditionSEG_TRIGGER_MASTER_REACHED_REL,
}

func (v SegmentTriggerCondition) String() string {
	if s, ok := EnumNamesSegmentTriggerCondition[v]; ok {
		return s
	}
	return "SegmentTriggerCondition(" + strconv.FormatInt(int64(v), 10) + ")"
}