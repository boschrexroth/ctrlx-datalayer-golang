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

package sseclient_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/sseclient"
	"github.com/stretchr/testify/require"
)

var multiEntries = false
var mux sync.Mutex

func setMultiEntries(m bool) {
	mux.Lock()
	defer mux.Unlock()
	multiEntries = m
}

func getMultiEntries() bool {
	mux.Lock()
	defer mux.Unlock()
	return multiEntries
}

func newServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("event: update\n"))
		require.NoError(t, err)
		_, err = w.Write([]byte("id: 12345\n"))
		require.NoError(t, err)
		if getMultiEntries() {
			_, err = w.Write([]byte("data: {\n"))
			require.NoError(t, err)
			_, err = w.Write([]byte("data: \"node\":\"plc/app/Application/sym/PLC_PRG/counter\", \"timestamp\":132669450604571037,\"type\":\"double\",\"value\":44.0\n"))
			require.NoError(t, err)
			_, err = w.Write([]byte("data: }\n"))
			require.NoError(t, err)
		} else {
			_, err = w.Write([]byte("data: {\"node\":\"plc/app/Application/sym/PLC_PRG/counter\", \"timestamp\":132669450604571037,\"type\":\"double\",\"value\":43.0}\n"))
			require.NoError(t, err)
		}
		_, err = w.Write([]byte("\n"))
		require.NoError(t, err)
	}))
	return server
}

func cleanup(server *httptest.Server) {
	server.CloseClientConnections()
	server.Close()
}

func TestNewSseClient(t *testing.T) {
	s := sseclient.NewSseClient("url", "Bearer xxx", true)
	require.NotEmpty(t, s)
}

func TestSubscribe(t *testing.T) {
	server := newServer(t)
	defer cleanup(server)

	s := sseclient.NewSseClient(server.URL+"/event", "Bearer xxx", true)
	ctx := context.Background()
	s.Subscribe(ctx, func(event string, data string) {
		require.Equal(t, "update", event)
		require.Equal(t, "{\"node\":\"plc/app/Application/sym/PLC_PRG/counter\", \"timestamp\":132669450604571037,\"type\":\"double\",\"value\":43.0}", data)
	})
}

func TestSubscribeMulti(t *testing.T) {
	setMultiEntries(true)
	server := newServer(t)
	defer cleanup(server)
	defer setMultiEntries(false)

	s := sseclient.NewSseClient(server.URL+"/event", "Bearer xxx", true)
	ctx := context.Background()
	s.Subscribe(ctx, func(event string, data string) {
		require.Equal(t, "update", event)
		require.Equal(t, "{\"node\":\"plc/app/Application/sym/PLC_PRG/counter\", \"timestamp\":132669450604571037,\"type\":\"double\",\"value\":44.0}", data)
	})
}
