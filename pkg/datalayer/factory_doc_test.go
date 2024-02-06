package datalayer_test

import (
	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
)

func ExampleFactory_CreateProvider() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Create a factory
	f := s.Factory()

	// Create and destroy a provider
	p := f.CreateProvider("")
	defer datalayer.DeleteProvider(p)
}

func ExampleFactory_CreateClient() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Create a factory
	f := s.Factory()

	// Create and destroy a client
	c := f.CreateClient("name")
	defer datalayer.DeleteClient(c)
}
