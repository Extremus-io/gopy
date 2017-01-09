package main

import (
	"net/http"
	"os"
	"path/filepath"
	"os/exec"
	"io"
	"golang.org/x/net/websocket"
)

var (
	gopath = os.Getenv("GOPATH")
	approot = filepath.Join(gopath, "src/github.com/Extremus-io/gopy/")
	host = "127.0.0.1:9000"
	webroot = filepath.Join(approot, "webapp/")
)

func main() {
	http.Handle("/api/ws/", websocket.Handler(apiwsCall))
	http.HandleFunc("/api/", apiCall)
	http.Handle("/", http.FileServer(http.Dir(webroot)))
	http.ListenAndServe(host, nil)
}

func apiCall(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("python", "-u", "func.py")
	w.WriteHeader(http.StatusOK)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	p := io.MultiReader(stdout, stderr)
	cmd.Start()
	io.Copy(w, p)
}

func apiwsCall(w *websocket.Conn) {
	cmd := exec.Command("python", "-u", "func.py")
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	p := io.MultiReader(stdout, stderr)
	cmd.Start()
	io.Copy(w, p)
	w.Close()
}