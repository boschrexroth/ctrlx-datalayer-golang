/*
 * SPDX-FileCopyrightText: Bosch Rexroth AG
 *
 * SPDX-License-Identifier: MIT
 */

package datalayer_test

import (
	"fmt"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
)

const (
	typeaddressstring = "types/datalayer/string"
	typeaddressbool   = "types/datalayer/bool8"
	typeaddressint32  = "types/datalayer/int32"
	addressbase       = "test-go-datalayer-provider/"
	addressString     = addressbase + "string"
	addressBool       = addressbase + "bool"
	addressInt1       = addressbase + "int1"
	addressInt2       = addressbase + "int2"
	addressFolder     = addressbase + "folder/"
)

// helper for 'local' provider
type testProvider struct {
	system   *datalayer.System
	provider *datalayer.Provider
	nodes    []nodeData
}

// interface for provider nodes
type nodeDataHandler interface {
	Name() string
	Node() *datalayer.ProviderNode
	Value() *datalayer.Variant
	OnMetadata(p *datalayer.ProviderNodeEvent) (datalayer.Result, *datalayer.Variant)
	OnCreate(p *datalayer.ProviderNodeEventData) datalayer.Result
	OnRemove(p *datalayer.ProviderNodeEvent) datalayer.Result
	OnRead(p *datalayer.ProviderNodeEventData) (datalayer.Result, *datalayer.Variant)
	OnWrite(p *datalayer.ProviderNodeEventData) datalayer.Result
}

// base for provider node
type nodeData struct {
	nodeDataHandler

	name string
	node *datalayer.ProviderNode
	val  *datalayer.Variant
}

// provider node from type 'string'
type nodeDataString struct {
	nodeData
}

// provider node from type 'bool'
type nodeDataBool struct {
	nodeData
}

// provider node from type 'int32' and dynamic values
type nodeDataInt32 struct {
	nodeData
	i int32
}

// provider node from type 'folder' and manages subnodes
type nodeDataFolder struct {
	nodeData
	provider        *datalayer.Provider
	subnodesHandler map[string]any
}

// create 'local' provider
func (suite *ClientTestSuite) initProvider() {
	suite.tp = &testProvider{}
	suite.tp.create()
}

// close 'local' provider
func (suite *ClientTestSuite) closeProvider() {
	if suite.tp == nil {
		return
	}
	suite.tp.close()
	suite.tp = nil
}

// this function executes bevor all tests executed
func (suite *ClientTestSuite) SetupSuite() {

	suite.initProvider()
	// set StartingNumber to one
	fmt.Println(">>> From SetupSuite")
}

// this function executes after all tests executed
func (suite *ClientTestSuite) TearDownSuite() {
	suite.closeProvider()
	fmt.Println(">>> From TearDownSuite")
}

// create a 'local' client
func (suite *ClientTestSuite) initClient() *datalayer.Client {
	system := suite.tp.system
	//system.Start(true)

	client := system.Factory().CreateClient("ipc://")
	client.SetTimeout(datalayer.TimeoutSettingPing, ctrlxClientTimeout())
	return client
}

