package datalayer_test

import (
	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
)

func ExampleSystem_Start() {
	// Initialize the system
	s := datalayer.NewSystem("")
	// Destroy the system
	defer datalayer.DeleteSystem(s)

	//Run a dalayer system
	s.Start(true)
}

func ExampleSystem_Stop() {
	// Initialize the system
	s := datalayer.NewSystem("")
	// Destroy the system
	defer datalayer.DeleteSystem(s)

	// Run and stop a dalayer system
	s.Start(true)
	s.Stop(true)
}

func ExampleSystem_Factory() {
	// Initialize the system
	s := datalayer.NewSystem("")
	// Destroy the system
	defer datalayer.DeleteSystem(s)

	// Create a factory
	s.Factory()
}

func ExampleSystem_JSONConverter() {
	// Initialize the system
	s := datalayer.NewSystem("")
	// Destroy the system
	defer datalayer.DeleteSystem(s)

	// Create a converter
	s.JSONConverter()
}

func ExampleSystem_SetBfbsPath() {
	// Initialize the system
	s := datalayer.NewSystem("")
	// Destroy the system
	defer datalayer.DeleteSystem(s)

	// Set the base path to bfbs files
	s.SetBfbsPath("somePath")
}
