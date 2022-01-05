// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import "strconv"

type NodeClass int8

const (
	NodeClassNode       NodeClass = 0
	NodeClassMethod     NodeClass = 1
	NodeClassType       NodeClass = 2
	NodeClassVariable   NodeClass = 3
	NodeClassCollection NodeClass = 4
	NodeClassResource   NodeClass = 5
	NodeClassProgram    NodeClass = 6
	NodeClassFolder     NodeClass = 7
)

var EnumNamesNodeClass = map[NodeClass]string{
	NodeClassNode:       "Node",
	NodeClassMethod:     "Method",
	NodeClassType:       "Type",
	NodeClassVariable:   "Variable",
	NodeClassCollection: "Collection",
	NodeClassResource:   "Resource",
	NodeClassProgram:    "Program",
	NodeClassFolder:     "Folder",
}

var EnumValuesNodeClass = map[string]NodeClass{
	"Node":       NodeClassNode,
	"Method":     NodeClassMethod,
	"Type":       NodeClassType,
	"Variable":   NodeClassVariable,
	"Collection": NodeClassCollection,
	"Resource":   NodeClassResource,
	"Program":    NodeClassProgram,
	"Folder":     NodeClassFolder,
}

func (v NodeClass) String() string {
	if s, ok := EnumNamesNodeClass[v]; ok {
		return s
	}
	return "NodeClass(" + strconv.FormatInt(int64(v), 10) + ")"
}
