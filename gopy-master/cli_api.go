package main

import (
	"golang.org/x/net/websocket"
	"github.com/Extremus-io/gopy/cli"
	"github.com/Extremus-io/gopy/log"
)

func cliWsApi(ws *websocket.Conn) {
	client := cli.NewWsCli(ws)
	log.Notice("Client connected")
	client.Start()
	log.Notice("Client disconnected")
}