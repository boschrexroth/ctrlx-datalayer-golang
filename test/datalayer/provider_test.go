package datalayer_test

import (
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	a "github.com/stretchr/testify/assert"
)

const providerAddress string = "tcp://boschrexroth:boschrexroth@192.168.64.129:2070"

var goFloat float32 = 0.369

func TestProvider(t *testing.T) {
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(false) })

	provider := system.Factory().CreateProvider(providerAddress)
	defer datalayer.DeleteProvider(provider)
	a.NotNil(t, provider)

	node := datalayer.NewProviderNode()
	defer datalayer.DeleteProviderNode(node)
	a.NotNil(t, node)

	a.NotPanics(t, func() { go startTestNode(node, t) })

	a.Equal(t, datalayer.Result(0), provider.RegisterNode("godata/test", node))
	a.Equal(t, datalayer.Result(0), provider.Start())
	a.Equal(t, true, provider.IsConnected())
}

func startTestNode(node *datalayer.ProviderNode, t *testing.T) {
	for {
		select {
		case event := <-node.Channels().OnCreate:
			event.Callback(datalayer.Result(0), nil)
		case event := <-node.Channels().OnRemove:
			event.Callback(datalayer.Result(0), nil)
		case event := <-node.Channels().OnBrowse:
			newData := datalayer.NewVariant()
			defer datalayer.DeleteVariant(newData)
			newData.SetArrayString([]string{"d1", "d2"})
			event.Callback(datalayer.Result(0), newData)
		case event := <-node.Channels().OnRead:
			newData := datalayer.NewVariant()
			defer datalayer.DeleteVariant(newData)
			newData.SetFloat32(goFloat)
			event.Callback(datalayer.Result(0), newData)
		case event := <-node.Channels().OnWrite:
			goFloat = event.Data.GetFloat32()
			event.Callback(datalayer.Result(0), event.Data)
		case event := <-node.Channels().OnMetadata:
			event.Callback(datalayer.Result(0), nil)
		case <-node.Channels().Done:
			return
		}
	}
}
