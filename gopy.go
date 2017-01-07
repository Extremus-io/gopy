package main

import "net/http"

var (
	host = "127.0.0.1:9000"
	webroot = "webapp/"
)

func main() {
	http.ListenAndServe(host, http.FileServer(http.Dir(webroot)))
}
