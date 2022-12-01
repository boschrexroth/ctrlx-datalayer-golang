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
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/pkg/fbs/comm/datalayer"
	a "github.com/stretchr/testify/assert"
)

func init() {
	clientAddress = ctrlxAddress()
}

func TestClientSubscription(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done := make(chan bool)

	cm := map[string]bool{
		"framework/metrics/system/cpu-utilisation-percent": false,
		"framework/metrics/system/memused-percent":         false,
	}

	onSubscribe := func(result datalayer.Result, items map[string]datalayer.Variant) {
		a.Equal(t, datalayer.Result(0), result)
		for k, v := range items {
			fmt.Println(k, " ", v.GetUint32())
			cm[k] = true
		}
		for _, b := range cm {
			if !b {
				return
			}
		}
		done <- true
	}

	s, r := client.CreateSubscription("mySub", datalayer.DefaultSubscriptionProperties(), onSubscribe)
	a.Equal(t, datalayer.Result(0), r)
	a.NotNil(t, s)

	r = s.Subscribe("framework/metrics/system/cpu-utilisation-percent",
		"framework/metrics/system/memused-percent")
	a.Equal(t, datalayer.Result(0), r)

	a.Equal(t, s.Id(), "mySub")
	a.True(t, len(s.Addresses()) == 2)

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		a.Fail(t, "callback timeout")
	}

	a.NotPanics(t, func() { client.DeleteSubscription(s) })
	close(done)
}

func TestClientUnSubscription(t *testing.T) {
	if clientAddress == "" {
		t.Skip("ctrlX device does not exist")
	}
	system, client := initClient()
	defer datalayer.DeleteSystem(system)
	defer datalayer.DeleteClient(client)
	a.True(t, client.IsConnected())

	done1 := make(chan bool)
	done2 := make(chan bool)

	tc := struct {
		mu       sync.Mutex
		useCase1 bool
		useCase2 bool
		cm       map[string]int
	}{
		useCase1: true,
		useCase2: false,
		cm: map[string]int{"framework/metrics/system/cpu-utilisation-percent": 0,
			"framework/metrics/system/memused-percent": 0},
	}

	onSubscribe := func(result datalayer.Result, items map[string]datalayer.Variant) {
		//fmt.Println(">          onSub: ", time.Now().String(), len(items))
		a.Equal(t, datalayer.Result(0), result)
		tc.mu.Lock()
		defer tc.mu.Unlock()
		for k, v := range items {
			fmt.Println(k, " ", v.GetUint32())
			tc.cm[k]++
		}

		for _, b := range tc.cm {
			if b == 0 {
				return
			}
		}
		if tc.useCase1 {
			done1 <- true
			//fmt.Println("<          onSub: ", time.Now().String(), len(items))
		}
		if tc.useCase2 {
			done2 <- true
			//fmt.Println("<          onSub: ", time.Now().String(), len(items))
		}
	}
	p := datalayer.DefaultSubscriptionProperties()
	p.SamplingInterval = 1000
	p.DataChangeTrigger = fbs.DataChangeTriggerStatusValueTimestamp

	s, r := client.CreateSubscription("mySub", p, onSubscribe)
	a.Equal(t, datalayer.Result(0), r)
	a.NotNil(t, s)

	r = s.Subscribe("framework/metrics/system/cpu-utilisation-percent",
		"framework/metrics/system/memused-percent")
	a.Equal(t, datalayer.Result(0), r)

	a.Equal(t, s.Id(), "mySub")
	a.True(t, len(s.Addresses()) == 2)

	select {
	case msg := <-done1:
		a.True(t, msg, clientAddress)
	case <-time.After(5 * time.Second):
		a.Fail(t, "callback timeout")
	}

	//fmt.Println("> unsub", time.Now().String())
	a.NotPanics(t, func() { s.Unsubscribe("framework/metrics/system/memused-percent") })
	a.True(t, len(s.Addresses()) == 1)

	tc.mu.Lock()
	tc.cm["framework/metrics/system/cpu-utilisation-percent"] = 0
	tc.cm["framework/metrics/system/memused-percent"] = -1

	tc.useCase1 = false
	tc.useCase2 = true

	tc.mu.Unlock()
	//fmt.Println("< unsub", time.Now().String())

	select {
	case <-done2:
	case <-time.After(10 * time.Second):
		a.Fail(t, "callback timeout", clientAddress)
	}

	a.NotPanics(t, func() { client.DeleteSubscription(s) })
	tc.mu.Lock()
	a.True(t, tc.cm["framework/metrics/system/cpu-utilisation-percent"] > 0, clientAddress)
	a.True(t, tc.cm["framework/metrics/system/memused-percent"] == -1)
	tc.mu.Unlock()
	close(done1)
	close(done2)
}
