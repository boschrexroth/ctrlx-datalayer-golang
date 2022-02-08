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
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	a "github.com/stretchr/testify/assert"
)

const (
	bfbsPathForTest string = "../../deps/public/bfbs/comm.datalayer/comm/datalayer/"
)

func TestGenerateJsonComplex(t *testing.T) {
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(true) })
	a.NotPanics(t, func() {
		system.SetBfbsPath(bfbsPathForTest)
	})

	converter := system.JSONConverter()
	a.NotNil(t, converter)

	r, memory := converter.GetSchema(datalayer.SchemaMemory)
	a.Equal(t, r, datalayer.ResultOk)

	r, reflection := converter.GetSchema(datalayer.SchemaReflection)
	a.Equal(t, r, datalayer.ResultOk)

	r, schematext := converter.GenerateJsonComplex(memory, reflection, -1)
	a.Equal(t, r, datalayer.ResultOk)

	r, data, err := converter.ParseJsonComplex(schematext, reflection)
	a.Equal(t, r, datalayer.ResultOk)
	a.True(t, len(data) != 0)
	a.Nil(t, err)
}
func TestComplexError(t *testing.T) {
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(true) })
	a.NotPanics(t, func() {
		system.SetBfbsPath(bfbsPathForTest)
	})

	converter := system.JSONConverter()
	a.NotNil(t, converter)

	r, reflection := converter.GetSchema(datalayer.SchemaReflection)
	a.Equal(t, r, datalayer.ResultOk)

	r, data, err := converter.ParseJsonComplex([]byte("error"), reflection)
	a.NotEqual(t, r, datalayer.ResultOk)
	a.True(t, len(data) == 0)
	a.NotNil(t, err)
}

func TestGetSchema(t *testing.T) {
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(true) })
	a.NotPanics(t, func() {
		system.SetBfbsPath(bfbsPathForTest)
	})

	converter := system.JSONConverter()
	a.NotNil(t, converter)

	r, buf := converter.GetSchema(datalayer.SchemaMemory)
	a.Equal(t, r, datalayer.ResultOk)
	a.True(t, len(buf) != 0)
}

func TestGenerateJsonSimple(t *testing.T) {
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(true) })
	a.NotPanics(t, func() {
		system.SetBfbsPath(bfbsPathForTest)
	})

	converter := system.JSONConverter()
	a.NotNil(t, converter)

	input := datalayer.NewVariant()
	defer datalayer.DeleteVariant(input)
	input.SetString("Das ist ein Test")
	r, schematext := converter.GenerateJsonSimple(input, -1)
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, schematext, []byte("{\"type\":\"string\",\"value\":\"Das ist ein Test\"}"))
}

func TestParseJsonSimpleString(t *testing.T) {
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(true) })
	a.NotPanics(t, func() {
		system.SetBfbsPath(bfbsPathForTest)
	})

	converter := system.JSONConverter()
	a.NotNil(t, converter)

	input := datalayer.NewVariant()
	defer datalayer.DeleteVariant(input)
	input.SetString("Das ist ein Test")
	r, schematext := converter.GenerateJsonSimple(input, -1)
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, schematext, []byte("{\"type\":\"string\",\"value\":\"Das ist ein Test\"}"))

	r, output, err := converter.ParseJsonSimple(schematext)
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, datalayer.VariantTypeString, output.GetType())
	a.Equal(t, "Das ist ein Test", output.GetString())
	a.Nil(t, err)
}

func TestParseJsonSimpleNumber(t *testing.T) {
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(true) })
	a.NotPanics(t, func() {
		system.SetBfbsPath(bfbsPathForTest)
	})

	converter := system.JSONConverter()
	a.NotNil(t, converter)

	input := datalayer.NewVariant()
	defer datalayer.DeleteVariant(input)
	input.SetInt32(123)
	r, schematext := converter.GenerateJsonSimple(input, -1)
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, []byte("{\"type\":\"int32\",\"value\":123}"), schematext)

	r, data, err := converter.ParseJsonSimple(schematext)
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, datalayer.VariantTypeInt32, data.GetType())
	a.Equal(t, 123, int(data.GetInt32()))
	a.Nil(t, err)
}

func TestSimpleError(t *testing.T) {
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(true) })
	a.NotPanics(t, func() {
		system.SetBfbsPath(bfbsPathForTest)
	})

	converter := system.JSONConverter()
	a.NotNil(t, converter)

	r, _, err := converter.ParseJsonSimple([]byte("error wird gesetzt"))
	a.NotEqual(t, r, datalayer.ResultOk)
	a.NotNil(t, err)
}
