# ctrlX Data Layer client and provider library for golang

This project provides:
  
1. ctrlX Data Layer access via Golang

    It wraps the original ctrlX Data Layer written in C++.

    You can find documentation and examples here [ctrlX Software Development Kit](https://boschrexroth.github.io/ctrlx-automation-sdk/) and [GoDoc](https://pkg.go.dev/github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer).

2. Server-Sent Events(SSE) Client

    Server-Sent Events is HTTP/HTTP2 common standard, the SSE server pushes the update to the client, the client receives automatic updates via HTTP connection. 

    ctrlX Data Layer provides a SSE Server, which makes the data on the Data Layer accessible to the web application.  

    The SSE client library can be used to connect to the SSE Server of the ctrlX Data Layer running on the ctrlX Device. It is possible to subscribe to nodes on the server and the server will automatically send new updates on value changes.

    You can find documentation here [GoDoc](https://pkg.go.dev/github.com/boschrexroth/ctrlx-datalayer-golang/pkg/sseclient)

3. Token handling

    JSON Web Token is used for authentification in ctrlX Core. 

    The token library can be used to get token from ctrlX Core and later on check if the token is still valid. 

    You can find documentation here [GoDoc](https://pkg.go.dev/github.com/boschrexroth/ctrlx-datalayer-golang/pkg/token)

## Status


![Make Test](https://github.com/boschrexroth/ctrlx-datalayer-golang/actions/workflows/main.yml/badge.svg)

## Usage

### Prerequisites

```bash
sudo apt-get install libsystemd-dev libsystemd-dev:arm64 libzmq3-dev libzmq3-dev:arm64
```

### Download and install ctrlX Data Layer debian package

Check the current [ctrlx Data Layer](https://github.com/boschrexroth/ctrlx-automation-sdk/releases) debian package, download and install this, see example.

```bash
wget https://github.com/boschrexroth/ctrlx-automation-sdk/releases/download/1.18.0/ctrlx-datalayer-1.9.1.deb
sudo dpkg -i ctrlx-datalayer-1.9.1.deb
```

### Integrate in you project

To integrate this library in your project, simply install it via:

`go get github.com/boschrexroth/ctrlx-datalayer-golang`

## About

Copyright © 2022 Bosch Rexroth AG. All rights reserved.

<https://www.boschrexroth.com>

Bosch Rexroth AG  
Bgm.-Dr.-Nebel-Str. 2  
97816 Lohr am Main  
GERMANY

## Licenses

MIT License

Copyright (c) 2021-2022 Bosch Rexroth AG

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
