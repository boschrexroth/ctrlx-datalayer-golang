// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package enumeration

import "strconv"

type PubSubConfigurationRefMask uint32

const (
	PubSubConfigurationRefMaskElementAdd             PubSubConfigurationRefMask = 1
	PubSubConfigurationRefMaskElementMatch           PubSubConfigurationRefMask = 2
	PubSubConfigurationRefMaskElementModify          PubSubConfigurationRefMask = 4
	PubSubConfigurationRefMaskElementRemove          PubSubConfigurationRefMask = 8
	PubSubConfigurationRefMaskReferenceWriter        PubSubConfigurationRefMask = 16
	PubSubConfigurationRefMaskReferenceReader        PubSubConfigurationRefMask = 32
	PubSubConfigurationRefMaskReferenceWriterGroup   PubSubConfigurationRefMask = 64
	PubSubConfigurationRefMaskReferenceReaderGroup   PubSubConfigurationRefMask = 128
	PubSubConfigurationRefMaskReferenceConnection    PubSubConfigurationRefMask = 256
	PubSubConfigurationRefMaskReferencePubDataset    PubSubConfigurationRefMask = 512
	PubSubConfigurationRefMaskReferenceSubDataset    PubSubConfigurationRefMask = 1024
	PubSubConfigurationRefMaskReferenceSecurityGroup PubSubConfigurationRefMask = 2048
	PubSubConfigurationRefMaskReferencePushTarget    PubSubConfigurationRefMask = 4096
)

var EnumNamesPubSubConfigurationRefMask = map[PubSubConfigurationRefMask]string{
	PubSubConfigurationRefMaskElementAdd:             "ElementAdd",
	PubSubConfigurationRefMaskElementMatch:           "ElementMatch",
	PubSubConfigurationRefMaskElementModify:          "ElementModify",
	PubSubConfigurationRefMaskElementRemove:          "ElementRemove",
	PubSubConfigurationRefMaskReferenceWriter:        "ReferenceWriter",
	PubSubConfigurationRefMaskReferenceReader:        "ReferenceReader",
	PubSubConfigurationRefMaskReferenceWriterGroup:   "ReferenceWriterGroup",
	PubSubConfigurationRefMaskReferenceReaderGroup:   "ReferenceReaderGroup",
	PubSubConfigurationRefMaskReferenceConnection:    "ReferenceConnection",
	PubSubConfigurationRefMaskReferencePubDataset:    "ReferencePubDataset",
	PubSubConfigurationRefMaskReferenceSubDataset:    "ReferenceSubDataset",
	PubSubConfigurationRefMaskReferenceSecurityGroup: "ReferenceSecurityGroup",
	PubSubConfigurationRefMaskReferencePushTarget:    "ReferencePushTarget",
}

var EnumValuesPubSubConfigurationRefMask = map[string]PubSubConfigurationRefMask{
	"ElementAdd":             PubSubConfigurationRefMaskElementAdd,
	"ElementMatch":           PubSubConfigurationRefMaskElementMatch,
	"ElementModify":          PubSubConfigurationRefMaskElementModify,
	"ElementRemove":          PubSubConfigurationRefMaskElementRemove,
	"ReferenceWriter":        PubSubConfigurationRefMaskReferenceWriter,
	"ReferenceReader":        PubSubConfigurationRefMaskReferenceReader,
	"ReferenceWriterGroup":   PubSubConfigurationRefMaskReferenceWriterGroup,
	"ReferenceReaderGroup":   PubSubConfigurationRefMaskReferenceReaderGroup,
	"ReferenceConnection":    PubSubConfigurationRefMaskReferenceConnection,
	"ReferencePubDataset":    PubSubConfigurationRefMaskReferencePubDataset,
	"ReferenceSubDataset":    PubSubConfigurationRefMaskReferenceSubDataset,
	"ReferenceSecurityGroup": PubSubConfigurationRefMaskReferenceSecurityGroup,
	"ReferencePushTarget":    PubSubConfigurationRefMaskReferencePushTarget,
}

func (v PubSubConfigurationRefMask) String() string {
	if s, ok := EnumNamesPubSubConfigurationRefMask[v]; ok {
		return s
	}
	return "PubSubConfigurationRefMask(" + strconv.FormatInt(int64(v), 10) + ")"
}
