package cli

import (
	"io"
	"encoding/json"
	"github.com/Extremus-io/gopy/log"
	"fmt"
	"errors"
	"golang.org/x/net/websocket"
)

var ErrAlreadyRunning = errors.New("Error cli already started")

type Msg struct {
	Type string `json:"type"`
	Msg  json.RawMessage `json:"data"`
}

func (m *Msg) Execute() {
	fmt.Print("Executing")
}

type Cli struct {
	reader  io.Reader
	writer  io.Writer
	close   io.Closer
	running bool
}

func NewWsCli(ws *websocket.Conn) *Cli {
	return &Cli{
		reader:ws,
		writer:ws,
		close:ws,
		running:false,
	}
}

func (c *Cli) Start() error {
	dec := json.NewDecoder(c.reader)
	if c.running == true {
		return ErrAlreadyRunning
	}
	c.running = true
	go func() {
		defer func() {
			c.running = false
		}()
		for {
			msg := Msg{}
			err := dec.Decode(&msg)
			if err != nil {
				log.Errorf("Cli msg decoding error `%s`", err.Error())
				c.running = false
				break
			}
			msg.Execute()
		}
	}()
	return nil
}

func (c *Cli) Close() error {
	return c.close.Close()
}

func (c *Cli) Running() bool {
	return c.running
}