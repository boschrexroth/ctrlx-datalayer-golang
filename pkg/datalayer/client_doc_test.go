package datalayer_test

import (
	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
)

func ExampleClient_Browse() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Browse returns all subnodes of address
	c.Browse("address of the node")
}

func ExampleClient_BrowseSync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// BrowseSync searches a node
	c.BrowseSync("address of the node")
}

func ExampleClient_CreateSync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Create and destroy the variant
	v := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v)

	// CreateSync creates an object
	c.CreateSync("address of the node", v)
}

func ExampleClient_ReadSyncArgs() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Create and destroy the variant
	v := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v)

	// Read an object
	c.ReadSyncArgs("address of the node", v)
}

func ExampleClient_ReadSync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Read an object
	c.ReadSync("address of the node")
}

func ExampleClient_WriteSync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Create and destroy the variant
	v := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v)

	// Write an object
	c.WriteSync("address of the node", v)
}

func ExampleClient_MetadataSync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Read metadata of an object
	c.MetadataSync("address of the node")
}

func ExampleClient_Metadata() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Read metadata of a node
	c.Metadata("address of the node")
}

func ExampleClient_GetAuthToken() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Get persistent security access token for authentication
	c.GetAuthToken()
}

func ExampleClient_RemoveSync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Remove the client
	c.RemoveSync("")
}

func ExampleClient_ReadJsonSync() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Initialize the converter
	cv := s.JSONConverter()

	// Read a values as a JSON string
	c.ReadJsonSync(cv, "address of the node", 2)
}

func ExampleClient_WriteJsonSync() {
	// Initialize and destroy the system
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)

	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	// Initialize the converter
	cv := s.JSONConverter()

	// Write a JSON value
	c.ReadJsonSync(cv, "address of the node", 2)
}
