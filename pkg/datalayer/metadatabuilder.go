/**
 * MIT License
 *
 * Copyright (c) 2021-2023 Bosch Rexroth AG
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package datalayer

import (
	"sort"

	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
	flatbuffers "github.com/google/flatbuffers/go"
)

// AllowedOperation enum
type AllowedOperation uint

// AllowedOperation enum defintion
const (
	AllowedOperationNone   AllowedOperation = 0x00000
	AllowedOperationRead   AllowedOperation = 0x00001
	AllowedOperationWrite  AllowedOperation = 0x00010
	AllowedOperationCreate AllowedOperation = 0x00100
	AllowedOperationDelete AllowedOperation = 0x01000
	AllowedOperationBrowse AllowedOperation = 0x10000
	AllowedOperationAll    AllowedOperation = AllowedOperationRead | AllowedOperationWrite | AllowedOperationCreate | AllowedOperationDelete | AllowedOperationBrowse
)

// ReferenceType enum
type ReferenceType uint8

// ReferenceType enum defintion
const (
	ReferenceTypeRead     ReferenceType = 0
	ReferenceTypeReadIn   ReferenceType = 1
	ReferenceTypeReadOut  ReferenceType = 2
	ReferenceTypeWrite    ReferenceType = 3
	ReferenceTypeWriteIn  ReferenceType = 4
	ReferenceTypeWriteOut ReferenceType = 5
	ReferenceTypeCreate   ReferenceType = 6
	ReferenceTypeUses     ReferenceType = 7
	ReferenceTypeHasSave  ReferenceType = 8
)

// EnumNamesReferenceType enum defintion
var EnumNamesReferenceType = map[ReferenceType]string{
	ReferenceTypeRead:     "readType",
	ReferenceTypeReadIn:   "readInType",
	ReferenceTypeReadOut:  "readOutType",
	ReferenceTypeWrite:    "writeType",
	ReferenceTypeWriteIn:  "writeInType",
	ReferenceTypeWriteOut: "writeOutType",
	ReferenceTypeCreate:   "createType",
	ReferenceTypeUses:     "uses",
	ReferenceTypeHasSave:  "hasSave",
}

// MetaDataBuilder struct
type MetaDataBuilder struct {
	name           string
	description    string
	descriptionurl string
	unit           string
	allowed        AllowedOperation
	nodeclass      fbs.NodeClass
	displayformat  fbs.DisplayFormat
	refers         map[string]string
	extensions     map[string]string
	descriptions   map[string]string
	displaynames   map[string]string
}

// NewMetaDataBuilder generates MetaDataBuilder instance
func NewMetaDataBuilder(a AllowedOperation, desc string, descurl string) *MetaDataBuilder {
	m := &MetaDataBuilder{description: desc, descriptionurl: descurl}
	m.name = ""
	m.unit = ""
	m.nodeclass = fbs.NodeClassNode
	m.displayformat = fbs.DisplayFormatAuto
	m.refers = make(map[string]string)
	m.extensions = make(map[string]string)
	m.descriptions = make(map[string]string)
	m.displaynames = make(map[string]string)
	m.Operations(a)
	return m
}

// Build this instance
func (m *MetaDataBuilder) Build() *Variant {
	builder := flatbuffers.NewBuilder(4096)

	//Serialize Descriptions data
	descriptions := m.builddescriptions()

	//Serialize DisplayNames data
	displaynames := m.builddisplaynames()

	//Serialize References data
	references := m.buildreferences()

	//Serialize Extensions data
	extensions := m.buildextensions()

	//Serialize AllowedOperations data
	operations := m.operations()

	meta := fbs.MetadataT{}
	meta.Description = m.description
	meta.DescriptionUrl = m.descriptionurl
	meta.DisplayName = m.name
	meta.Unit = m.unit
	meta.DisplayFormat = m.displayformat

	meta.NodeClass = m.nodeclass
	meta.Operations = operations
	meta.References = references
	meta.Extensions = extensions
	meta.Descriptions = descriptions
	meta.DisplayNames = displaynames
	mi := meta.Pack(builder)
	builder.Finish(mi)
	v := NewVariant()
	v.SetFlatbuffers(builder.FinishedBytes())
	return v
}

func (m *MetaDataBuilder) buildextensions() []*fbs.ExtensionT {

	extensions := []*fbs.ExtensionT{}
	keys := sortkeys(m.extensions)

	for _, k := range keys {
		re := m.addExt(k, m.extensions[k])
		extensions = append(extensions, re)
	}

	return extensions
}

func (m *MetaDataBuilder) addExt(key, val string) *fbs.ExtensionT {
	ext := &fbs.ExtensionT{Key: key, Value: val}
	return ext
}

func sortkeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (m *MetaDataBuilder) buildreferences() []*fbs.ReferenceT {

	references := []*fbs.ReferenceT{}
	keys := sortkeys(m.refers)
	for _, k := range keys {
		//fmt.Println("refs: ", k, m.refers[k])
		re := m.addRef(k, m.refers[k])
		references = append(references, re)
	}

	return references
}

func (m *MetaDataBuilder) addRef(ty, addr string) *fbs.ReferenceT {
	ref := &fbs.ReferenceT{Type: ty, TargetAddress: addr}
	return ref
}

func (m *MetaDataBuilder) operations() *fbs.AllowedOperationsT {
	ops := &fbs.AllowedOperationsT{}
	ops.Read = m.isallowed(AllowedOperationRead)
	ops.Browse = m.isallowed(AllowedOperationBrowse)
	ops.Create = m.isallowed(AllowedOperationCreate)
	ops.Delete = m.isallowed(AllowedOperationDelete)
	ops.Write = m.isallowed(AllowedOperationWrite)
	return ops
}

// Operations set the allowed operations
func (m *MetaDataBuilder) Operations(a AllowedOperation) *MetaDataBuilder {
	m.allowed = a
	return m
}

func (m *MetaDataBuilder) isallowed(a AllowedOperation) bool {
	return m.allowed&a == a
}

// Unit sets the unit
func (m *MetaDataBuilder) Unit(u string) *MetaDataBuilder {
	m.unit = u
	return m
}

// DisplayName sets the display name
func (m *MetaDataBuilder) DisplayName(n string) *MetaDataBuilder {
	m.name = n
	return m
}

// NodeClass sets the node class
func (m *MetaDataBuilder) NodeClass(nc fbs.NodeClass) *MetaDataBuilder {
	m.nodeclass = nc
	return m
}

// DisplayFormat sets the display format
func (m *MetaDataBuilder) DisplayFormat(df fbs.DisplayFormat) *MetaDataBuilder {
	m.displayformat = df
	return m
}

// AddReference adds the reference
func (m *MetaDataBuilder) AddReference(r ReferenceType, t string) *MetaDataBuilder {
	m.refers[ref2name(r)] = t
	return m
}

// AddExtension adds the extension
func (m *MetaDataBuilder) AddExtension(key string, val string) *MetaDataBuilder {
	m.extensions[key] = val
	return m
}

func ref2name(r ReferenceType) string {
	if s, ok := EnumNamesReferenceType[r]; ok {
		return s
	}
	return ""
}

func (m *MetaDataBuilder) builddescriptions() []*fbs.LocaleTextT {
	return m.buildLocalText(m.descriptions)
}

func (m *MetaDataBuilder) builddisplaynames() []*fbs.LocaleTextT {
	return m.buildLocalText(m.displaynames)
}

func (m *MetaDataBuilder) buildLocalText(l map[string]string) []*fbs.LocaleTextT {
	texts := []*fbs.LocaleTextT{}
	keys := sortkeys(l)

	for _, k := range keys {
		re := m.addLocalText(k, l[k])
		texts = append(texts, re)
	}

	return texts
}

func (m *MetaDataBuilder) addLocalText(key, val string) *fbs.LocaleTextT {
	lt := &fbs.LocaleTextT{Id: key, Text: val}
	return lt
}

// AddLocalizationDescription adds localization of description
func (m *MetaDataBuilder) AddLocalizationDescription(id, txt string) *MetaDataBuilder {
	m.descriptions[id] = txt
	return m
}

// AddLocalizationDisplayName adds localization of display name
func (m *MetaDataBuilder) AddLocalizationDisplayName(id, txt string) *MetaDataBuilder {
	m.displaynames[id] = txt
	return m
}
