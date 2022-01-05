/**
 * MIT License
 *
 * Copyright (c) 2021 Bosch Rexroth AG
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
	"sync"
	"testing"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	a "github.com/stretchr/testify/assert"
)

func init() {
	clientAddress = ctrlxAddress()
}

type callbackvalue struct {
	res datalayer.Result
	val *datalayer.Variant
}

func timeoutStressTest() time.Duration {
	return asyncTestTimeout() * 5
}

func TestClientPingAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			r := client.PingAsync(func(res datalayer.Result, v *datalayer.Variant) {
				a.Equal(t, res, datalayer.Result(0))
				wg.Done()
			})
			a.Equal(t, datalayer.Result(0), r)
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
		a.Fail(t, "callback timeout")
	}
}

func TestClientReadAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		//t.Log("callback: ", res)
		a.Equal(t, res, datalayer.Result(0))
		val := datalayer.NewVariant()
		a.NotPanics(t, func() { v.Copy(val) })
		done <- callbackvalue{res: res, val: val}
	}

	args := datalayer.NewVariant()
	defer datalayer.DeleteVariant(args)

	r := client.ReadAsync("/diagnosis/cfg/realtime/numbers", args, rc)
	a.Equal(t, datalayer.Result(0), r)

	select {
	case vals := <-done:
		a.Equal(t, vals.res, datalayer.Result(0))
		a.Equal(t, vals.val.GetType(), datalayer.VariantTypeBool8)
		datalayer.DeleteVariant(vals.val)
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
	close(done)
}

func TestClientWriteAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan bool)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		//t.Log("callback: ", res)
		a.Equal(t, res, datalayer.Result(0))
		close(done)
	}

	data := datalayer.NewVariant()
	defer datalayer.DeleteVariant(data)
	data.SetBool8(false)

	r := client.WriteAsync("/diagnosis/cfg/realtime/numbers", data, rc)
	a.Equal(t, datalayer.Result(0), r)

	select {
	case <-done:
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
}

func TestClientBrowseAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		//t.Log("callback: ", res)
		a.Equal(t, res, datalayer.Result(0))
		val := datalayer.NewVariant()
		a.NotPanics(t, func() { v.Copy(val) })
		done <- callbackvalue{res: res, val: val}
	}

	r := client.BrowseAsync("", rc)
	a.Equal(t, datalayer.Result(0), r)

	select {
	case vals := <-done:
		a.Equal(t, vals.res, datalayer.Result(0))
		a.Equal(t, vals.val.GetType(), datalayer.VariantTypeArrayString)
		datalayer.DeleteVariant(vals.val)
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
	close(done)
}

func TestClientMetadataAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		//t.Log("callback: ", res)
		a.Equal(t, res, datalayer.Result(0))
		val := datalayer.NewVariant()
		a.NotPanics(t, func() { v.Copy(val) })
		done <- callbackvalue{res: res, val: val}
	}

	r := client.MetadataAsync("/diagnosis/cfg/realtime/numbers", rc)
	a.Equal(t, datalayer.Result(0), r)

	select {
	case vals := <-done:
		a.Equal(t, vals.res, datalayer.Result(0))
		a.Equal(t, vals.val.GetType(), datalayer.VariantTypeFlatbuffers)
		datalayer.DeleteVariant(vals.val)
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
	close(done)
}

func TestClientCreateAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		done <- callbackvalue{res: res, val: nil}
	}

	args := datalayer.NewVariant()
	defer datalayer.DeleteVariant(args)

	r := client.CreateAsync("", args, rc)
	a.Equal(t, datalayer.Result(0), r)

	select {
	case vals := <-done:
		a.Equal(t, vals.res, datalayer.Result(0x80010001)) // DL_INVALID_ADDRESS
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
	close(done)
}

func TestClientRemoveAsync(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan callbackvalue)

	rc := func(res datalayer.Result, v *datalayer.Variant) {
		done <- callbackvalue{res: res, val: nil}
	}

	args := datalayer.NewVariant()
	defer datalayer.DeleteVariant(args)

	r := client.RemoveAsync("", rc)
	a.Equal(t, datalayer.Result(0), r)

	select {
	case vals := <-done:
		a.Equal(t, vals.res, datalayer.Result(0x80010001)) // DL_INVALID_ADDRESS
	case <-time.After(asyncTestTimeout()):
		a.Fail(t, "callback timeout")
	}
	close(done)
}
