package datalayer

//#include <stdbool.h>
//#include <stdlib.h>
//#include "converter.h"
import "C"
import "unsafe"

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

// GenerateJSONSimple converter
func (c *Converter) GenerateJSONSimple(data Variant, json Variant, indentStep int) Result {
	return Result(C.DLR_converterGenerateJsonSimple(c.this, data.this, json.this, C.int(indentStep)))
}

// GenerateJSONComplex converter
func (c *Converter) GenerateJSONComplex(data Variant, ty Variant, json Variant, indentStep int) Result {
	return Result(C.DLR_converterGenerateJsonComplex(c.this, data.this, ty.this, json.this, C.int(indentStep)))
}

// ParseJSONSimple converter
func (c *Converter) ParseJSONSimple(json string, data Variant, err Variant) Result {
	cjson := C.CString(json)
	defer C.free(unsafe.Pointer(cjson))
	return Result(C.DLR_converterParseJsonSimple(c.this, cjson, data.this, err.this))
}

// ParseJSONComplex converter
func (c *Converter) ParseJSONComplex(json string, ty Variant, data Variant, err Variant) Result {
	cjson := C.CString(json)
	defer C.free(unsafe.Pointer(cjson))
	return Result(C.DLR_converterParseJsonComplex(c.this, cjson, ty.this, data.this, err.this))
}

// ConverterGetSchema converter
func (c *Converter) ConverterGetSchema(schema Schema, data Variant) Result {
	return Result(C.DLR_converterGetSchema(c.this, C.DLR_SCHEMA(schema), data.this))
}
