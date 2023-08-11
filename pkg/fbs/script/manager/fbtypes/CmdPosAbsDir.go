// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// Selected direction for PosAbs with modulo axis
type CmdPosAbsDir int8

const (
	/// use shortest way to reach target position
	CmdPosAbsDirSHORTEST_WAY CmdPosAbsDir = 0
	/// move only in positive direction to reach target position
	CmdPosAbsDirPOSITIVE_DIR CmdPosAbsDir = 1
	/// move only in negative direction to reach target position
	CmdPosAbsDirNEGATIVE_DIR CmdPosAbsDir = 2
	/// when currently moving: continue to move in this direction to reach target position; 
	/// when in stand still: use SHORTEST_WAY
	CmdPosAbsDirCURRENT_DIR  CmdPosAbsDir = 3
)

var EnumNamesCmdPosAbsDir = map[CmdPosAbsDir]string{
	CmdPosAbsDirSHORTEST_WAY: "SHORTEST_WAY",
	CmdPosAbsDirPOSITIVE_DIR: "POSITIVE_DIR",
	CmdPosAbsDirNEGATIVE_DIR: "NEGATIVE_DIR",
	CmdPosAbsDirCURRENT_DIR:  "CURRENT_DIR",
}

var EnumValuesCmdPosAbsDir = map[string]CmdPosAbsDir{
	"SHORTEST_WAY": CmdPosAbsDirSHORTEST_WAY,
	"POSITIVE_DIR": CmdPosAbsDirPOSITIVE_DIR,
	"NEGATIVE_DIR": CmdPosAbsDirNEGATIVE_DIR,
	"CURRENT_DIR":  CmdPosAbsDirCURRENT_DIR,
}

func (v CmdPosAbsDir) String() string {
	if s, ok := EnumNamesCmdPosAbsDir[v]; ok {
		return s
	}
	return "CmdPosAbsDir(" + strconv.FormatInt(int64(v), 10) + ")"
}