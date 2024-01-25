/*
 * SPDX-FileCopyrightText: Bosch Rexroth AG
 *
 * SPDX-License-Identifier: MIT
 */
package datalayer_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
	"github.com/stretchr/testify/suite"
)

// test suite client
type ClientSubscriptionTestSuite struct {
	ClientTestSuite
}

// test entry function
func TestClientSubTestSuite(t *testing.T) {
	suite.Run(t, new(ClientSubscriptionTestSuite))
}

func (suite *ClientSubscriptionTestSuite) TestClientSubscription() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

	done := make(chan bool)

	cm := map[string]bool{
		addressInt1: false,
		addressInt2: false,
	}

	onSubscribe := func(result datalayer.Result, items map[string]datalayer.Variant) {
		suite.Equal(datalayer.Result(0), result)
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
	suite.Equal(datalayer.Result(0), r)
	suite.NotNil(s)

	r = s.Subscribe(addressInt1,
		addressInt2)
	suite.Equal(datalayer.Result(0), r)

	suite.Equal(s.Id(), "mySub")
	suite.True(len(s.Addresses()) == 2)

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		suite.Fail("callback timeout")
	}

	suite.NotPanics(func() { client.DeleteSubscription(s) })
	close(done)
}

func (suite *ClientSubscriptionTestSuite) TestClientUnSubscription() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)
	suite.True(client.IsConnected())

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
		cm: map[string]int{addressInt1: 0,
			addressInt2: 0},
	}

	onSubscribe := func(result datalayer.Result, items map[string]datalayer.Variant) {
		//fmt.Println(">          onSub: ", time.Now().String(), len(items))
		suite.Equal(datalayer.Result(0), result)
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
	suite.Equal(datalayer.Result(0), r)
	suite.NotNil(s)

	r = s.Subscribe(addressInt1,
		addressInt2)
	suite.Equal(datalayer.Result(0), r)

	suite.Equal(s.Id(), "mySub")
	suite.True(len(s.Addresses()) == 2)

	select {
	case msg := <-done1:
		suite.True(msg, clientAddress)
	case <-time.After(5 * time.Second):
		suite.Fail("callback timeout")
	}

	//fmt.Println("> unsub", time.Now().String())
	suite.NotPanics(func() { s.Unsubscribe(addressInt2) })
	suite.True(len(s.Addresses()) == 1)

	tc.mu.Lock()
	tc.cm[addressInt1] = 0
	tc.cm[addressInt2] = -1

	tc.useCase1 = false
	tc.useCase2 = true

	tc.mu.Unlock()
	//fmt.Println("< unsub", time.Now().String())

	select {
	case <-done2:
	case <-time.After(10 * time.Second):
		suite.Fail("callback timeout", clientAddress)
	}

	suite.NotPanics(func() { client.DeleteSubscription(s) })
	tc.mu.Lock()
	suite.True(tc.cm[addressInt1] > 0, clientAddress)
	suite.True(tc.cm[addressInt2] == -1)
	tc.mu.Unlock()
	close(done1)
	close(done2)
}
