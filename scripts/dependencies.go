package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

//go:generate go run ./dependencies.go

type release struct {
	Assets []struct {
		BrowserDownloadURL string `json:"browser_download_url"`
		ID                 int    `json:"id"`
	}
}

type github struct {
	Github struct {
		User  string
		Token string
	}
}

func main() {
	var auth string
	var g github

	b, err := ioutil.ReadFile(".github")
	if err == nil {
		if err := yaml.Unmarshal(b, &g); err == nil {
			if g.Github.User != "" && g.Github.Token != "" {
				auth = fmt.Sprintf("%s:%s@", g.Github.User, g.Github.Token)
			}
		}
	}
	println("Looking for dependencies")
	resp, err := http.Get(fmt.Sprintf("https://%sapi.github.com/repos/boschrexroth/ctrlx-automation-sdk/releases/latest", auth))
	if err != nil {
		println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		println(resp.Status)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
		return
	}
	var r release
	json.Unmarshal(body, &r)

	println("Downloading dependencies")
	u := fmt.Sprintf("https://%sapi.github.com/repos/boschrexroth/ctrlx-automation-sdk/releases/assets/%v", auth, r.Assets[0].ID)
	req, err := http.NewRequest("GET", u, nil)
	req.Header.Set("Accept", "application/octet-stream")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		println(resp.Status)
		return
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
		return
	}
	prefix := "deps"
	if stat, err := os.Stat(prefix); err == nil && stat.IsDir() {
		os.RemoveAll(prefix)
	} else if err == nil && !stat.IsDir() {
		fmt.Printf("Error: %s exists, but is not a dir", prefix)
		return
	}
	z, err := zip.NewReader(bytes.NewReader(b), resp.ContentLength)
	if err != nil {
		println(err)
		return
	}
	wanted := []string{
		"public/bin/zmq/",
		"public/bin/comm.datalayer/",
		"public/include/comm.datalayer/comm/datalayer/c/",
	}
	for _, f := range z.File {
		for _, w := range wanted {
			if strings.HasPrefix(f.Name, w) {
				p := path.Join(prefix, f.Name)
				if f.FileInfo().IsDir() {
					os.MkdirAll(p, f.FileInfo().Mode())
				} else {
					zf, err := f.Open()
					if err != nil {
						println(err)
						return
					}
					defer zf.Close()
					cf, err := os.Create(p)
					if err != nil {
						println(err)
						return
					}
					fmt.Printf("Extracting %s to %s\n", f.FileInfo().Name(), cf.Name())
					io.Copy(cf, zf)
				}
			}
		}
	}
}
