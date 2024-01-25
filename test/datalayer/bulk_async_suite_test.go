/*
 * SPDX-FileCopyrightText: Bosch Rexroth AG
 *
 * SPDX-License-Identifier: MIT
 */
package datalayer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
	"github.com/stretchr/testify/suite"
)

// test suite client bulk async
type ClientBulkAsyncTestSuite struct {
	ClientTestSuite
}

// test entry function
func TestClientBulkAsyncTestSuite(t *testing.T) {
	suite.Run(t, new(ClientBulkAsyncTestSuite))
}

// TestBulkReadAsync - tests an asynchronous bulk read operation
func (suite *ClientBulkAsyncTestSuite) TestBulkReadAsync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	m := []string{
		addressInt1,
		addressInt2,
	}

	done := make(chan bool)

	//callback function
	rc := func(resps []datalayer.Response) {
		//fmt.Println("TestBulkReadAsync: ", len(resps))
		suite.Equal(len(resps), len(m))
		for key, value := range m {
			println(resps[key].Result)
			suite.Equal(resps[key].Address, value)
			suite.Equal(resps[key].Result, datalayer.ResultOk)
			suite.NotNil(resps[key].Data)
			suite.NotNil(resps[key].Time)
		}
		close(done)
	}

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)

	r := bulk.ReadAsync(rc, datalayer.BulkReadArg{Address: m[0], Argument: nil}, datalayer.BulkReadArg{Address: m[1], Argument: nil})
	suite.Equal(r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
}

// TestBulkCreateAsync - tests an asynchronous bulk create operation
func (suite *ClientBulkAsyncTestSuite) TestBulkCreateAsync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan bool)

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)

	args := []datalayer.BulkCreateArg{}
	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetBool8(true)
	args = append(args, datalayer.BulkCreateArg{Address: addressFolder + "mybool", Data: v1})

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetInt32(123)
	args = append(args, datalayer.BulkCreateArg{Address: addressFolder + "myint32", Data: v2})

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkCreateAsync: ", len(resps))
		for key, _ := range resps {
			println(resps[key].Result)
			suite.Equal(resps[key].Result, datalayer.ResultOk)
		}
		close(done)
	}

	r := bulk.CreateAsync(rc, args...)
	suite.Equal(r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
	r = bulk.Delete(addressFolder+"mybool", addressFolder+"myint32")
	suite.Equal(r, datalayer.ResultOk)

}

// TestBulkDeleteAsync - tests an asynchronous bulk delete operation
func (suite *ClientBulkAsyncTestSuite) TestBulkDeleteAsync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	bulkCreateSuite(client, suite.T())

	m := []string{
		addressFolder + "mybool",
		addressFolder + "myint32",
	}

	done := make(chan bool)

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkDeleteAsync: ", len(resps))
		suite.Equal(len(resps), len(m))
		for key, value := range m {
			println(resps[key].Result)
			suite.Equal(resps[key].Address, value)
			suite.Equal(resps[key].Result, datalayer.ResultOk)
			suite.NotNil(resps[key].Data)
			suite.NotNil(resps[key].Time)
		}
		close(done)
	}

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)

	r := bulk.DeleteAsync(rc, m...)
	suite.Equal(r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
}

// TestBulkWriteAsync - tests an asynchronous bulk write operation
func (suite *ClientBulkAsyncTestSuite) TestBulkWriteAsync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan bool)

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)

	args := []datalayer.BulkWriteArg{}
	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetString("Hallo Async")
	arg := datalayer.BulkWriteArg{Address: addressString, Data: v1}
	args = append(args, arg)

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetBool8(false)
	arg = datalayer.BulkWriteArg{Address: addressBool, Data: v2}
	args = append(args, arg)

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkWriteAsync: ", len(resps))
		for key, value := range resps {
			println(resps[key].Result)
			//println(resps[key].Address)
			//println(value.Address)
			suite.Equal(resps[key].Address, value.Address)
			suite.Equal(resps[key].Result, datalayer.ResultOk)
			suite.NotNil(resps[key].Data)
			suite.NotNil(resps[key].Time)
		}
		close(done)
	}

	r := bulk.WriteAsync(rc, args...)
	suite.Equal(r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
}

// TestBulkBrowseAsync - tests an asynchronous bulk browse operation
func (suite *ClientBulkAsyncTestSuite) TestBulkBrowseAsync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan bool)

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkBrowseAsync: ", len(resps))
		for key, _ := range resps {
			println(resps[key].Result)
			suite.Equal(resps[key].Result, datalayer.ResultOk)
			data := resps[key].Data
			suite.NotNil(data)
			for _, v := range data.GetArrayString() {
				println(v)
			}
		}
		close(done)
	}

	r := bulk.BrowseAsync(rc, addressbase, "types")
	suite.Equal(r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
}

// TestBulkMetadataAsync - tests an asynchronous bulk read metadata operation
func (suite *ClientBulkAsyncTestSuite) TestBulkMetadataAsync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan bool)

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkMetadataAsync: ", len(resps))
		for key, _ := range resps {
			println("Result:", resps[key].Result)
			suite.Equal(resps[key].Result, datalayer.ResultOk)
			data := resps[key].Data
			suite.NotNil(data)
			suite.Equal(data.GetType(), datalayer.VariantTypeFlatbuffers)
			suite.NotNil(data.GetFlatbuffers())
			metadata := fbs.GetRootAsMetadata(data.GetFlatbuffers(), 0)
			println(metadata.NodeClass())
		}
		close(done)
	}

	r := bulk.MetadataAsync(rc, addressString, addressBool)
	suite.Equal(r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
}
