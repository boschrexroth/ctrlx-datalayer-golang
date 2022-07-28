/**
 * MIT License
 *
 * Copyright (c) 2021-2022 Bosch Rexroth AG
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
package datalayer_test

import (
	"fmt"
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/pkg/fbs/comm/datalayer"
	"github.com/stretchr/testify/assert"
)

func TestMetadataBuilder(t *testing.T) {
	v := datalayer.NewMetaDataBuilder(datalayer.AllowedOperationRead|datalayer.AllowedOperationBrowse, "test", "test_url").Build()
	defer datalayer.DeleteVariant(v)
	assert.NotNil(t, v)

	m := fbs.GetRootAsMetadata(v.GetFlatbuffers(), 0)
	assert.NotNil(t, m)
	assert.Equal(t, string(m.Description()), "test")
	assert.Equal(t, string(m.DescriptionUrl()), "test_url")
	op := m.Operations(nil)
	assert.NotNil(t, op)
	assert.True(t, op.Read())
	assert.True(t, op.Browse())
	assert.False(t, op.Write())
}
func TestMetadataReferenceBuilder(t *testing.T) {
	m := datalayer.NewMetaDataBuilder(datalayer.AllowedOperationRead|datalayer.AllowedOperationWrite|datalayer.AllowedOperationBrowse, "test", "test_url")

	m.AddReference(datalayer.ReferenceTypeWrite, "type_path_to_"+datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeWrite])
	m.AddReference(datalayer.ReferenceTypeRead, "type_path_to_"+datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeRead])
	m.AddReference(datalayer.ReferenceTypeWriteOut, "type_path_to_"+datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeWriteOut])
	m.AddReference(datalayer.ReferenceTypeReadIn, "type_path_to_"+datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeReadIn])
	m.AddReference(datalayer.ReferenceTypeCreate, "type_path_to_"+datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeCreate])
	m.AddReference(datalayer.ReferenceTypeWriteIn, "type_path_to_"+datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeWriteIn])
	m.AddReference(datalayer.ReferenceTypeReadOut, "type_path_to_"+datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeReadOut])

	v := m.Build()
	defer datalayer.DeleteVariant(v)
	assert.NotNil(t, v)
	d := fbs.GetRootAsMetadata(v.GetFlatbuffers(), 0)
	assert.NotNil(t, d)

	lst := []string{
		datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeCreate],
		datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeReadIn],
		datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeReadOut],
		datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeRead],
		datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeWriteIn],
		datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeWriteOut],
		datalayer.EnumNamesReferenceType[datalayer.ReferenceTypeWrite],
	}
	assert.Equal(t, d.ReferencesLength(), len(lst))

	for i, s := range lst {
		r := &fbs.Reference{}
		b := d.References(r, i)
		assert.True(t, b)
		assert.NotNil(t, r)
		fmt.Println(i, s, string(r.Type()), string(r.TargetAddress()))
		assert.Equal(t, string(r.Type()), s)
		assert.Equal(t, string(r.TargetAddress()), "type_path_to_"+s)
	}
}

func TestMetadataExtensionsBuilder(t *testing.T) {
	m := datalayer.NewMetaDataBuilder(datalayer.AllowedOperationRead|datalayer.AllowedOperationBrowse, "test", "test_url")
	lst := []string{"1key", "2key", "3key", "4key"}
	m.AddExtension(lst[2], "value_"+lst[2])
	m.AddExtension(lst[1], "value_"+lst[1])
	m.AddExtension(lst[3], "value_"+lst[3])
	m.AddExtension(lst[0], "value_"+lst[0])
	v := m.Build()
	defer datalayer.DeleteVariant(v)
	assert.NotNil(t, v)
	d := fbs.GetRootAsMetadata(v.GetFlatbuffers(), 0)
	assert.NotNil(t, d)
	assert.Equal(t, d.ExtensionsLength(), len(lst))
	for i, s := range lst {
		e := &fbs.Extension{}
		b := d.Extensions(e, i)
		assert.True(t, b)
		assert.NotNil(t, e)
		fmt.Println(i, s, string(e.Key()), string(e.Value()))
		assert.Equal(t, string(e.Key()), s)
		assert.Equal(t, string(e.Value()), "value_"+s)
	}
}
