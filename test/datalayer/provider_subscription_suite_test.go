/*
 * SPDX-FileCopyrightText: Bosch Rexroth AG
 *
 * SPDX-License-Identifier: MIT
 */

package datalayer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	"github.com/stretchr/testify/suite"
)

const (
	addressbasesub   = "test-go-datalayer-provider-sub/folder/"
	addressFolderSub = addressbasesub + "**"
	addressSubInt1   = addressbasesub + "int1"
)

// helper for 'local' provider
type testProviderSub struct {
	system   *datalayer.System
	provider *datalayer.Provider
	node     *newNodeDataFolderSub
}

// test suite client
type ProviderSubTestSuite struct {
	suite.Suite
	tp *testProviderSub
}

// test entry function
func TestProviderSubTestSuite(t *testing.T) {
	suite.Run(t, new(ProviderSubTestSuite))
}

func (suite *ProviderSubTestSuite) TestProviderClientSubscription() {
	client := suite.initClient()
	defer datalayer.DeleteClient(client)

	suite.True(client.IsConnected())

	done := make(chan bool)

	onSubscribe := func(result datalayer.Result, items map[string]datalayer.Variant) {
		suite.Equal(datalayer.Result(0), result)
		suite.Equal(len(items), 1)
		done <- true
	}

	s, r := client.CreateSubscription("myProviderSub", datalayer.DefaultSubscriptionProperties(), onSubscribe)
	suite.Equal(datalayer.Result(0), r)
	suite.NotNil(s)

	r = s.Subscribe(addressSubInt1)
	suite.Equal(datalayer.Result(0), r)

	suite.Equal(s.Id(), "myProviderSub")
	suite.True(len(s.Addresses()) == 1)

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		suite.Fail("callback timeout")
	}

	suite.NotPanics(func() { client.DeleteSubscription(s) })
	close(done)

}

// create 'local' provider
func (suite *ProviderSubTestSuite) initProvider() {
	suite.tp = &testProviderSub{}
	suite.tp.create()
}

// close 'local' provider
func (suite *ProviderSubTestSuite) closeProvider() {
	if suite.tp == nil {
		return
	}
	suite.tp.close()
	suite.tp = nil
}

// this function executes bevor all tests executed
func (suite *ProviderSubTestSuite) SetupSuite() {

	suite.initProvider()
	fmt.Println(">>> From SetupSuite")
}

// this function executes after all tests executed
func (suite *ProviderSubTestSuite) TearDownSuite() {
	suite.closeProvider()
	fmt.Println(">>> From TearDownSuite")
}

// create a 'local' client
func (suite *ProviderSubTestSuite) initClient() *datalayer.Client {
	system := suite.tp.system
	//system.Start(true)

	client := system.Factory().CreateClient("ipc://")
	client.SetTimeout(datalayer.TimeoutSettingPing, ctrlxClientTimeout())
	return client
}

// create a 'local' provider with provider nodes
func (p *testProviderSub) create() {
	p.system = datalayer.NewSystem("")
	p.system.Start(true)
	p.provider = p.system.Factory().CreateProvider("ipc://")
	r := p.provider.Start()
	fmt.Println("Start Provider: ", r)
	if r != datalayer.ResultOk {
		return
	}
	p.createNodes()
}

// create folder provider nodes and folder manage different subnodes
func (p *testProviderSub) createNodes() {
	ndf := &newNodeDataFolderSub{
		name:          addressFolderSub,
		node:          datalayer.NewProviderNodeSub(),
		Subscriptions: make(map[uint64]*datalayer.ProviderSubscription),
	}
	p.node = ndf
	go startnodeDataSubHandler(*ndf)
	r := p.provider.RegisterNode(ndf.Name(), ndf.Node())
	if r != datalayer.ResultOk {
		fmt.Println("ERROR Registering node ", ndf.Name(), " failed: ", r)
	}
}

