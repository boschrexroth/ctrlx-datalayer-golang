package main

import (
	"errors"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
)

var (
	constFloat32 float32 = 42.0
	constInt64   int64   = 42
)

type Provider struct {
	provider   *datalayer.Provider
	done       chan struct{}
	nodes      map[string]*datalayer.ProviderNode
	events     chan datalayer.ProviderNodeEvent
	dataEvents chan datalayer.ProviderNodeEventData
}

func NewBechmarkProvider() *Provider {
	return &Provider{
		done: make(chan struct{}),
	}
}

func (p *Provider) Connect(system *datalayer.System, address string) error {
	p.provider = system.Factory().CreateProvider(address)
	p.startConstInt32()
	p.provider.Start()
	time.Sleep(1 * time.Second)
	connected := p.provider.IsConnected()
	if connected == false {
		return errors.New("Failed to connect")
	}
	return nil
}

func (p *Provider) Close() {
	datalayer.DeleteProvider(p.provider)
}

func (p *Provider) initConstInt32() {

}

func (p *Provider) handleConstInt32() {

}

func (p *Provider) startConstInt32() {
	a := "gobenchmark/int32"
	n := datalayer.NewProviderNode()
	p.provider.RegisterNode(a, n)
	v := datalayer.NewVariant()
	go func() {
		for {
			select {
			case event := <-n.Channels().OnCreate:
				p.dataEvents <- event
			case event := <-n.Channels().OnRemove:
				p.events <- event
			case event := <-n.Channels().OnBrowse:
				p.events <- event
			case event := <-n.Channels().OnRead:
				v.SetFloat32(constFloat32)
				event.Callback(datalayer.Result(0), v)
			case event := <-n.Channels().OnWrite:
				p.dataEvents <- event
			case event := <-n.Channels().OnMetadata:
				p.events <- event
			case <-n.Channels().Done:
				datalayer.DeleteProviderNode(n)
				return
			}
		}
	}()
	datalayer.DeleteVariant(v)
}
