package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProviderSuite struct {
	suite.Suite
	provider   *datalayerProvider
	system     *datalayer.System
	dlProvider *datalayer.Provider
	connect    string
	client     *datalayer.Client
}

func TestProviderSuite(t *testing.T) {
	suite.Run(t, new(ProviderSuite))
}

func (s *ProviderSuite) SetupTest() {
	s.connect = "ipc://"
	s.system = datalayer.NewSystem("web.common.datalayer.provider.test")
	s.system.Start(true)

	s.dlProvider = s.system.Factory().CreateProvider(s.connect)

	s.provider = NewProvider("test/base/address").(*datalayerProvider)
	result := s.dlProvider.Start()
	s.Require().Equal(datalayer.ResultOk, result, fmt.Sprintf("Failed to start provider: %s", result))
	s.client = s.system.Factory().CreateClient(s.connect)
}

func (s *ProviderSuite) TearDownTest() {
	s.provider.Unregister()
	_ = s.dlProvider.Stop()
	datalayer.DeleteProvider(s.dlProvider)
	datalayer.DeleteClient(s.client)
	s.system.Stop(true)
}

func (s *ProviderSuite) TestProvideCreateInvalidProvider() {
	provider := newMockprovider(s.T())
	s.provider.Provide(s.dlProvider, provider, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	data := datalayer.NewVariant()
	defer datalayer.DeleteVariant(data)

	result := s.client.CreateSync("test/base/address/child", data)
	s.Assert().Equal(datalayer.ResultUnsupported, result)
}

func (s *ProviderSuite) TestProvideRemoveInvalidProvider() {
	provider := newMockprovider(s.T())
	s.provider.Provide(s.dlProvider, provider, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	result := s.client.RemoveSync("test/base/address/child")
	s.Assert().Equal(datalayer.ResultUnsupported, result)
}

func (s *ProviderSuite) TestProvideBrowseInvalidProvider() {
	provider := newMockprovider(s.T())
	s.provider.Provide(s.dlProvider, provider, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	responseCallback := func(result datalayer.Result, v *datalayer.Variant) {
		s.Assert().Equal(datalayer.ResultOk, result)
		s.Assert().NotNil(v)
	}

	result := s.client.BrowseAsync("test/base/address/child", responseCallback)
	s.Assert().Equal(datalayer.ResultOk, result)
}

func (s *ProviderSuite) TestProvideReadInvalidProvider() {
	provider := newMockprovider(s.T())
	s.provider.Provide(s.dlProvider, provider, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	result, data := s.client.ReadSync("test/base/address/child")
	s.Assert().Equal(datalayer.ResultUnsupported, result)
	s.Assert().NotNil(data)
}

func (s *ProviderSuite) TestProvideWriteInvalidProvider() {
	provider := newMockprovider(s.T())
	s.provider.Provide(s.dlProvider, provider, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	data := datalayer.NewVariant()
	defer datalayer.DeleteVariant(data)

	result := s.client.WriteSync("test/base/address/child", data)
	s.Assert().Equal(datalayer.ResultUnsupported, result)
}

func (s *ProviderSuite) TestProvideMetadataInvalidProvider() {
	provider := newMockprovider(s.T())
	s.provider.Provide(s.dlProvider, provider, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	result, data := s.client.MetadataSync("test/base/address/child")
	s.Assert().Equal(datalayer.ResultUnsupported, result)
	s.Assert().NotNil(data)
}

func (s *ProviderSuite) TestProvideCreateSuccess() {
	create := newMockonCreateProvider(s.T())
	create.On("OnCreate", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		event := args.Get(0).(datalayer.ProviderNodeEventData)
		event.Callback(datalayer.ResultOk, nil)
	})
	s.provider.Provide(s.dlProvider, create, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	data := datalayer.NewVariant()
	defer datalayer.DeleteVariant(data)

	result := s.client.CreateSync("test/base/address/child", data)
	s.Assert().Equal(datalayer.ResultOk, result)
}

func (s *ProviderSuite) TestProvideRemoveSuccess() {
	remove := newMockonRemoveProvider(s.T())
	remove.On("OnRemove", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		event := args.Get(0).(datalayer.ProviderNodeEvent)
		event.Callback(datalayer.ResultOk, nil)
	})
	s.provider.Provide(s.dlProvider, remove, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	result := s.client.RemoveSync("test/base/address/child")
	s.Assert().Equal(datalayer.ResultOk, result)
}

/*
func (s *ProviderSuite) TestProvideBrowseSuccess()   {
	browse := mocks.NewOnBrowseProvider(s.T())
	browse.On("OnBrowse", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		event := args.Get(0).(datalayer.ProviderNodeEvent)
		event.Callback(datalayer.ResultOk, nil)
	})
	s.provider.Provide(s.dlProvider, browse, "*", "testField")
	time.Sleep(time.Millisecond * 100)

	result, data := s.client.BrowseSync("test/base/address/child")
	s.Assert().Equal(datalayer.ResultOk, result)
	s.Assert().NotNil(data)
}*/

func (s *ProviderSuite) TestProvideReadSuccess() {
	read := newMockonReadProvider(s.T())
	read.On("OnRead", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		event := args.Get(0).(datalayer.ProviderNodeEventData)
		event.Callback(datalayer.ResultOk, nil)
	})
	s.provider.Provide(s.dlProvider, read, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	result, data := s.client.ReadSync("test/base/address/child")
	s.Assert().Equal(datalayer.ResultOk, result)
	s.Assert().NotNil(data)
}

func (s *ProviderSuite) TestProvideWriteSuccess() {
	write := newMockonWriteProvider(s.T())
	write.On("OnWrite", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		event := args.Get(0).(datalayer.ProviderNodeEventData)
		event.Callback(datalayer.ResultOk, nil)
	})
	s.provider.Provide(s.dlProvider, write, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	data := datalayer.NewVariant()
	defer datalayer.DeleteVariant(data)

	result := s.client.WriteSync("test/base/address/child", data)
	s.Assert().Equal(datalayer.ResultOk, result)
}

func (s *ProviderSuite) TestProvideMetadataSuccess() {
	metadata := newMockonMetadataProvider(s.T())
	metadata.On("OnMetadata", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		event := args.Get(0).(datalayer.ProviderNodeEvent)
		event.Callback(datalayer.ResultOk, nil)
	})
	s.provider.Provide(s.dlProvider, metadata, "/child", "testField")
	time.Sleep(time.Millisecond * 100)

	data := datalayer.NewVariant()
	defer datalayer.DeleteVariant(data)

	result, data := s.client.MetadataSync("test/base/address/child")
	s.Assert().Equal(datalayer.ResultOk, result)
	s.Assert().NotNil(data)
}

func (s *ProviderSuite) TestStopProvider() {
	p := newMockprovider(s.T())
	p.On("Stop")
	s.provider.provider = p
	s.provider.StopProvider()
}
