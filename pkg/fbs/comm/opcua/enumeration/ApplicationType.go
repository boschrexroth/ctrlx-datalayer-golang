// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package enumeration

import "strconv"

type ApplicationType int32

const (
	ApplicationTypeServer          ApplicationType = 0
	ApplicationTypeClient          ApplicationType = 1
	ApplicationTypeClientAndServer ApplicationType = 2
	ApplicationTypeDiscoveryServer ApplicationType = 3
)

var EnumNamesApplicationType = map[ApplicationType]string{
	ApplicationTypeServer:          "Server",
	ApplicationTypeClient:          "Client",
	ApplicationTypeClientAndServer: "ClientAndServer",
	ApplicationTypeDiscoveryServer: "DiscoveryServer",
}

var EnumValuesApplicationType = map[string]ApplicationType{
	"Server":          ApplicationTypeServer,
	"Client":          ApplicationTypeClient,
	"ClientAndServer": ApplicationTypeClientAndServer,
	"DiscoveryServer": ApplicationTypeDiscoveryServer,
}

func (v ApplicationType) String() string {
	if s, ok := EnumNamesApplicationType[v]; ok {
		return s
	}
	return "ApplicationType(" + strconv.FormatInt(int64(v), 10) + ")"
}