// create a 'local' provider with provider nodes
func (p *testProvider) create() {
	p.nodes = make([]nodeData, 0)
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
func (p *testProvider) createNodes() {
	ndf := newNodeDataFolder(addressFolder+"**", p.provider)
	go startnodeDataHandler(ndf)
	r := p.provider.RegisterNode(ndf.Name(), ndf.Node())
	if r != datalayer.ResultOk {
		fmt.Println("ERROR Registering node ", ndf.Name(), " failed: ", r)
	}
	p.nodes = append(p.nodes, ndf.nodeData)

	//create sub nodes
	nds := newNodeDataString(addressString)
	r = p.provider.RegisterNode(nds.Name(), ndf.Node())
	if r != datalayer.ResultOk {
		fmt.Println("ERROR Registering node ", nds.Name(), " failed: ", r)
	}
	ndf.subnodesHandler[nds.Name()] = nds

	ndb := newNodeDataBool(addressBool)
	r = p.provider.RegisterNode(ndb.Name(), ndf.Node())
	if r != datalayer.ResultOk {
		fmt.Println("ERROR Registering node ", ndb.Name(), " failed: ", r)
	}
	ndf.subnodesHandler[ndb.Name()] = ndb

	ndi1 := newNodeDataInt32(addressInt1)
	r = p.provider.RegisterNode(ndi1.Name(), ndf.Node())
	if r != datalayer.ResultOk {
		fmt.Println("ERROR Registering node ", ndi1.Name(), " failed: ", r)
	}
	ndf.subnodesHandler[ndi1.Name()] = ndi1

	ndi2 := newNodeDataInt32(addressInt2)
	r = p.provider.RegisterNode(ndi2.Name(), ndf.Node())
	if r != datalayer.ResultOk {
		fmt.Println("ERROR Registering node ", ndi2.Name(), " failed: ", r)
	}
	ndf.subnodesHandler[ndi2.Name()] = ndi2
}

// close provider
func (p *testProvider) closeProvider() {
	for _, n := range p.nodes {
		p.provider.UnregisterNode(n.Name())
		deletenodeData(&n)
	}
	p.provider.Stop()
	datalayer.DeleteProvider(p.provider)
}

// close runtime
func (p *testProvider) close() {
	p.closeProvider()
	p.system.Stop(true)
	datalayer.DeleteSystem(p.system)
}

// node name
func (n *nodeData) Name() string {
	return n.name
}

// provider node
func (n *nodeData) Node() *datalayer.ProviderNode {
	return n.node
}

// node value
func (n *nodeData) Value() *datalayer.Variant {
	return n.val
}

// setting the write values
func (n *nodeData) OnWrite(p *datalayer.ProviderNodeEventData) datalayer.Result {
	return p.Data.Copy(n.Value())
}

// creating a provider subnode
func (n *nodeData) OnCreate(p *datalayer.ProviderNodeEventData) datalayer.Result {
	return datalayer.ResultInvalidAddress
}

// deleting a provider subnode
func (n *nodeData) OnRemove(p *datalayer.ProviderNodeEvent) datalayer.Result {
	deletenodeData(n)
	return datalayer.ResultOk
}

// helper delete
func deletenodeData(n *nodeData) {
	if n == nil {
		return
	}
	datalayer.DeleteProviderNode(n.node)
	datalayer.DeleteVariant(n.val)
}

// creating a provider subnode from type 'string'
func newNodeDataString(n string) *nodeDataString {
	nd := &nodeDataString{nodeData{name: n, node: nil, val: datalayer.NewVariant()}}
	nd.val.SetString("test")
	return nd
}

// OnMetadata descript 'string' MetaData of the provider node
func (n *nodeDataString) OnMetadata(p *datalayer.ProviderNodeEvent) (datalayer.Result, *datalayer.Variant) {
	m := datalayer.NewMetaDataBuilder(datalayer.AllowedOperationAll, "String variable", "String_variable_url")
	m.Unit("-").DisplayName(n.name).NodeClass(fbs.NodeClassVariable)
	m.AddReference(datalayer.ReferenceTypeRead, typeaddressstring)
	m.AddReference(datalayer.ReferenceTypeWrite, typeaddressstring)
	v := m.Build()
	return datalayer.ResultOk, v
}

// creating a provider subnode from type 'bool'
func newNodeDataBool(n string) *nodeDataBool {
	nd := &nodeDataBool{nodeData{name: n, node: nil, val: datalayer.NewVariant()}}
	nd.val.SetBool8(true)
	return nd
}

// OnMetadata descript 'bool' MetaData of the provider node
func (n *nodeDataBool) OnMetadata(p *datalayer.ProviderNodeEvent) (datalayer.Result, *datalayer.Variant) {
	m := datalayer.NewMetaDataBuilder(datalayer.AllowedOperationAll, "bool variable", "bool_variable_url")
	m.Unit("-").DisplayName(n.name).NodeClass(fbs.NodeClassVariable)
	m.AddReference(datalayer.ReferenceTypeRead, typeaddressbool)
	m.AddReference(datalayer.ReferenceTypeWrite, typeaddressbool)
	v := m.Build()
	return datalayer.ResultOk, v
}

// creating a provider subnode from type 'int32'
func newNodeDataInt32(n string) *nodeDataInt32 {
	nd := &nodeDataInt32{
		nodeData: nodeData{
			name: n,
			node: nil,
			val:  datalayer.NewVariant(),
		},
		i: int32(0),
	}
	nd.val.SetInt32(nd.i)
	return nd
}

// OnMetadata descript 'int32' MetaData of the provider node
func (n *nodeDataInt32) OnMetadata(p *datalayer.ProviderNodeEvent) (datalayer.Result, *datalayer.Variant) {
	m := datalayer.NewMetaDataBuilder(datalayer.AllowedOperationAll, "int32 variable", "int32_variable_url")
	m.Unit("-").DisplayName(n.name).NodeClass(fbs.NodeClassVariable)
	m.AddReference(datalayer.ReferenceTypeRead, typeaddressint32)
	m.AddReference(datalayer.ReferenceTypeWrite, typeaddressint32)
	v := m.Build()
	return datalayer.ResultOk, v
}

// dynamic values
func (n *nodeDataInt32) Value() *datalayer.Variant {
	n.i++
	n.val.SetInt32(n.i)
	return n.val
}

// OnMetadata descript 'folder' MetaData of the provider node
func (n *nodeDataFolder) OnMetadata(p *datalayer.ProviderNodeEvent) (datalayer.Result, *datalayer.Variant) {
	if p.Address == n.Name() {
		m := datalayer.NewMetaDataBuilder(datalayer.AllowedOperationBrowse|datalayer.AllowedOperationCreate|datalayer.AllowedOperationDelete|datalayer.AllowedOperationWrite, "Folder variable", "Folder_variable_url")
		m.Unit("-").DisplayName(n.name).NodeClass(fbs.NodeClassFolder)
		v := m.Build()
		return datalayer.ResultOk, v
	}
	// OnMetadata from subnodes
	e := n.subnodesHandler[p.Address]
	if e == nil {
		return datalayer.ResultInvalidAddress, nil
	}
	return e.(nodeDataHandler).OnMetadata(p)
}

// creating provider subnodes
func (n *nodeDataFolder) OnCreate(p *datalayer.ProviderNodeEventData) datalayer.Result {
	e := n.subnodesHandler[p.Address]
	if e != nil {
		return datalayer.ResultAlreadyExists
	}
	if p.Data != nil {
		if p.Data.GetType() == datalayer.VariantTypeBool8 {
			nb := newNodeDataBool(p.Address)
			r := n.provider.RegisterNode(nb.Name(), n.Node())
			if r != datalayer.ResultOk {
				fmt.Println("ERROR Registering node ", nb.Name(), " failed: ", r)
				return r
			}
			n.subnodesHandler[p.Address] = nb
			p.Data.Copy(nb.Value())
			return datalayer.ResultOk
		}
		if p.Data.GetType() == datalayer.VariantTypeInt32 {
			ni := newNodeDataInt32(p.Address)
			r := n.provider.RegisterNode(ni.Name(), n.Node())
			if r != datalayer.ResultOk {
				fmt.Println("ERROR Registering node ", ni.Name(), " failed: ", r)
				return r
			}
			n.subnodesHandler[p.Address] = ni
			p.Data.Copy(ni.Value())
			return datalayer.ResultOk
		}
	}
	//default node -> string
	ns := newNodeDataString(p.Address)
	r := n.provider.RegisterNode(ns.Name(), n.Node())
	if r != datalayer.ResultOk {
		fmt.Println("ERROR Registering node ", ns.Name(), " failed: ", r)
		return r
	}
	n.subnodesHandler[p.Address] = ns
	if p.Data != nil {
		p.Data.Copy(ns.Value())
	}

	return datalayer.ResultOk
}

// deleting provider subnodes
func (n *nodeDataFolder) OnRemove(p *datalayer.ProviderNodeEvent) datalayer.Result {
	if p.Address == n.Name() {
		return datalayer.ResultUnsupported
	}
	e := n.subnodesHandler[p.Address]
	if e == nil {
		return datalayer.ResultInvalidAddress
	}
	n.provider.UnregisterNode(p.Address)
	delete(n.subnodesHandler, p.Address)

	return e.(nodeDataHandler).OnRemove(p)
}

// reading proider subnodes value
func (n *nodeDataFolder) OnRead(p *datalayer.ProviderNodeEventData) (datalayer.Result, *datalayer.Variant) {
	if p.Address == n.Name() {
		return datalayer.ResultOk, n.Value()
	}
	e := n.subnodesHandler[p.Address]
	if e == nil {
		return datalayer.ResultInvalidAddress, nil
	}
	return datalayer.ResultOk, e.(nodeDataHandler).Value()
}

// writing proider subnodes value
func (n *nodeDataFolder) OnWrite(p *datalayer.ProviderNodeEventData) datalayer.Result {
	if p.Address == n.Name() {
		return p.Data.Copy(n.Value())
	}
	e := n.subnodesHandler[p.Address]
	if e == nil {
		return datalayer.ResultInvalidAddress
	}
	return e.(nodeDataHandler).OnWrite(p)
}

// creating a provider subnode from type 'folder'
func newNodeDataFolder(n string, provider *datalayer.Provider) *nodeDataFolder {
	nd := &nodeDataFolder{
		nodeData: nodeData{
			name: n,
			node: datalayer.NewProviderNode(),
			val:  datalayer.NewVariant(),
		},
		provider:        provider,
		subnodesHandler: make(map[string]any),
	}
	nd.val.SetBool8(true)
	return nd
}

// handler for provider node
func startnodeDataHandler(d nodeDataHandler) {
	for {
		if d.Node() == nil {
			return
		}
		select {
		case event, ok := <-d.Node().Channels().OnCreate:
			if !ok {
				return
			}
			r := d.OnCreate(&event)
			fmt.Println("event: oncreate: ", d.Name(), event.Address, r)
			event.Callback(r, d.Value())

		case event, ok := <-d.Node().Channels().OnRemove:
			if !ok {
				return
			}

			fmt.Println("event: OnRemove: ", d.Name())
			r := d.OnRemove(&event)
			event.Callback(r, nil)

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
				r, v := d.OnRead(&event)
				if r == datalayer.ResultOk {
					v.Copy(newData)
				}
				fmt.Println("event: OnRead: ", d.Name(), event.Address, r)
				event.Callback(r, newData)
			}()

		case event, ok := <-d.Node().Channels().OnWrite:
			if !ok {
				return
			}

			r := d.OnWrite(&event)
			fmt.Println("event: OnWrite: ", d.Name())
			event.Callback(r, event.Data)

		case event, ok := <-d.Node().Channels().OnMetadata:
			func() {
				if !ok {
					return
				}

				fmt.Println("event: OnMetadata: ", d.Name())

				r, v := d.OnMetadata(&event)
				defer datalayer.DeleteVariant(v)
				event.Callback(r, v)
			}()

		case <-d.Node().Channels().Done:
			fmt.Println("event: Done: ", d.Name())
			return
		}
	}
}
