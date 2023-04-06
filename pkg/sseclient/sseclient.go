// MIT License
//
// Copyright (c) 2021-2022 Bosch Rexroth AG
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package sseclient

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// This file contains an implementation of a Server Sent Event (SSE) client as standardized as part of
// HTML5 by the W3C. See: https://html.spec.whatwg.org/multipage/server-sent-events.html

// SseReceiverFunc defines the callback function that will be called on new event data.
type SseReceiverFunc func(event string, data string)

// SseClient is an instance of an SSE-Client sending events.
type SseClient struct {
	URL        string
	Connection *http.Client
	Headers    map[string]string
}

// SseEvent is a go representation of an http server-sent event.
type SseEvent struct {
	Event string
	Data  string
}

// NewSseClient creates a new client.
func NewSseClient(url string, authorization string, insecureSkipVerify bool) *SseClient {
	c := &SseClient{
		URL:        url,
		Connection: &http.Client{},
		Headers:    make(map[string]string),
	}
	if authorization != "" {
		c.Headers["Authorization"] = authorization
	}
	c.Connection.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify}}
	return c
}

// startSseConnection starts a sse stream connection and returns the established connection on success.
func (c *SseClient) startSseConnection() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting sse request for %q: %v", c.URL, err)
	}

	req.Header.Set("Accept", "text/event-stream")

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	res, err := c.Connection.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing request for %q: %v", c.URL, err)
	}
	if res.StatusCode > 201 {
		return nil, fmt.Errorf("wrong response for %q with statuscode %d", c.URL, res.StatusCode)
	}

	return res, nil
}

// Subscribe starts the event stream and reads the incoming data.
// Data is parsed as Sse Event and forwarded to the given callback func.
func (c *SseClient) Subscribe(ctx context.Context, receiver SseReceiverFunc) error {
	res, err := c.startSseConnection()
	if err != nil {
		return fmt.Errorf("failed to start sse request, %v", err)
	}

	br := bufio.NewReader(res.Body)
	defer res.Body.Close()

	var event, data string
	// Endless Loop - until error occurs or server disconnects
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			bytesRead, err := br.ReadBytes('\n')
			if err == io.EOF {
				// Server is down
				return fmt.Errorf("got end of sse data input")
			}
			if err != nil {
				return fmt.Errorf("error: %v", err)
			}

			//	Server sends:
			//		"event: update\n"
			//		"id: ...\n"
			//		"data: {"node":".....","timestamp":....,"type":"...","value":...}\n"
			//		"\n" --> End-of-message found
			//
			//		"event: error\n"
			//		"id: ...\n"
			//		"data: {"type":".....","title":....,"status":...,"instance":"...","severity":""}\n"
			//		"\n" --> End-of-message found
			//
			//		"event: keepalive\n"
			//		"id: ...\n"
			//		"data: {"timestamp":....}\n"
			//		"\n" --> End-of-message found

			// empty line marks the end of the sse event message
			if len(bytesRead) == 1 && bytesRead[0] == '\n' {
				// Forward the received data to caller
				receiver(event, data)
				event = ""
				data = ""
				continue
			}

			stringRead := strings.Trim(string(bytesRead), "\n")

			parts := strings.SplitN(stringRead, ": ", 2)

			if len(parts) < 2 {
				continue
			}

			switch parts[0] {
			case "event":
				event = parts[1]
			case "data":
				data = data + parts[1]
				// all other events will be ignored, continue
			}
		}
	}
}
