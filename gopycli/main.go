package main

import (
	"flag"
	"golang.org/x/net/websocket"
)

var host = flag.String("host", "localhost:8765", "server address")
var ws *websocket.Conn

func main() {
	var err error
	ws, err = websocket.Dial("ws://" + *host + "/api/cli/ws", "", "http://" + *host)
	if err != nil {
		panic(err)
	}
	RunUi()
}