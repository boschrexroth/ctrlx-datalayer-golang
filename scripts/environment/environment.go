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
	cflags          string
	ldflags         string
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
		panic(fmt.Errorf("Unsupported os"))
	}

	include, err := filepath.Abs("./deps/public/include/comm.datalayer/comm/datalayer/c")
	if err != nil {
		panic(err)
	}
	if stat, err := os.Stat(include); err != nil || !stat.IsDir() {
		panic(fmt.Errorf("Something wrong with %s", include))
	}
	dlr, err := filepath.Abs(fmt.Sprintf("./deps/public/bin/comm.datalayer/%s-gcc-%s/release", osString, archString))
	if err != nil {
		panic(err)
	}
	if stat, err := os.Stat(dlr); err != nil || !stat.IsDir() {
		panic(fmt.Errorf("Something wrong with %s", dlr))
	}
	zmq, err := filepath.Abs(fmt.Sprintf("./deps/public/bin/zmq/%s-gcc-%s/release", osString, archString))
	if err != nil {
		panic(err)
	}
	if stat, err := os.Stat(zmq); err != nil || !stat.IsDir() {
		panic(fmt.Errorf("Something wrong with %s", zmq))
	}
	e := environment{
		fmt.Sprintf("-g -O2 -I%s", include),
		fmt.Sprintf("-L%s -L%s", zmq, dlr),
		fmt.Sprintf("%s;%s", zmq, dlr),
	}

	env := reflect.ValueOf(&e).Elem()
	t := env.Type()

	for i := 0; i < env.NumField(); i++ {
		f := env.Field(i)
		fmt.Printf("%s=\"%v\"\n", strings.ToUpper(t.Field(i).Name), f)
	}
}
