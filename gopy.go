package main

import (
	"net/http"
	"os"
	"path/filepath"
)

var (
	gopath = os.Getenv("GOPATH")
	approot = filepath.Join(gopath, "src/github.com/Extremus-io/gopy/")
	host = "127.0.0.1:9000"
	webroot = filepath.Join(approot, "webapp/")
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(webroot)))
	http.ListenAndServe(host, nil)
}
