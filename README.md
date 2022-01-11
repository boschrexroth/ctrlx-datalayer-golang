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

### Windows 
Vet/Test/Build from the commandline can be done using mage (https://magefile.org/). 
You can either install mage as described there or run
```bash
./mage.sh [vet,build:[amd64,arm64,win64,linux,all],test,testcover]
```
to execute distinct build steps.

Run `./mage.sh` without arguments to get a description of all available targets.

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

### Windows
Because this package is heavily using CGO, CGO_LDFLAGS and LD_LIBRARY_PATH need to be set correctly in projects using it as dependency.
If you build your package using mage, you can use the `magelib` package from this repository by adding

```golang
import "github.com/boschrexroth/ctrlx-datalayer-golang/magelib"
```
to your `magefile.go`. The `magelib` has methods to get CGO_LDFLAGS and LD_LIBRARY_PATH, to build executables and libraries
importing `pkg/datalayer` and to run go commands like `go vet` or `go test` on your package with the correct environment variables.
