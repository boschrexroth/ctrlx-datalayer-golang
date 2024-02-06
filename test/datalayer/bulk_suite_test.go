/*
 * SPDX-FileCopyrightText: Bosch Rexroth AG
 *
 * SPDX-License-Identifier: MIT
 */
package datalayer_test

import (
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// test suite client bulk
type ClientBulkTestSuite struct {
	ClientTestSuite
}

// test entry function
func TestClientBulkTestSuite(t *testing.T) {
	suite.Run(t, new(ClientBulkTestSuite))
}

// TestBulk - simple test
func (suite *ClientBulkTestSuite) TestBulk() {

	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)
}

// TestBulkRead - tests a bulk read operation
func (suite *ClientBulkTestSuite) TestBulkRead() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)
	m := []string{
		addressInt1,
		addressInt2,
	}

	args := []datalayer.BulkReadArg{}
	for _, value := range m {
		arg := datalayer.BulkReadArg{Address: value, Argument: nil}
		args = append(args, arg)
	}

	r := bulk.Read(args...)
	suite.Equal(r, datalayer.ResultOk)
	suite.Equal(len(bulk.Response), 2)

	for key, value := range m {
		println(bulk.Response[key].Result)
		suite.Equal(bulk.Response[key].Address, value)
		suite.Equal(bulk.Response[key].Result, datalayer.ResultOk)
		data := bulk.Response[key].Data
		suite.NotNil(data)
		println(data.GetInt32())
		suite.NotNil(bulk.Response[key].Time)
	}
}

// Prerequisite: sdk-cpp-alldata snap is installed
func bulkCreateSuite(client *datalayer.Client, t *testing.T) {
	bulk := client.CreateBulk()
	assert.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	args := []datalayer.BulkCreateArg{}
	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetBool8(false)
	args = append(args, datalayer.BulkCreateArg{Address: addressFolder + "mybool", Data: v1})

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetInt32(314)
	args = append(args, datalayer.BulkCreateArg{Address: addressFolder + "myint32", Data: v2})

	r := bulk.Create(args...)
	assert.Equal(t, r, datalayer.ResultOk)
	assert.Equal(t, len(bulk.Response), 2)

	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
		assert.Equal(t, bulk.Response[key].Result, datalayer.ResultOk)
	}
}

// TestBulkDelete - tests a bulk delete operation
// Prerequisites:
// - TestBulkCreate was running before
func (suite *ClientBulkTestSuite) TestBulkCreateDelete() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	bulkCreateSuite(client, suite.T())

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)
	m := []string{
		addressFolder + "mybool",
		addressFolder + "myint32",
	}

	r := bulk.Delete(m...)
	suite.Equal(r, datalayer.ResultOk)
	suite.Equal(len(bulk.Response), 2)

	for key, value := range m {
		println(bulk.Response[key].Result)
		suite.Equal(bulk.Response[key].Address, value)
		suite.Equal(bulk.Response[key].Result, datalayer.ResultOk)
		suite.NotNil(bulk.Response[key].Data)
		suite.NotNil(bulk.Response[key].Time)
	}
}

// TestBulkWrite - tests a bulk write operation
func (suite *ClientBulkTestSuite) TestBulkWrite() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)

	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetString("hallo")

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetBool8(true)

	r := bulk.Write(datalayer.BulkWriteArg{Address: addressString, Data: v1},
		datalayer.BulkWriteArg{Address: addressBool, Data: v2})
	suite.Equal(r, datalayer.ResultOk)
	suite.Equal(len(bulk.Response), 2)

	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
		suite.Equal(bulk.Response[key].Result, datalayer.ResultOk)
		suite.NotNil(bulk.Response[key].Data)
		suite.NotNil(bulk.Response[key].Time)
	}
}

// TestBulkBrowse - tests a bulk browse operation
func (suite *ClientBulkTestSuite) TestBulkBrowse() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)

	r := bulk.Browse(addressbase, "types")
	suite.Equal(r, datalayer.ResultOk)
	suite.Equal(len(bulk.Response), 2)

	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
		suite.Equal(bulk.Response[key].Result, datalayer.ResultOk)
		data := bulk.Response[key].Data
		suite.NotNil(data)
		for _, v := range data.GetArrayString() {
			println(v)
		}
	}
}

// TestBulkMetadata - tests a bulk read metadata operation
func (suite *ClientBulkTestSuite) TestBulkMetadata() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	bulk := client.CreateBulk()
	suite.NotNil(bulk)
	defer datalayer.DeleteBulk(bulk)

	r := bulk.Metadata(addressString, addressBool)
	suite.Equal(r, datalayer.ResultOk)
	suite.Equal(len(bulk.Response), 2)

	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
		suite.Equal(bulk.Response[key].Result, datalayer.ResultOk)
		data := bulk.Response[key].Data
		suite.NotNil(data)
		suite.Equal(data.GetType(), datalayer.VariantTypeFlatbuffers)
		suite.NotNil(data.GetFlatbuffers())
		metadata := fbs.GetRootAsMetadata(data.GetFlatbuffers(), 0)
		println(metadata.NodeClass())
	}
}
