package main

import (
	"github.com/Extremus-io/gopy/machines"
	"github.com/Extremus-io/gopy/log"
	"flag"
	"net/http"
)

var host = flag.String("host", "0.0.0.0:8765", "host:port for the server")

func init() {
	flag.Parse()
}
func main() {
	log.SetMin(log.DEBU)
	machines.SetupServer()
	log.Noticef("Starting Master server at %s", *host)
	err := http.ListenAndServe(*host, nil)
	log.Critical(err)
}