// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type SegmentActionMode uint32

const (
	SegmentActionModeSEG_ACTION_EXTERNAL_INSTANT         SegmentActionMode = 0
	SegmentActionModeSEG_ACTION_EXTERNAL_AFTERSEG        SegmentActionMode = 1
	SegmentActionModeSEG_ACTION_EXTERNAL_AFTERGROUP      SegmentActionMode = 2
	SegmentActionModeSEG_ACTION_EXTERNAL_INSTANT_NEXTSEG SegmentActionMode = 3
)

var EnumNamesSegmentActionMode = map[SegmentActionMode]string{
	SegmentActionModeSEG_ACTION_EXTERNAL_INSTANT:         "SEG_ACTION_EXTERNAL_INSTANT",
	SegmentActionModeSEG_ACTION_EXTERNAL_AFTERSEG:        "SEG_ACTION_EXTERNAL_AFTERSEG",
	SegmentActionModeSEG_ACTION_EXTERNAL_AFTERGROUP:      "SEG_ACTION_EXTERNAL_AFTERGROUP",
	SegmentActionModeSEG_ACTION_EXTERNAL_INSTANT_NEXTSEG: "SEG_ACTION_EXTERNAL_INSTANT_NEXTSEG",
}

var EnumValuesSegmentActionMode = map[string]SegmentActionMode{
	"SEG_ACTION_EXTERNAL_INSTANT":         SegmentActionModeSEG_ACTION_EXTERNAL_INSTANT,
	"SEG_ACTION_EXTERNAL_AFTERSEG":        SegmentActionModeSEG_ACTION_EXTERNAL_AFTERSEG,
	"SEG_ACTION_EXTERNAL_AFTERGROUP":      SegmentActionModeSEG_ACTION_EXTERNAL_AFTERGROUP,
	"SEG_ACTION_EXTERNAL_INSTANT_NEXTSEG": SegmentActionModeSEG_ACTION_EXTERNAL_INSTANT_NEXTSEG,
}

func (v SegmentActionMode) String() string {
	if s, ok := EnumNamesSegmentActionMode[v]; ok {
		return s
	}
	return "SegmentActionMode(" + strconv.FormatInt(int64(v), 10) + ")"
}
