package provider

import (
	"strings"
	"sync"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
)

type DatalayerProvider interface {
	Provide(dlProvider *datalayer.Provider, provider Provider, address string, field string)
	StopProvider()
	Unregister()
}

type datalayerProvider struct {
	nodes       []string
	stop        map[string]*chan bool
	dlProvider  *datalayer.Provider
	provider    Provider
	mutex       sync.Mutex
	baseAddress string
}

func NewProvider(baseAddress string) DatalayerProvider {
	return &datalayerProvider{
		baseAddress: strings.TrimSuffix(baseAddress, "/"),
	}
}

func (dlData *datalayerProvider) Provide(dlProvider *datalayer.Provider, provider Provider, address string, field string) {
	dlData.mutex.Lock()
	defer dlData.mutex.Unlock()
	node := datalayer.NewProviderNode()
	nodeName := dlData.baseAddress + "/" + strings.TrimPrefix(address, "/")
	if len(dlData.nodes) == 0 {
		dlData.nodes = []string{}
		dlData.stop = make(map[string]*chan bool)
	}
	dlData.nodes = append(dlData.nodes, nodeName)
	stopChan := make(chan bool)
	dlData.stop[nodeName] = &stopChan
	dlData.dlProvider = dlProvider
	dlData.provider = provider
	_ = (*dlProvider).RegisterNode(nodeName, node)
	go func() {
		defer func() {
			datalayer.DeleteProviderNode(node)
		}()
		for {
			select {
			case event := <-node.Channels().OnCreate:
				if createProvider, ok := provider.(onCreateProvider); ok {
					createProvider.OnCreate(event, field)
				} else {
					event.Callback(datalayer.ResultUnsupported, nil)
				}
			case event := <-node.Channels().OnRemove:
				if removeProvider, ok := provider.(onRemoveProvider); ok {
					removeProvider.OnRemove(event, field)
				} else {
					event.Callback(datalayer.ResultUnsupported, nil)
				}
			case event := <-node.Channels().OnBrowse:
				if browseProvider, ok := provider.(onBrowseProvider); ok {
					browseProvider.OnBrowse(event, field)
				}
				// no else needed - default browse will answer!
			case event := <-node.Channels().OnRead:
				if readProvider, ok := provider.(onReadProvider); ok {
					readProvider.OnRead(event, field)
				} else {
					event.Callback(datalayer.ResultUnsupported, nil)
				}
			case event := <-node.Channels().OnWrite:
				if writeProvider, ok := provider.(onWriteProvider); ok {
					writeProvider.OnWrite(event, field)
				} else {
					event.Callback(datalayer.ResultUnsupported, nil)
				}
			case event := <-node.Channels().OnMetadata:
				if metadataProvider, ok := provider.(onMetadataProvider); ok {
					metadataProvider.OnMetadata(event, field)
				} else {
					event.Callback(datalayer.ResultUnsupported, nil)
				}
			case <-node.Channels().Done:
				return
			case <-stopChan:
				return
			}
		}
	}()
}

func (dlData *datalayerProvider) Unregister() {
	dlData.mutex.Lock()
	defer dlData.mutex.Unlock()
	for _, nodeName := range dlData.nodes {
		close((*dlData.stop[nodeName]))
		_ = dlData.dlProvider.UnregisterNode(nodeName)
	}
}

// StopProvider implements DatalayerProvider.
func (dlData *datalayerProvider) StopProvider() {
	dlData.provider.Stop()
}
