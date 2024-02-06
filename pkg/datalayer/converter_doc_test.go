package datalayer_test

import (
	"fmt"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
)

func ExampleConverter_GenerateJsonComplex() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Create a converter
	c := s.JSONConverter()

	r, m := c.GetSchema(datalayer.SchemaMemory)
	fmt.Print(r)
	r, rs := c.GetSchema(datalayer.SchemaReflection)
	r, st := c.GenerateJsonComplex(m, rs, -1)
	r, d, err := c.ParseJsonComplex(st, rs)
	fmt.Print(r, d, err)
}

func ExampleConverter_GetSchema() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Create a converter
	c := s.JSONConverter()

	r, m := c.GetSchema(datalayer.SchemaMemory)
	fmt.Print(r, m)
}

func ExampleConverter_GenerateJsonSimple() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Create a converter
	c := s.JSONConverter()

	// Create and destroy a variant
	i := datalayer.NewVariant()
	defer datalayer.DeleteVariant(i)
	i.SetString("someText")

	// Generate a a JSON string
	r, st := c.GenerateJsonSimple(i, -1)
	fmt.Print(r, st)
}

func ExampleConverter_ParseJsonSimple() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Create a converter
	c := s.JSONConverter()

	// Create and destroy a variant
	i := datalayer.NewVariant()
	defer datalayer.DeleteVariant(i)
	i.SetString("someText")

	// Generate a JSON string
	r, st := c.GenerateJsonSimple(i, -1)
	fmt.Print(r)

	// Parse a string
	r, o, err := c.ParseJsonSimple(st)
	fmt.Print(r, o, err)
}
