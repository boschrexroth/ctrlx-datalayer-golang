/**
 * MIT License
 *
 * Copyright (c) 2022 Bosch Rexroth AG
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package datalayer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
	a "github.com/stretchr/testify/assert"
)

var alldataprovider string = ""

func init() {
	clientAddress = ctrlxAddress()
	alldataprovider = allDataProvider()
}

// TestBulkReadAsync - tests an asynchronous bulk read operation
func TestBulkReadAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}

	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	m := []string{
		"/framework/metrics/system/cpu-utilisation-percent",
		"/diagnosis/cfg/realtime/numbers",
	}

	done := make(chan bool)

	//callback function
	rc := func(resps []datalayer.Response) {
		//fmt.Println("TestBulkReadAsync: ", len(resps))
		a.Equal(t, len(resps), len(m))
		for key, value := range m {
			println(resps[key].Result)
			a.Equal(t, resps[key].Address, value)
			a.Equal(t, resps[key].Result, datalayer.ResultOk)
			a.NotNil(t, resps[key].Data)
			a.NotNil(t, resps[key].Time)
		}
		close(done)
	}

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	r := bulk.ReadAsync(rc, datalayer.BulkReadArg{Address: m[0], Argument: nil}, datalayer.BulkReadArg{Address: m[1], Argument: nil})
	a.Equal(t, r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
}

// TestBulkCreateAsync - tests an asynchronous bulk create operation
// Prerequisite: Snap sdk-cpp-alldata has to be installed
func TestBulkCreateAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	if alldataprovider == "" {
		t.Skip("all data provider does not exist")
	}

	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan bool)

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	args := []datalayer.BulkCreateArg{}
	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetFloat32(32.45678)
	args = append(args, datalayer.BulkCreateArg{Address: "sdk-cpp-alldata/dynamic/float32x", Data: v1})

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetFloat64(64.56789)
	args = append(args, datalayer.BulkCreateArg{Address: "sdk-cpp-alldata/dynamic/float64x", Data: v2})

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkCreateAsync: ", len(resps))
		for key, _ := range resps {
			println(resps[key].Result)
			a.Equal(t, resps[key].Result, datalayer.ResultOk)
		}
		close(done)
	}

	r := bulk.CreateAsync(rc, args...)
	a.Equal(t, r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
	r = bulk.Delete("sdk-cpp-alldata/dynamic/float32x", "sdk-cpp-alldata/dynamic/float64x")
	a.Equal(t, r, datalayer.ResultOk)

}

// TestBulkDeleteAsync - tests an asynchronous bulk delete operation
// Prerequisites:
// - Snap sdk-cpp-alldata has to be installed
func TestBulkDeleteAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	if alldataprovider == "" {
		t.Skip("all data provider does not exist")
	}

	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	bulkCreate(t, client)

	m := []string{
		"sdk-cpp-alldata/dynamic/float32x",
		"sdk-cpp-alldata/dynamic/float64x",
	}

	done := make(chan bool)

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkDeleteAsync: ", len(resps))
		a.Equal(t, len(resps), len(m))
		for key, value := range m {
			println(resps[key].Result)
			a.Equal(t, resps[key].Address, value)
			a.Equal(t, resps[key].Result, datalayer.ResultOk)
			a.NotNil(t, resps[key].Data)
			a.NotNil(t, resps[key].Time)
		}
		close(done)
	}

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	r := bulk.DeleteAsync(rc, m...)
	a.Equal(t, r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
}

// TestBulkWriteAsync - tests an asynchronous bulk write operation
func TestBulkWriteAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	if alldataprovider == "" {
		t.Skip("all data provider does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan bool)

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	args := []datalayer.BulkWriteArg{}
	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetFloat32(32.45678)
	arg := datalayer.BulkWriteArg{Address: "sdk-cpp-alldata/dynamic/float32", Data: v1}
	args = append(args, arg)

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetFloat64(64.56789)
	arg = datalayer.BulkWriteArg{Address: "sdk-cpp-alldata/dynamic/float64", Data: v2}
	args = append(args, arg)

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkWriteAsync: ", len(resps))
		for key, value := range resps {
			println(resps[key].Result)
			//println(resps[key].Address)
			//println(value.Address)
			a.Equal(t, resps[key].Address, value.Address)
			a.Equal(t, resps[key].Result, datalayer.ResultOk)
			a.NotNil(t, resps[key].Data)
			a.NotNil(t, resps[key].Time)
		}
		close(done)
	}

	r := bulk.WriteAsync(rc, args...)
	a.Equal(t, r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
}

// TestBulkBrowseAsync - tests an asynchronous bulk browse operation
func TestBulkBrowseAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan bool)

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkBrowseAsync: ", len(resps))
		for key, _ := range resps {
			println(resps[key].Result)
			a.Equal(t, resps[key].Result, datalayer.ResultOk)
			data := resps[key].Data
			a.NotNil(t, data)
			for _, v := range data.GetArrayString() {
				println(v)
			}
		}
		close(done)
	}

	r := bulk.BrowseAsync(rc, "framework", "types")
	a.Equal(t, r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
}

// TestBulkMetadataAsync - tests an asynchronous bulk read metadata operation
func TestBulkMetadataAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan bool)

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	//callback function
	rc := func(resps []datalayer.Response) {
		fmt.Println("TestBulkMetadataAsync: ", len(resps))
		for key, _ := range resps {
			println("Result:", resps[key].Result)
			a.Equal(t, resps[key].Result, datalayer.ResultOk)
			data := resps[key].Data
			a.NotNil(t, data)
			a.Equal(t, data.GetType(), datalayer.VariantTypeFlatbuffers)
			a.NotNil(t, data.GetFlatbuffers())
			metadata := fbs.GetRootAsMetadata(data.GetFlatbuffers(), 0)
			println(metadata.NodeClass())
		}
		close(done)
	}

	r := bulk.MetadataAsync(rc, "framework/metrics/system/cpu-utilisation-percent", "framework/metrics/system/memavailable-mb")
	a.Equal(t, r, datalayer.ResultOk)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
}
