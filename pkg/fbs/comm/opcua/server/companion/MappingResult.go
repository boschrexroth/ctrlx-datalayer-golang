// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package companion

import "strconv"

type MappingResult int8

const (
	/// Mapping is good
	MappingResultGood                    MappingResult = 0
	/// sourceUaNodeId from datalayer is not valid
	MappingResultInvalidSource           MappingResult = 1
	/// targetUaNodeId is not present in the loaded companion models
	MappingResultInvalidTarget           MappingResult = 2
	/// targetUaNodeId is used multiple times in the mapping tables
	MappingResultTargetAlreadyUsed       MappingResult = 3
	/// Could not load the type definition of the targetUaNodeId from companion model
	MappingResultTargetTypeUnsupported   MappingResult = 4
	/// sourceUaNodeId and targetUaNodeId have different datatypes, only checked if 'typeSafety' is true
	MappingResultTypeMismatch            MappingResult = 5
	/// sourceUaNodeId does not provide valid metadata
	MappingResultSourceNodeHasNoMetadata MappingResult = 6
)

var EnumNamesMappingResult = map[MappingResult]string{
	MappingResultGood:                    "Good",
	MappingResultInvalidSource:           "InvalidSource",
	MappingResultInvalidTarget:           "InvalidTarget",
	MappingResultTargetAlreadyUsed:       "TargetAlreadyUsed",
	MappingResultTargetTypeUnsupported:   "TargetTypeUnsupported",
	MappingResultTypeMismatch:            "TypeMismatch",
	MappingResultSourceNodeHasNoMetadata: "SourceNodeHasNoMetadata",
}

var EnumValuesMappingResult = map[string]MappingResult{
	"Good":                    MappingResultGood,
	"InvalidSource":           MappingResultInvalidSource,
	"InvalidTarget":           MappingResultInvalidTarget,
	"TargetAlreadyUsed":       MappingResultTargetAlreadyUsed,
	"TargetTypeUnsupported":   MappingResultTargetTypeUnsupported,
	"TypeMismatch":            MappingResultTypeMismatch,
	"SourceNodeHasNoMetadata": MappingResultSourceNodeHasNoMetadata,
}

func (v MappingResult) String() string {
	if s, ok := EnumNamesMappingResult[v]; ok {
		return s
	}
	return "MappingResult(" + strconv.FormatInt(int64(v), 10) + ")"
}
