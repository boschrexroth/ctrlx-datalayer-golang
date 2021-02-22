package datalayer_test

import (
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	a "github.com/stretchr/testify/assert"
)

const clientAddress string = "tcp://boschrexroth:boschrexroth@192.168.64.129:2069"

func TestClient(t *testing.T) {
	system := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(system)
	a.NotNil(t, system)

	a.NotPanics(t, func() { system.Start(false) })

	client := system.Factory().CreateClient(clientAddress)
	defer datalayer.DeleteClient(client)
	a.NotNil(t, client)

	a.Equal(t, datalayer.Result(0), client.SetTimeout(datalayer.TimeoutSettingPing, 2000))
	a.Equal(t, datalayer.Result(0), client.PingSync())
}
