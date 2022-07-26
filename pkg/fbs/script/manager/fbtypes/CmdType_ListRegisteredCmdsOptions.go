// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type CmdTypeListRegisteredCmdsOptions int8

const (
	CmdType_ListRegisteredCmdsOptionsUNKNOWN    CmdTypeListRegisteredCmdsOptions = 0
	CmdType_ListRegisteredCmdsOptionsCMD        CmdTypeListRegisteredCmdsOptions = 1
	CmdType_ListRegisteredCmdsOptionsCMD_OPTION CmdTypeListRegisteredCmdsOptions = 2
)

var EnumNamesCmdType_ListRegisteredCmdsOptions = map[CmdTypeListRegisteredCmdsOptions]string{
	CmdType_ListRegisteredCmdsOptionsUNKNOWN:    "UNKNOWN",
	CmdType_ListRegisteredCmdsOptionsCMD:        "CMD",
	CmdType_ListRegisteredCmdsOptionsCMD_OPTION: "CMD_OPTION",
}

var EnumValuesCmdType_ListRegisteredCmdsOptions = map[string]CmdTypeListRegisteredCmdsOptions{
	"UNKNOWN":    CmdType_ListRegisteredCmdsOptionsUNKNOWN,
	"CMD":        CmdType_ListRegisteredCmdsOptionsCMD,
	"CMD_OPTION": CmdType_ListRegisteredCmdsOptionsCMD_OPTION,
}

func (v CmdTypeListRegisteredCmdsOptions) String() string {
	if s, ok := EnumNamesCmdType_ListRegisteredCmdsOptions[v]; ok {
		return s
	}
	return "CmdTypeListRegisteredCmdsOptions(" + strconv.FormatInt(int64(v), 10) + ")"
}
