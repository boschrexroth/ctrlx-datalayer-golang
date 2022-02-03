# Datalayer Client and Provider for golang

## Preparation

Installation of necessary libraries:
```bash
sudo apt-get install libsystemd-dev libsystemd-dev:arm64 libzmq3-dev libzmq3-dev:arm64
```

Download and install ctrlX CORE deb-package:
```bash
wget https://github.com/boschrexroth/ctrlx-automation-sdk/releases/download/1.12.0/ctrlx-datalayer.deb
dpkg -i ctrlx-datalayer.deb
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

Copyright (c) 2022 Bosch Rexroth AG. All rights reserved.

<https://www.boschrexroth.com>

Bosch Rexroth AG
Bgm.-Dr.-Nebel-Str. 2
97816 Lohr am Main
GERMANY
