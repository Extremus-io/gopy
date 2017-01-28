package cli

import (
	"io"
	"encoding/json"
	"github.com/Extremus-io/gopy/log"
	"errors"
	"golang.org/x/net/websocket"
)

var ErrAlreadyRunning = errors.New("Error cli already started")
var ErrCommandNotFound = errors.New("Error requested command not found")

type Msg struct {
	Type string `json:"type"`
	Msg  json.RawMessage `json:"data"`
}

func (m *Msg) Execute() error {
	f, found := commands[m.Type]
	if !found {
		return ErrCommandNotFound
	}
	return f([]byte(m.Msg))
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
	enc := json.NewEncoder(c.writer)
	if c.running == true {
		return ErrAlreadyRunning
	}
	c.running = true
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
		err = msg.Execute()
		if err != nil {
			enc.Encode(map[string]string{
				"result":"failed",
				"reason":err.Error(),
			})
			continue
		}
		enc.Encode(map[string]string{
			"result":"success",
		})

	}
	return nil
}

func (c *Cli) Close() error {
	return c.close.Close()
}

func (c *Cli) Running() bool {
	return c.running
}