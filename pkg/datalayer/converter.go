// MIT License
//
// Copyright (c) 2021 Bosch Rexroth AG
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package datalayer

//#include <stdbool.h>
//#include <stdlib.h>
//#include <converter.h>
import "C"
import (
	"errors"
	"unsafe"
)

// Schema enum
type Schema C.enum_DLR_SCHEMA

// Schema enum definition
const (
	SchemaMetadata   Schema = C.DLR_SCHEMA_METADATA
	SchemaReflection Schema = C.DLR_SCHEMA_REFLECTION
	SchemaMemory     Schema = C.DLR_SCHEMA_MEMORY
	SchemaMemoryMap  Schema = C.DLR_SCHEMA_MEMORY_MAP
	SchemaToken      Schema = C.DLR_SCHEMA_TOKEN
	SchemaProblem    Schema = C.DLR_SCHEMA_PROBLEM
	SchemaDiagnosis  Schema = C.DLR_SCHEMA_DIAGNOSIS
)

// Converter class
type Converter struct {
	this C.DLR_CONVERTER
}

// GenerateJsonSimple generates a JSON string out of a Variant with a simple data type.
// Parameter data is a Variant that contains data with simple data type.
// Parameter indentStep is an indentation length for json string.
// It returns generated JSON as a string.
func (c *Converter) GenerateJsonSimple(data *Variant, indentStep int) (Result, []byte) {
	json := NewVariant()
	defer DeleteVariant(json)
	r := Result(C.DLR_converterGenerateJsonSimple(c.this, data.this, json.this, C.int(indentStep)))
	if r != ResultOk {
		return r, nil
	}
	return r, []byte(json.GetString())
}

// generateJSONComplex converter
func (c *Converter) generateJSONComplex(data Variant, ty Variant, indentStep int) (Result, *Variant) {
	json := NewVariant()
	r := Result(C.DLR_converterGenerateJsonComplex(c.this, data.this, ty.this, json.this, C.int(indentStep)))
	return r, json
}

// GenerateJsonComplex generates a JSON string out of a Variant with a complex type (flatbuffers) and the metadata of this data.
// Parameter data is a Varaint that contains data of complex data type (flatbuffers). If data is empty (VariantType::UNKNOWN) type is converted to json schema.
// Parameter type is a Variant that contains type of data (Variant with flatbuffers BFBS)
// Parameter indentStep is a indentation length for json string.
// It returns generated JSON as a Variant (string).
func (c *Converter) GenerateJsonComplex(data []byte, ty []byte, indentStep int) (Result, []byte) {
	d := NewVariant()
	defer DeleteVariant(d)
	d.SetFlatbuffers(data)
	t := NewVariant()
	defer DeleteVariant(t)
	t.SetFlatbuffers(ty)
	r, json := c.generateJSONComplex(*d, *t, indentStep)
	defer DeleteVariant(json)
	if r != ResultOk {
		return r, nil
	}
	return r, []byte(json.GetString())
}

// ParseJsonSimple generates a Variant out of a JSON string containing the (simple) data.
// Parameter json is a data of the Variant as a json string.
// It returns a result status of the function, data string which contains the data, error Error object.
func (c *Converter) ParseJsonSimple(json []byte) (Result, *Variant, error) {
	cjson := C.CString(string(json))
	defer C.free(unsafe.Pointer(cjson))
	data := NewVariant()
	err := NewVariant()
	defer DeleteVariant(err)
	r := Result(C.DLR_converterParseJsonSimple(c.this, cjson, data.this, err.this))
	if r != ResultOk {
		return r, data, errors.New(err.GetString())
	}
	return r, data, nil
}

// parseJSONComplex converter
func (c *Converter) parseJSONComplex(json []byte, ty Variant) (Result, *Variant, *Variant) {
	cjson := C.CString(string(json))
	defer C.free(unsafe.Pointer(cjson))
	data := NewVariant()
	err := NewVariant()
	r := Result(C.DLR_converterParseJsonComplex(c.this, cjson, ty.this, data.this, err.this))
	return r, data, err
}

// ParseJsonComplex generates a Variant out of a JSON string containing the (complex) data.
// Parameter json is a data of the Variant as a json string.
// Parameter type (ty) is a Variant that contains type of data (Variant with bfbs flatbuffer content).
// It returns a result status of the function, data (array of byte) which contains the data, error Error as error object.
func (c *Converter) ParseJsonComplex(json []byte, ty []byte) (Result, []byte, error) {
	t := NewVariant()
	defer DeleteVariant(t)
	t.SetFlatbuffers(ty)

	r, data, err := c.parseJSONComplex(json, *t)
	defer DeleteVariant(data)
	defer DeleteVariant(err)
	if r != ResultOk {
		return r, []byte{}, errors.New(err.GetString())
	}
	b := make([]byte, len(data.GetFlatbuffers()))
	copy(b, data.GetFlatbuffers())
	return r, b, nil
}

// converterGetSchema converter
func (c *Converter) converterGetSchema(schema Schema) (Result, *Variant) {
	data := NewVariant()
	r := Result(C.DLR_converterGetSchema(c.this, C.DLR_SCHEMA(schema), data.this))
	return r, data
}

// GetSchema returns the type (schema).
// Parameter schema is a requested schema.
// It returns a result status of the function, data which contains the type (schema).
func (c *Converter) GetSchema(schema Schema) (Result, []byte) {
	r, d := c.converterGetSchema(schema)
	defer DeleteVariant(d)
	if r != ResultOk {
		return r, []byte{}
	}
	f := make([]byte, len(d.GetFlatbuffers()))
	copy(f, d.GetFlatbuffers())
	return r, f
}
