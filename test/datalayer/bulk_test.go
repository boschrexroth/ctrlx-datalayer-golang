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
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
	a "github.com/stretchr/testify/assert"
)

func init() {
	clientAddress = ctrlxAddress()
	alldataprovider = allDataProvider()
}

func TestBulk(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)
}

// TestBulkRead - tests a bulk read operation
func TestBulkRead(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)
	m := []string{
		"framework/metrics/system/cpu-utilisation-percent",
		"framework/metrics/system/memavailable-mb",
	}

	args := []datalayer.BulkReadArg{}
	for _, value := range m {
		arg := datalayer.BulkReadArg{Address: value, Argument: nil}
		args = append(args, arg)
	}

	r := bulk.Read(args...)
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, len(bulk.Response), 2)

	for key, value := range m {
		println(bulk.Response[key].Result)
		a.Equal(t, bulk.Response[key].Address, value)
		a.Equal(t, bulk.Response[key].Result, datalayer.ResultOk)
		data := bulk.Response[key].Data
		a.NotNil(t, data)
		println(data.GetFloat64())
		a.NotNil(t, bulk.Response[key].Time)
	}
}

// Prerequisite: sdk-cpp-alldata snap is installed
func bulkCreate(t *testing.T, client *datalayer.Client) {
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

	r := bulk.Create(args...)
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, len(bulk.Response), 2)

	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
		a.Equal(t, bulk.Response[key].Result, datalayer.ResultOk)
	}
}

// TestBulkDelete - tests a bulk delete operation
// Prerequisites:
// - sdk-cpp-alldata snap is installed
// - TestBulkCreate was running before
func TestBulkCreateDelete(t *testing.T) {
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

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)
	m := []string{
		"sdk-cpp-alldata/dynamic/float32x",
		"sdk-cpp-alldata/dynamic/float64x",
	}

	r := bulk.Delete(m...)
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, len(bulk.Response), 2)

	for key, value := range m {
		println(bulk.Response[key].Result)
		a.Equal(t, bulk.Response[key].Address, value)
		a.Equal(t, bulk.Response[key].Result, datalayer.ResultOk)
		a.NotNil(t, bulk.Response[key].Data)
		a.NotNil(t, bulk.Response[key].Time)
	}
}

// TestBulkWrite - tests a bulk write operation
// Prerequisite: sdk-cpp-alldata snap is installed
func TestBulkWrite(t *testing.T) {
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

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	v1 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v1)
	v1.SetFloat32(32.45678)

	v2 := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v2)
	v2.SetFloat64(64.56789)

	r := bulk.Write(datalayer.BulkWriteArg{Address: "sdk-cpp-alldata/dynamic/float32", Data: v1},
		datalayer.BulkWriteArg{Address: "sdk-cpp-alldata/dynamic/float64", Data: v2})
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, len(bulk.Response), 2)

	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
		a.Equal(t, bulk.Response[key].Result, datalayer.ResultOk)
		a.NotNil(t, bulk.Response[key].Data)
		a.NotNil(t, bulk.Response[key].Time)
	}
}

// TestBulkBrowse - tests a bulk browse operation
func TestBulkBrowse(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	r := bulk.Browse("framework", "types")
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, len(bulk.Response), 2)

	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
		a.Equal(t, bulk.Response[key].Result, datalayer.ResultOk)
		data := bulk.Response[key].Data
		a.NotNil(t, data)
		for _, v := range data.GetArrayString() {
			println(v)
		}
	}
}

// TestBulkMetadata - tests a bulk read metadata operation
func TestBulkMetadata(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	bulk := client.CreateBulk()
	a.NotNil(t, bulk)
	defer datalayer.DeleteBulk(bulk)

	r := bulk.Metadata("framework/metrics/system/cpu-utilisation-percent", "framework/metrics/system/memavailable-mb")
	a.Equal(t, r, datalayer.ResultOk)
	a.Equal(t, len(bulk.Response), 2)

	for key, _ := range bulk.Response {
		println(bulk.Response[key].Result)
		a.Equal(t, bulk.Response[key].Result, datalayer.ResultOk)
		data := bulk.Response[key].Data
		a.NotNil(t, data)
		a.Equal(t, data.GetType(), datalayer.VariantTypeFlatbuffers)
		a.NotNil(t, data.GetFlatbuffers())
		metadata := fbs.GetRootAsMetadata(data.GetFlatbuffers(), 0)
		println(metadata.NodeClass())
	}
}
