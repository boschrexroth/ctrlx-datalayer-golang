/**
 * MIT License
 *
 * Copyright (c) 2021-2022 Bosch Rexroth AG
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

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	a "github.com/stretchr/testify/assert"
)

var clientAddress string = ""

func init() {
	clientAddress = ctrlxAddress()
}

func TestDeleteClient(t *testing.T) {
	a.NotPanics(t, func() { datalayer.DeleteClient(nil) })
}

func TestClientBasics(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(false) })

	client := system.Factory().CreateClient(clientAddress)
	defer datalayer.DeleteClient(client)
	a.NotNil(t, client)

	a.Equal(t, datalayer.Result(0), client.SetTimeout(datalayer.TimeoutSettingPing, ctrlxClientTimeout()))
	a.True(t, client.IsConnected())
	a.Equal(t, datalayer.Result(0), client.PingSync())

	at := client.GetAuthToken()
	a.True(t, len(at) != 0)
}

func initClient() (*datalayer.System, *datalayer.Client) {
	system := datalayer.NewSystem("")
	system.Start(false)

	client := system.Factory().CreateClient(clientAddress)
	client.SetTimeout(datalayer.TimeoutSettingPing, ctrlxClientTimeout())
	return system, client
}
func TestClientBrowseSync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.BrowseSync("")
	defer datalayer.DeleteVariant(v)
	a.Equal(t, datalayer.Result(0), r)
	a.True(t, v.GetType() == datalayer.VariantTypeArrayString)
}

func TestClientBrowseSyncError(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.BrowseSync("not_exists/")
	defer datalayer.DeleteVariant(v)
	a.NotEqual(t, datalayer.Result(0), r)
}

func TestClientBrowse(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.Browse("")
	a.Equal(t, datalayer.Result(0), r)
	a.True(t, len(v) != 0)
}

func TestClientBrowseError(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.Browse("not_exists/")
	a.NotEqual(t, datalayer.Result(0), r)
	a.True(t, len(v) == 0)
}
func TestClientReadSyncArgs(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	v := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v)
	r := client.ReadSyncArgs("/diagnosis/cfg/realtime/numbers", v)
	a.Equal(t, datalayer.Result(0), r, clientAddress)
	a.True(t, v.GetType() == datalayer.VariantTypeBool8)
}

func TestClientReadSync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.ReadSync("/diagnosis/cfg/realtime/numbers")
	defer datalayer.DeleteVariant(v)
	a.Equal(t, datalayer.Result(0), r, clientAddress)
	a.True(t, v.GetType() == datalayer.VariantTypeBool8)
}

func TestClientReadSyncError(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.ReadSync("/diagnosis/cfg/realtime/not_exists")
	defer datalayer.DeleteVariant(v)
	a.NotEqual(t, datalayer.Result(0), r, clientAddress)
}

func TestClientWriteSync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	v := datalayer.NewVariant()
	defer datalayer.DeleteVariant(v)
	v.SetBool8(true)
	r := client.WriteSync("/diagnosis/cfg/realtime/numbers", v)
	a.Equal(t, datalayer.Result(0), r)
	a.True(t, v.GetType() == datalayer.VariantTypeUnknown)
}

func TestClientMetadataSync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.MetadataSync("/diagnosis/cfg/realtime/numbers")
	defer datalayer.DeleteVariant(v)
	a.Equal(t, datalayer.Result(0), r)
	a.True(t, v.GetType() == datalayer.VariantTypeFlatbuffers)
}

func TestClientMetadataSyncError(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.MetadataSync("/diagnosis/cfg/realtime/not_exists")
	defer datalayer.DeleteVariant(v)
	a.NotEqual(t, datalayer.Result(0), r)
}

func TestClientMetadata(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.Metadata("/diagnosis/cfg/realtime/numbers")
	a.Equal(t, datalayer.Result(0), r)
	a.NotNil(t, v)
}

func TestClientMetadataError(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	r, v := client.Metadata("/diagnosis/cfg/realtime/not_exists")
	a.NotEqual(t, datalayer.Result(0), r)
	a.Nil(t, v)
}

func TestAuthToken(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}

	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())
	s := client.GetAuthToken()
	a.NotNil(t, s)
	a.True(t, len(s) != 0)
	a.NotPanics(t, func() { client.SetAuthToken(s) })
}

func TestCreateSync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(false) })

	client := system.Factory().CreateClient(clientAddress)
	defer datalayer.DeleteClient(client)
	a.NotNil(t, client)

	a.Equal(t, datalayer.Result(0), client.SetTimeout(datalayer.TimeoutSettingPing, ctrlxClientTimeout()))
	a.True(t, client.IsConnected())

	d := datalayer.NewVariant()
	defer datalayer.DeleteVariant(d)
	r := client.CreateSync("", d)
	a.Equal(t, datalayer.Result(0x80010001), r) //0x80010001 == DL_INVALID_ADDRESS
}

func TestRemoveSync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(false) })

	client := system.Factory().CreateClient(clientAddress)
	defer datalayer.DeleteClient(client)
	a.NotNil(t, client)

	a.Equal(t, datalayer.Result(0), client.SetTimeout(datalayer.TimeoutSettingPing, ctrlxClientTimeout()))
	a.True(t, client.IsConnected())

	r := client.RemoveSync("")
	a.Equal(t, datalayer.Result(0x80010001), r) //0x80010001 == DL_INVALID_ADDRESS
}

func TestReadJsonSync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(false) })

	client := system.Factory().CreateClient(clientAddress)
	defer datalayer.DeleteClient(client)
	a.NotNil(t, client)

	a.Equal(t, datalayer.Result(0), client.SetTimeout(datalayer.TimeoutSettingPing, ctrlxClientTimeout()))
	a.True(t, client.IsConnected())

	cv := system.JSONConverter()
	a.NotNil(t, cv)

	r, d := client.ReadJsonSync(cv, "scheduler/admin/state", 2)
	a.Equal(t, datalayer.Result(0), r)
	a.NotNil(t, d)
	a.True(t, len(d) != 0)
}

func TestWriteJsonSync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(false) })

	client := system.Factory().CreateClient(clientAddress)
	defer datalayer.DeleteClient(client)
	a.NotNil(t, client)

	a.Equal(t, datalayer.Result(0), client.SetTimeout(datalayer.TimeoutSettingPing, ctrlxClientTimeout()))
	a.True(t, client.IsConnected())

	cv := system.JSONConverter()
	a.NotNil(t, cv)

	r, d := client.ReadJsonSync(cv, "scheduler/admin/state", 2)
	a.Equal(t, datalayer.Result(0), r)

	r, err := client.WriteJsonSync(cv, "scheduler/admin/state", d)
	a.Equal(t, datalayer.Result(0), r)
	a.Nil(t, err)
}
