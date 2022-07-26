// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// state of a calculation pipeline
type PipelineState int32

const (
	/// the pipeline has been changed and validation was not successful
	PipelineStateUNVALIDATED PipelineState = 0
	/// the pipeline is currently being edited
	PipelineStateIN_EDIT     PipelineState = 1
	/// the pipeline has been successfully validated and can be used now
	PipelineStateVALIDATED   PipelineState = 2
	/// the pipeline is currently in use and may not be changed
	PipelineStateIN_USE      PipelineState = 3
)

var EnumNamesPipelineState = map[PipelineState]string{
	PipelineStateUNVALIDATED: "UNVALIDATED",
	PipelineStateIN_EDIT:     "IN_EDIT",
	PipelineStateVALIDATED:   "VALIDATED",
	PipelineStateIN_USE:      "IN_USE",
}

var EnumValuesPipelineState = map[string]PipelineState{
	"UNVALIDATED": PipelineStateUNVALIDATED,
	"IN_EDIT":     PipelineStateIN_EDIT,
	"VALIDATED":   PipelineStateVALIDATED,
	"IN_USE":      PipelineStateIN_USE,
}

func (v PipelineState) String() string {
	if s, ok := EnumNamesPipelineState[v]; ok {
		return s
	}
	return "PipelineState(" + strconv.FormatInt(int64(v), 10) + ")"
}
