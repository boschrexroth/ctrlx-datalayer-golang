// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package companion

import "strconv"

type MirrorMappingResult int8

const (
	/// Mapping is good
	MirrorMappingResultGood              MirrorMappingResult = 0
	/// sourceUaNodeId from Server Address Space is invalid
	MirrorMappingResultInvalidSource     MirrorMappingResult = 1
	/// targetDlAddress is used multiple times in the mapping tables
	MirrorMappingResultTargetAlreadyUsed MirrorMappingResult = 2
)

var EnumNamesMirrorMappingResult = map[MirrorMappingResult]string{
	MirrorMappingResultGood:              "Good",
	MirrorMappingResultInvalidSource:     "InvalidSource",
	MirrorMappingResultTargetAlreadyUsed: "TargetAlreadyUsed",
}

var EnumValuesMirrorMappingResult = map[string]MirrorMappingResult{
	"Good":              MirrorMappingResultGood,
	"InvalidSource":     MirrorMappingResultInvalidSource,
	"TargetAlreadyUsed": MirrorMappingResultTargetAlreadyUsed,
}

func (v MirrorMappingResult) String() string {
	if s, ok := EnumNamesMirrorMappingResult[v]; ok {
		return s
	}
	return "MirrorMappingResult(" + strconv.FormatInt(int64(v), 10) + ")"
}
