/*
 * SPDX-FileCopyrightText: Bosch Rexroth AG
 *
 * SPDX-License-Identifier: MIT
 */

package datalayer_test

import (
	"fmt"
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	"github.com/stretchr/testify/suite"
)

// test suite client
type ClientTestSuite struct {
	suite.Suite
	tp *testProvider
}

// test entry function
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) TestClientBasics() {

	system := suite.tp.system
	suite.NotNil(system)

	client := system.Factory().CreateClient("ipc://")
	defer datalayer.DeleteClient(client)
	suite.NotNil(client)

	suite.Equal(datalayer.Result(0), client.SetTimeout(datalayer.TimeoutSettingPing, ctrlxClientTimeout()))
	suite.True(client.IsConnected())
	suite.Equal(datalayer.Result(0), client.PingSync())

	at := client.GetAuthToken()
	suite.True(len(at) == 0)
}

func (suite *ClientTestSuite) TestClientBrowseSync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.BrowseSync("")
	defer datalayer.DeleteVariant(v)
	suite.Equal(datalayer.Result(0), r)
	suite.True(v.GetType() == datalayer.VariantTypeArrayString)
}

func (suite *ClientTestSuite) TestClientBrowseSyncError() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.BrowseSync("not_exists/")
	defer datalayer.DeleteVariant(v)
	suite.NotEqual(datalayer.Result(0), r)
}

func (suite *ClientTestSuite) TestClientBrowse() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.Browse("")
	suite.Equal(datalayer.Result(0), r)
	suite.True(len(v) != 0)
	fmt.Println("len(Browse):", len(v))
	fmt.Println("Browse[0]:", v[0])
}

func (suite *ClientTestSuite) TestClientBrowseError() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.Browse("not_exists/")
	suite.NotEqual(datalayer.Result(0), r)
	suite.True(len(v) == 0)
}

func (suite *ClientTestSuite) TestClientReadSyncArgs() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	v := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v)
	r := client.ReadSyncArgs(addressString, v)
	suite.Equal(datalayer.Result(0), r, clientAddress)
	suite.True(v.GetType() == datalayer.VariantTypeString)
}

func (suite *ClientTestSuite) TestClientReadSync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.ReadSync(addressString)
	defer datalayer.DeleteVariant(v)
	suite.Equal(datalayer.Result(0), r, clientAddress)
	suite.True(v.GetType() == datalayer.VariantTypeString)
}

func (suite *ClientTestSuite) TestClientReadSyncError() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.ReadSync(addressbase + "not_exists")
	defer datalayer.DeleteVariant(v)
	suite.NotEqual(datalayer.Result(0), r, clientAddress)
}

func (suite *ClientTestSuite) TestClientWriteSync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	v := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v)
	v.SetBool8(true)
	r := client.WriteSync(addressBool, v)
	suite.Equal(datalayer.Result(0), r)
	suite.True(v.GetType() == datalayer.VariantTypeBool8)
}

func (suite *ClientTestSuite) TestClientMetadataSync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.MetadataSync(addressBool)
	defer datalayer.DeleteVariant(v)
	suite.Equal(datalayer.Result(0), r)
	suite.True(v.GetType() == datalayer.VariantTypeFlatbuffers)
}

func (suite *ClientTestSuite) TestClientMetadataSyncError() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.MetadataSync(addressbase + "not_exists")
	defer datalayer.DeleteVariant(v)
	suite.NotEqual(datalayer.Result(0), r)
}

func (suite *ClientTestSuite) TestClientMetadata() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.Metadata(addressBool)
	suite.Equal(datalayer.Result(0), r)
	suite.NotNil(v)
}

func (suite *ClientTestSuite) TestClientMetadataError() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	r, v := client.Metadata(addressbase + "not_exists")
	suite.NotEqual(datalayer.Result(0), r)
	suite.Nil(v)
}

func (suite *ClientTestSuite) TestCreateSync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.NotNil(client)

	d := datalayer.NewVariant()
	defer datalayer.DeleteVariant(d)
	r := client.CreateSync("", d)
	suite.Equal(datalayer.ResultInvalidAddress, r) //0x80010001 == DL_INVALID_ADDRESS
}

func (suite *ClientTestSuite) TestRemoveSync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.NotNil(client)

	r := client.RemoveSync("")
	suite.Equal(datalayer.ResultInvalidAddress, r) //0x80010001 == DL_INVALID_ADDRESS
}

func (suite *ClientTestSuite) TestReadJsonSync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.NotNil(client)

	cv := suite.tp.system.JSONConverter()
	suite.NotNil(cv)

	r, d := client.ReadJsonSync(cv, addressString, 2)
	suite.Equal(datalayer.Result(0), r)
	suite.NotNil(d)
	suite.True(len(d) != 0)

	r, d = client.ReadJsonSyncArgs(cv, addressString, 2, []byte("{\"type\": \"string\", \"value\": \"testdata\"}"))
	suite.Equal(datalayer.Result(0), r)
	suite.NotNil(d)
	suite.True(len(d) != 0)
}

func (suite *ClientTestSuite) TestWriteJsonSync() {
	client := suite.initClient()

	defer datalayer.DeleteClient(client)
	suite.NotNil(client)

	cv := suite.tp.system.JSONConverter()
	suite.NotNil(cv)

	r, d := client.ReadJsonSync(cv, addressString, 2)
	suite.Equal(datalayer.Result(0), r)

	r, err := client.WriteJsonSync(cv, addressString, d)
	suite.Equal(datalayer.Result(0), r)
	suite.Nil(err)
}

func (suite *ClientTestSuite) TestCreateSyncAddr() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.NotNil(client)

	d := datalayer.NewVariant()
	defer datalayer.DeleteVariant(d)
	d.SetBool8(false)
	r := client.CreateSync(addressFolder+"mybool", d)
	suite.Equal(datalayer.ResultOk, r) //0x80010001 == DL_INVALID_ADDRESS
}
