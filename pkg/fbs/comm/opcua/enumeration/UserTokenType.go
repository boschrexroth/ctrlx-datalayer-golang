// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package enumeration

import "strconv"

type UserTokenType int32

const (
	UserTokenTypeAnonymous   UserTokenType = 0
	UserTokenTypeUserName    UserTokenType = 1
	UserTokenTypeCertificate UserTokenType = 2
	UserTokenTypeIssuedToken UserTokenType = 3
)

var EnumNamesUserTokenType = map[UserTokenType]string{
	UserTokenTypeAnonymous:   "Anonymous",
	UserTokenTypeUserName:    "UserName",
	UserTokenTypeCertificate: "Certificate",
	UserTokenTypeIssuedToken: "IssuedToken",
}

var EnumValuesUserTokenType = map[string]UserTokenType{
	"Anonymous":   UserTokenTypeAnonymous,
	"UserName":    UserTokenTypeUserName,
	"Certificate": UserTokenTypeCertificate,
	"IssuedToken": UserTokenTypeIssuedToken,
}

func (v UserTokenType) String() string {
	if s, ok := EnumNamesUserTokenType[v]; ok {
		return s
	}
	return "UserTokenType(" + strconv.FormatInt(int64(v), 10) + ")"
}