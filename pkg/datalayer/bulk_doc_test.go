package datalayer_test

import (
	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
)

func ExampleBulk_Browse() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	r := bulk.Browse("address of the node1", "address of the node2")
	if r != datalayer.ResultOk {
		return
	}
	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
		data := bulk.Response[key].Data
		for _, v := range data.GetArrayString() {
			println(v)
		}
	}
}

func ExampleBulk_BrowseAsync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	//callback function
	rc := func(resps []datalayer.Response) {
	}

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	r := bulk.BrowseAsync(rc, "address of the node1", "address of the node2")
	if r != datalayer.ResultOk {
		return
	}
	// waiting for callback
}

func ExampleBulk_Delete() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	r := bulk.Delete("address of the node1", "address of the node2")
	if r != datalayer.ResultOk {
		return
	}
	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
	}
}

func ExampleBulk_DeleteAsync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	//callback function
	rc := func(resps []datalayer.Response) {
	}

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	r := bulk.DeleteAsync(rc, "address of the node1", "address of the node2")
	if r != datalayer.ResultOk {
		return
	}
	// waiting for callback
}

func ExampleBulk_Metadata() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	r := bulk.Metadata("address of the node1", "address of the node2")
	if r != datalayer.ResultOk {
		return
	}
	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
	}
}

func ExampleBulk_MetadataAsync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	//callback function
	rc := func(resps []datalayer.Response) {
	}

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	r := bulk.MetadataAsync(rc, "address of the node1", "address of the node2")
	if r != datalayer.ResultOk {
		return
	}
	// waiting for callback
}

func ExampleBulk_Read() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	r := bulk.Read(datalayer.BulkReadArg{Address: "address of the node1", Argument: nil}, datalayer.BulkReadArg{Address: "address of the node2"})
	if r != datalayer.ResultOk {
		return
	}
	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
	}
}

func ExampleBulk_ReadAsync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	//callback function
	rc := func(resps []datalayer.Response) {
	}

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	r := bulk.ReadAsync(rc, datalayer.BulkReadArg{Address: "address of the node1", Argument: nil}, datalayer.BulkReadArg{Address: "address of the node2"})
	if r != datalayer.ResultOk {
		return
	}
	// waiting for callback
}

func ExampleBulk_Write() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	args := []datalayer.BulkWriteArg{}

	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetFloat32(32.45678)
	args = append(args, datalayer.BulkWriteArg{Address: "address of the node1", Data: v1})

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetFloat64(64.56789)
	args = append(args, datalayer.BulkWriteArg{Address: "address of the node2", Data: v2})

	r := bulk.Write(args...)
	if r != datalayer.ResultOk {
		return
	}
	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
	}
}

func ExampleBulk_WriteAsync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	//callback function
	rc := func(resps []datalayer.Response) {
	}

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	args := []datalayer.BulkWriteArg{}

	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetFloat32(32.45678)
	args = append(args, datalayer.BulkWriteArg{Address: "address of the node1", Data: v1})

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetFloat64(64.56789)
	args = append(args, datalayer.BulkWriteArg{Address: "address of the node2", Data: v2})

	r := bulk.WriteAsync(rc, args...)
	if r != datalayer.ResultOk {
		return
	}
	// waiting for callback
}

func ExampleBulk_Create() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	args := []datalayer.BulkCreateArg{}

	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetFloat32(32.45678)
	args = append(args, datalayer.BulkCreateArg{Address: "address of the node1", Data: v1})

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetFloat64(64.56789)
	args = append(args, datalayer.BulkCreateArg{Address: "address of the node2", Data: v2})

	r := bulk.Create(args...)
	if r != datalayer.ResultOk {
		return
	}
	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
	}
}

func ExampleBulk_CreateAsync() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	//callback function
	rc := func(resps []datalayer.Response) {
	}

	bulk := c.CreateBulk()
	defer datalayer.DeleteBulk(bulk)

	args := []datalayer.BulkCreateArg{}

	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetFloat32(32.45678)
	args = append(args, datalayer.BulkCreateArg{Address: "address of the node1", Data: v1})

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetFloat64(64.56789)
	args = append(args, datalayer.BulkCreateArg{Address: "address of the node2", Data: v2})

	r := bulk.CreateAsync(rc, args...)
	if r != datalayer.ResultOk {
		return
	}
	// waiting for callback
}
