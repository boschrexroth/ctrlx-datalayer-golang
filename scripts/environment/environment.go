package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type environment struct {
	cflags          string
	ldflags         string
	ld_library_path string
}

func main() {
	include, err := filepath.Abs("./deps/public/include/comm.datalayer/comm/datalayer/c")
	if err != nil {
		panic(err)
	}
	if stat, err := os.Stat(include); err != nil || !stat.IsDir() {
		panic("Something wrong with " + include)
	}
	dlr, err := filepath.Abs("./deps/public/bin/comm.datalayer/linux-gcc-x64/release")
	if err != nil {
		panic(err)
	}
	if stat, err := os.Stat(dlr); err != nil || !stat.IsDir() {
		panic("Something wrong with " + dlr)
	}
	zmq, err := filepath.Abs("./deps/public/bin/zmq/linux-gcc-x64/release")
	if err != nil {
		panic(err)
	}
	if stat, err := os.Stat(zmq); err != nil || !stat.IsDir() {
		panic("Something wrong with " + zmq)
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
		fmt.Printf("%s=%v\n", strings.ToUpper(t.Field(i).Name), f)
	}
}
