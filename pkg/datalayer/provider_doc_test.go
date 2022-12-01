package datalayer_test

import (
	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
)

func ExampleProvider_RegisterType() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the provider
	p := s.Factory().CreateProvider("providerAddress")
	defer datalayer.DeleteProvider(p)

	// Register the node
	p.RegisterType("address", "path")

}

func ExampleProvider_UnregisterType() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the provider
	p := s.Factory().CreateProvider("providerAddress")
	defer datalayer.DeleteProvider(p)

	// Unregisters the node
	p.UnregisterType("address")
}

func ExampleProvider_RegisterNode() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the provider
	p := s.Factory().CreateProvider("providerAddress")
	defer datalayer.DeleteProvider(p)

	// Initialize and destroy the node
	n := datalayer.NewProviderNode()
	defer datalayer.DeleteProviderNode(n)

	// Register the node
	p.RegisterNode("address", n)
}
func ExampleProvider_UnregisterNode() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the provider
	p := s.Factory().CreateProvider("providerAddress")
	defer datalayer.DeleteProvider(p)

	// Initialize and destroy the node
	n := datalayer.NewProviderNode()
	defer datalayer.DeleteProviderNode(n)

	// Register and unregister the node
	p.RegisterNode("address", n)
	p.UnregisterNode("address")
}

func ExampleProvider_SetTimeoutNode() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the provider
	p := s.Factory().CreateProvider("providerAddress")
	defer datalayer.DeleteProvider(p)

	// Initialize and destroy the node
	n := datalayer.NewProviderNode()
	defer datalayer.DeleteProviderNode(n)

	// Set timeout for the node
	p.RegisterNode("address", n)
	p.SetTimeoutNode(n, 1000)
	p.UnregisterNode("address")
}

func ExampleProvider_Start() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the provider
	p := s.Factory().CreateProvider("providerAddress")
	defer datalayer.DeleteProvider(p)

	// Start the provider
	p.Start()
}
func ExampleProvider_Stop() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the provider
	p := s.Factory().CreateProvider("providerAddress")
	defer datalayer.DeleteProvider(p)

	// Start the provider
	p.Start()
	p.Stop()
}

func ExampleProvider_IsConnected() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the provider
	p := s.Factory().CreateProvider("providerAddress")
	defer datalayer.DeleteProvider(p)

	// Check the connection
	p.Start()
	p.IsConnected()
}

func ExampleProvider_GetToken() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the provider
	p := s.Factory().CreateProvider("providerAddress")
	defer datalayer.DeleteProvider(p)

	// Start the provider
	p.Start()

	// Initialize and destroy the variant
	v := p.GetToken()
	defer datalayer.DeleteVariant(v)
}
