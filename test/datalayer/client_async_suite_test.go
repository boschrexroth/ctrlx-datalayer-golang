/*
 * SPDX-FileCopyrightText: Bosch Rexroth AG
 *
 * SPDX-License-Identifier: MIT
 */
package datalayer_test

import (
	"sync"
	"testing"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	"github.com/stretchr/testify/suite"
)

// test suite client async
type ClientAsyncTestSuite struct {
	ClientTestSuite
}

// test entry function
func TestClientAsyncTestSuite(t *testing.T) {
	suite.Run(t, new(ClientAsyncTestSuite))
}

func (suite *ClientAsyncTestSuite) TestClientPingAsync() {
	client := suite.initClient()

	defer datalayer.DeleteClient(client)

	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			r := client.PingAsync(func(res datalayer.Result, v *datalayer.Variant) {
				suite.Equal(res, datalayer.Result(0))
				wg.Done()
			})
			suite.Equal(datalayer.Result(0), r)
		}()
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(timeoutStressTest()):
		suite.Fail("callback timeout")
	}
}

func (suite *ClientAsyncTestSuite) TestClientReadAsync() {
	client := suite.initClient()

	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		//t.Log("callback: ", res)
		suite.Equal(res, datalayer.Result(0))
		val := datalayer.NewVariant()
		suite.NotPanics(func() { v.Copy(val) })
		done <- callbackvalue{res: res, val: val}
	}

	args := datalayer.NewVariant()
	defer datalayer.DeleteVariant(args)

	r := client.ReadAsync(addressBool, args, rc)
	suite.Equal(datalayer.Result(0), r)

	select {
	case vals := <-done:
		suite.Equal(vals.res, datalayer.Result(0))
		suite.Equal(vals.val.GetType(), datalayer.VariantTypeBool8)
		datalayer.DeleteVariant(vals.val)
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
	close(done)
}

func (suite *ClientAsyncTestSuite) TestClientWriteAsync() {
	client := suite.initClient()

	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan bool)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		//t.Log("callback: ", res)
		suite.Equal(res, datalayer.Result(0))
		close(done)
	}

	data := datalayer.NewVariant()
	defer datalayer.DeleteVariant(data)
	data.SetBool8(false)

	r := client.WriteAsync(addressBool, data, rc)
	suite.Equal(datalayer.Result(0), r)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
}

func (suite *ClientAsyncTestSuite) TestClientBrowseAsync() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		//t.Log("callback: ", res)
		suite.Equal(res, datalayer.Result(0))
		val := datalayer.NewVariant()
		suite.NotPanics(func() { v.Copy(val) })
		done <- callbackvalue{res: res, val: val}
	}

	r := client.BrowseAsync("", rc)
	suite.Equal(datalayer.Result(0), r)

	select {
	case vals := <-done:
		suite.Equal(vals.res, datalayer.Result(0))
		suite.Equal(vals.val.GetType(), datalayer.VariantTypeArrayString)
		datalayer.DeleteVariant(vals.val)
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
	close(done)
}

func (suite *ClientTestSuite) TestClientMetadataAsync() {
	client := suite.initClient()

	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		//t.Log("callback: ", res)
		suite.Equal(res, datalayer.Result(0))
		val := datalayer.NewVariant()
		suite.NotPanics(func() { v.Copy(val) })
		done <- callbackvalue{res: res, val: val}
	}

	r := client.MetadataAsync(addressBool, rc)
	suite.Equal(datalayer.Result(0), r)

	select {
	case vals := <-done:
		suite.Equal(vals.res, datalayer.Result(0))
		suite.Equal(vals.val.GetType(), datalayer.VariantTypeFlatbuffers)
		datalayer.DeleteVariant(vals.val)
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
	close(done)
}

func (suite *ClientTestSuite) TestClientCreateAsync() {
	client := suite.initClient()

	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		done <- callbackvalue{res: res, val: nil}
	}

	args := datalayer.NewVariant()
	defer datalayer.DeleteVariant(args)

	r := client.CreateAsync("", args, rc)
	suite.Equal(datalayer.Result(0), r)

	select {
	case vals := <-done:
		suite.Equal(vals.res, datalayer.Result(0x80010001)) // DL_INVALID_ADDRESS
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
	close(done)
}

func (suite *ClientTestSuite) TestClientRemoveAsync() {
	client := suite.initClient()

	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		done <- callbackvalue{res: res, val: nil}
	}

	args := datalayer.NewVariant()
	defer datalayer.DeleteVariant(args)

	r := client.RemoveAsync("", rc)
	suite.Equal(datalayer.Result(0), r)

	select {
	case vals := <-done:
		suite.Equal(vals.res, datalayer.Result(0x80010001)) // DL_INVALID_ADDRESS
	case <-time.After(asyncTestTimeout()):
		suite.Fail("callback timeout")
	}
	close(done)
}
