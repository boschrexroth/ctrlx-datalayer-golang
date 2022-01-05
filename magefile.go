//go:build mage
// +build mage

package main

import (
	"context"
	"io/ioutil"

	"github.com/boschrexroth/ctrlx-datalayer-golang/magelib"
	"github.com/magefile/mage/mg"
)

type Build mg.Namespace

// Builds the executables in the cmd folder for windows amd64
func (Build) Win64(ctx context.Context) error {
	return buildExecutables(ctx, "windows", "amd64")
}

// Builds the executables in the cmd folder for linux amd64 (deprecated, only used on jenkins)
func (Build) Amd64(ctx context.Context) error {
	return buildExecutables(ctx, "linux", "amd64")
}

// Runs 'go vet ./...' with the required cgo build flags on windows amd64
func Vet(ctx context.Context) error {
	// if runtime.GOOS != "windows" {
	// 	return fmt.Errorf("veting with mage is windows only and not supported for GOOS %v. Please use make vet", runtime.GOOS)
	// }
	return magelib.RunGoCommand(ctx, "deps", "vet", "-mod", "vendor", "./pkg/datalayer")
}

// Runs the tests with the required cgo build flags on windows amd64
func Test(ctx context.Context) error {
	// if runtime.GOOS != "windows" {
	// 	return fmt.Errorf("testing with mage is windows only and not supported for GOOS %v. Please use make test", runtime.GOOS)
	// }
	return magelib.RunGoCommand(ctx, "deps", "test", "-mod", "vendor", "-x", "./test/datalayer")
}

// Runs the tests with the required cgo build flags with coverage and outputs coverage.out on windows amd64
func TestCover(ctx context.Context) error {
	// if runtime.GOOS != "windows" {
	// 	return fmt.Errorf("testing with mage is windows only and not supported for GOOS %v. Please use make test-coverage", runtime.GOOS)
	// }
	return magelib.RunGoCommand(ctx, "deps", "test", "-mod", "vendor", "-x", "-cover", "-coverprofile=coverage.out", "./test/datalayer", "-coverpkg=./pkg/datalayer")
}

func buildExecutables(ctx context.Context, goos, arch string) error {
	files, err := ioutil.ReadDir("cmd")
	if err != nil {
		return err
	}
	cmds := make([]string, 0)
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			cmds = append(cmds, "./cmd/"+fileInfo.Name())
		}
	}
	return magelib.BuildExecutables(ctx, goos, arch, "deps", "build", []string{"-mod", "vendor", "-x"}, cmds)
}
