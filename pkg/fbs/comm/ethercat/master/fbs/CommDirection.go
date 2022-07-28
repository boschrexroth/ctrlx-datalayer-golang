// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

type CommDirection byte

const (
	CommDirectiontx    CommDirection = 0
	CommDirectionrx    CommDirection = 1
	CommDirectiontx_rx CommDirection = 2
)

var EnumNamesCommDirection = map[CommDirection]string{
	CommDirectiontx:    "tx",
	CommDirectionrx:    "rx",
	CommDirectiontx_rx: "tx_rx",
}

var EnumValuesCommDirection = map[string]CommDirection{
	"tx":    CommDirectiontx,
	"rx":    CommDirectionrx,
	"tx_rx": CommDirectiontx_rx,
}

func (v CommDirection) String() string {
	if s, ok := EnumNamesCommDirection[v]; ok {
		return s
	}
	return "CommDirection(" + strconv.FormatInt(int64(v), 10) + ")"
}