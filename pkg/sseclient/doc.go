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

/*
SSE stands for Server-Sent Events. Server-Sent Events is a technology, that the server pushes the data to the client,
the client receives automatic updates from the server via an HTTP connection.

ctrlX Data Layer has implemented a SSE Server, which provides data access via SSE protocol.

Package sseclient implements a SSE Client, which is used to connect the SSE server on the ctrlX Data Layer.

Before we create SseClient, two things are needed:

1) Get authorization token
The token is a Bearer token, which is provied by ctrlX Identity Manager Component.
2) Create SSE subscription on the SSE Server on the Data Layer
The subscriptionID is generated at first on client side by using for example package uuid, this id then will be sent to the
server together with other SSE parameters especially nodelist by sending a POST request to the ctrlX Data Layer. The subscription
then will be created on the server side, and it can be accessed by using the subscriptionID.

Let us now create a SseClient:

The function NewSseClient needs three parameters:
* the url is like "/automation/api/v2/events/<subscriptionID>",
* the token is like "Bearer eyJhbGciOiJIUxxx"
* the flag insecureSkipVerify has to be set to true if no corresponding certificate has been installed on the ctrlX CORE, which is usually not the case in delivery state.

	client := sseclient.NewSseClient(url, token, true)

Reading event data:

The function Subscribe starts GET request and then reads the event data continuouslly.
* The paramter context is used to cancell the reading loop
* The callback function of type SseReceiverFunc is used to handle the incoming event data.

This function will only terminate in case of connection errors (err != nil) or if the subscription is terminated
by the main program side (err == nil).

  err = client.Subscribe(ctx, func(event string, data string) {
		...
	})

The SSE Server sends following types of event and data

"event: update\n"
"id: ...\n"
"data: {"node":".....","timestamp":....,"type":"...","value":...}\n"
"\n"

"event: error\n"
"id: ...\n"
"data: {"type":".....","title":....,"status":...,"instance":"...","severity":""}\n"
"\n"

"event: keepalive\n"
"id: ...\n"
"data: {"timestamp":....}\n"
"\n"

The new subscription should be created and the connection should be restablished if it is broken, the SSE Server on the
ctrlX Data Layer will delete the inactive existing subscription after a time period.

*/
package sseclient