// close provider
func (p *testProviderSub) closeProvider() {
	p.provider.UnregisterNode(addressFolderSub)
	datalayer.DeleteProviderNode(p.node.node)
	p.provider.Stop()
	datalayer.DeleteProvider(p.provider)
}

// close runtime
func (p *testProviderSub) close() {
	p.closeProvider()
	p.system.Stop(true)
	datalayer.DeleteSystem(p.system)
}

type newNodeDataFolderSub struct {
	name string
	node *datalayer.ProviderNode

	Subscriptions map[uint64]*datalayer.ProviderSubscription
}

func (n *newNodeDataFolderSub) Node() *datalayer.ProviderNode {
	return n.node
}

func (n *newNodeDataFolderSub) Name() string {
	return n.name
}

// handler for provider node
func startnodeDataSubHandler(d newNodeDataFolderSub) {
	for {
		if d.Node() == nil {
			return
		}
		select {
		case event, ok := <-d.Node().Channels().OnCreate:
			if !ok {
				return
			}
			fmt.Println("event: oncreate: ", d.Name(), event.Address)
			event.Callback(datalayer.ResultOk, nil)

		case event, ok := <-d.Node().Channels().OnRemove:
			if !ok {
				return
			}

			fmt.Println("event: OnRemove: ", d.Name())
			event.Callback(datalayer.ResultOk, nil)

		case event, ok := <-d.Node().Channels().OnBrowse:
			func() {
				if !ok {
					return
				}

				newData := datalayer.NewVariant()
				defer datalayer.DeleteVariant(newData)
				newData.SetArrayString([]string{d.Name()})
				fmt.Println("event: OnBrowse: ", d.Name())
				event.Callback(datalayer.Result(0), newData)
			}()

		case event, ok := <-d.Node().Channels().OnRead:
			func() {
				if !ok {
					return
				}

				newData := datalayer.NewVariant()
				defer datalayer.DeleteVariant(newData)
				newData.SetInt32(123)
				fmt.Println("event: OnRead: ", d.Name(), event.Address)
				event.Callback(datalayer.ResultOk, newData)
			}()

		case event, ok := <-d.Node().Channels().OnWrite:
			if !ok {
				return
			}

			fmt.Println("event: OnWrite: ", d.Name())
			event.Callback(datalayer.ResultOk, event.Data)

		case event, ok := <-d.Node().Channels().OnMetadata:
			func() {
				if !ok {
					return
				}

				fmt.Println("event: OnMetadata: ", d.Name())

				event.Callback(datalayer.ResultOk, nil)
			}()

		case event, ok := <-d.Node().Channels().OnSubscribe:
			func() {
				if !ok {
					return
				}

				fmt.Println("event: OnSubscribe: ", d.Name(), event.Address, event.Subsciption.GetUniqueId())
				fmt.Println("    Props: PublishInterval:", event.Subsciption.GetProps().PublishInterval())
				_, ok := d.Subscriptions[event.Subsciption.GetUniqueId()]
				if ok {
					return
				}
				d.Subscriptions[event.Subsciption.GetUniqueId()] = event.Subsciption
				n := datalayer.NewNotifyItemPublish(event.Address)
				l := []*datalayer.NotifyItemPublish{n}
				event.Subsciption.Publish(datalayer.ResultOk, l)
				datalayer.DeleteNotifyItemsPublish(l)
			}()
		case event, ok := <-d.Node().Channels().OnUnsubscribe:
			func() {
				if !ok {
					return
				}

				fmt.Println("event: OnUnsubscribe: ", d.Name(), event.Address, event.Subsciption.GetUniqueId())
				if len(event.Subsciption.GetNodes()) != 0 {
					return
				}
				delete(d.Subscriptions, event.Subsciption.GetUniqueId())
			}()
		case <-d.Node().Channels().Done:
			fmt.Println("event: Done: ", d.Name())
			return
		}
	}
}
