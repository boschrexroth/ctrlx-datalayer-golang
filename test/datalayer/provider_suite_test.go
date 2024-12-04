/*
 * SPDX-FileCopyrightText: Bosch Rexroth AG
 *
 * SPDX-License-Identifier: MIT
 */

package datalayer_test

import (
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	"github.com/stretchr/testify/suite"
)

// test suite client
type ProviderTestSuite struct {
	ClientTestSuite
}

// test entry function
func TestProviderTestSuite(t *testing.T) {
	suite.Run(t, new(ProviderTestSuite))
}

func (suite *ProviderTestSuite) TestGetRegisteredType() {
	r, d := suite.tp.provider.GetRegisteredType("identity") //see test "c_provider_test.cpp"

	defer datalayer.DeleteVariant(d)
	suite.Equal(datalayer.ResultOk, r)
	suite.NotNil(d)
	suite.Equal(d.GetType(), datalayer.VariantTypeArrayUint8)

	d.SetBool8(true)
	r = suite.tp.provider.RegisterTypeVariant("types/GetRegisteredType/test1", d)
	suite.Equal(datalayer.ResultOk, r)

	r, d1 := suite.tp.provider.GetRegisteredType("types/GetRegisteredType/test1") //see test "c_provider_test.cpp"
	defer datalayer.DeleteVariant(d1)
	suite.Equal(datalayer.ResultOk, r)
	suite.NotNil(d1)
	suite.Equal(d1.GetType(), datalayer.VariantTypeBool8)
}

func (suite *ProviderTestSuite) TestGetRegisteredNodePaths() {
	r, d := suite.tp.provider.GetRegisteredNodePaths()
	suite.Equal(datalayer.ResultOk, r)
	suite.NotNil(d)
	suite.True(len(d) >= 2)
}

func (suite *ProviderTestSuite) TestGetRejectedNodePaths() {
	r, d := suite.tp.provider.GetRejectedNodePaths()
	suite.Equal(datalayer.ResultOk, r)
	suite.NotNil(d)
	suite.True(len(d) >= 0)
}

func (suite *ProviderTestSuite) TestPublishEvent() {
	data := datalayer.NewVariant()
	defer datalayer.DeleteVariant(data)
	r := suite.tp.provider.PublishEvent(data, data)
	suite.Equal(datalayer.ResultTypeMismatch, r)
}
