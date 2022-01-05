package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type environment struct {
	cgo_ldflags     string
	ld_library_path string
}

func main() {
	var osString string
	switch runtime.GOOS {
	case "linux":
		osString = runtime.GOOS
	case "windows":
		osString = "win"
	default:
		panic(fmt.Errorf("Unsupported os"))
	}

	var archString string
	switch runtime.GOARCH {
	case "amd64":
		archString = "x64"
	case "arm64":
		archString = "aarch64"
	default:
		panic(fmt.Errorf("unsupported os"))
	}

	dlr, err := filepath.Abs(fmt.Sprintf("./deps/public/bin/comm.datalayer/%s-gcc-%s/release", osString, archString))
	if err != nil {
		panic(err)
	}
	if stat, err := os.Stat(dlr); err != nil || !stat.IsDir() {
		panic(fmt.Errorf("something wrong with %s", dlr))
	}
	e := environment{
		fmt.Sprintf("-L%s", dlr),
		dlr,
	}

	env := reflect.ValueOf(&e).Elem()
	t := env.Type()

	for i := 0; i < env.NumField(); i++ {
		f := env.Field(i)
		fmt.Printf("export %s=\"%v\"\n", strings.ToUpper(t.Field(i).Name), f)
	}
}
