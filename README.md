# Datalayer Client and Provider for golang

This project provides ctrlX Data Layer access via Golang.
It wraps the original Data Layer SDK written in C++.

## Preparation

### Installation of necessary libraries:
```bash
sudo apt-get install libsystemd-dev libsystemd-dev:arm64 libzmq3-dev libzmq3-dev:arm64
```

### Download and install ctrlX Data Layer deb-package:

Check the current [ctrlx Data Layer](https://github.com/boschrexroth/ctrlx-automation-sdk/releases) debian package, download and install this, see example. 

```bash
wget https://github.com/boschrexroth/ctrlx-automation-sdk/releases/download/1.12.1/ctrlx-datalayer-1.7.5.deb
sudo dpkg -i ctrlx-datalayer-1.7.5.deb
```


## Development Environment

We recommend to use Visual Studio Code for development. The `.vscode/settings.json` already contains environment specific settings that a recommended.

### Linux
Vet/Test/Build from the commandline can be done simply with make

Run `make help` to get a description of all available targets.

## Tests
A lot of tests require a connection to a ctrlX device. The ctrlX device IP address you can set with help of the environment variable 
```
export CTRLX_ADDRESS=a.b.c.d
export CTRLX_TIMEOUT=4000               (Default: 2000 ms)
```

## Use as dependency

### Linux
To use this package, you need to install the libraries mentioned under [Preparation](#Preparation)

To cross-compile for the ctrlX CORE you will additionally need the `arm64` versions of the libraries:
```bash
sudo apt-get install -y libsystemd-dev:arm64 libzmq3-dev:arm64
```
The `ctrlx-datalayer.deb` is multi-arch and the `arm64` variant is already installed.

Note that you need to prefix your go build commands with `CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc` to cross compile. The arm compiler
can be installed with:
```bash
sudo apt install gcc-aarch64-linux-gnu
```
## About

Copyright © 2022 Bosch Rexroth AG. All rights reserved.

<https://www.boschrexroth.com>

Bosch Rexroth AG  
Bgm.-Dr.-Nebel-Str. 2  
97816 Lohr am Main  
GERMANY

## Licenses

MIT License

Copyright (c) 2021-2022-2022 Bosch Rexroth AG

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