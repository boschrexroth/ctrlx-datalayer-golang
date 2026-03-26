package provider

import (
	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
)

type Provider interface {
	Stop()
}

type onCreateProvider interface {
	Provider
	OnCreate(event datalayer.ProviderNodeEventData, field string)
}

type onRemoveProvider interface {
	Provider
	OnRemove(event datalayer.ProviderNodeEvent, field string)
}

type onBrowseProvider interface {
	Provider
	OnBrowse(event datalayer.ProviderNodeEvent, field string)
}

type onReadProvider interface {
	Provider
	OnRead(event datalayer.ProviderNodeEventData, field string)
}

type onWriteProvider interface {
	Provider
	OnWrite(event datalayer.ProviderNodeEventData, field string)
}

type onMetadataProvider interface {
	Provider
	OnMetadata(event datalayer.ProviderNodeEvent, field string)
}
