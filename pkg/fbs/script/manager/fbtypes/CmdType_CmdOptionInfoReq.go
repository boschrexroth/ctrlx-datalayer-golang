// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type CmdTypeCmdOptionInfoReq int8

const (
	CmdType_CmdOptionInfoReqUNKNOWN    CmdTypeCmdOptionInfoReq = 0
	CmdType_CmdOptionInfoReqCMD        CmdTypeCmdOptionInfoReq = 1
	CmdType_CmdOptionInfoReqCMD_OPTION CmdTypeCmdOptionInfoReq = 2
)

var EnumNamesCmdType_CmdOptionInfoReq = map[CmdTypeCmdOptionInfoReq]string{
	CmdType_CmdOptionInfoReqUNKNOWN:    "UNKNOWN",
	CmdType_CmdOptionInfoReqCMD:        "CMD",
	CmdType_CmdOptionInfoReqCMD_OPTION: "CMD_OPTION",
}

var EnumValuesCmdType_CmdOptionInfoReq = map[string]CmdTypeCmdOptionInfoReq{
	"UNKNOWN":    CmdType_CmdOptionInfoReqUNKNOWN,
	"CMD":        CmdType_CmdOptionInfoReqCMD,
	"CMD_OPTION": CmdType_CmdOptionInfoReqCMD_OPTION,
}

func (v CmdTypeCmdOptionInfoReq) String() string {
	if s, ok := EnumNamesCmdType_CmdOptionInfoReq[v]; ok {
		return s
	}
	return "CmdType_CmdOptionInfoReq(" + strconv.FormatInt(int64(v), 10) + ")"
}
