package machines

import (
	"golang.org/x/net/websocket"
	"encoding/json"
	"time"
	"errors"
	"fmt"
)

const SERVER_URL = "ws://localhost:9000/_ah/machine/connect"
const CONNECT_TIMEOUT = time.Second * 2

func clientConnect(mc MachineConfig) (*Machine, error) {

	// Connect to server
	ws, err := websocket.Dial(SERVER_URL, "", "http://localhost:9000")
	if err != nil {
		return nil, err
	}

	// use these for encoding or decoding json data
	enc := json.NewEncoder(ws)
	dec := json.NewDecoder(ws)

	// Complete Handshake
	Id, err := cliHandshake(mc, enc, dec)

	// if no error make machine object and return
	if err != nil {
		return nil, err
	}
	m := &Machine{
		Id:Id,
		reader:ws,
		writer:ws,
		Conf:mc,
		ws:ws,
	}
	return m, nil
}

func cliHandshake(mc MachineConfig, enc *json.Encoder, dec *json.Decoder) (int, error) {
	// Initiate handshake
	enc.Encode(mc)
	var err error

	// datatype for storing info
	a := struct {
		Handshake bool        `json:"handshake"`
		Error     string      `json:"error"`
		Id        int         `json:"id"`
	}{}

	// decode the response with a timeout
	c := make(chan int)
	go func() {
		err = dec.Decode(&a)
		c <- 1
	}()
	select {
	case _, _ = <-c:
		if err != nil {
			return 0, err
		}
		break
	case <-time.After(CONNECT_TIMEOUT):
		return 0, errors.New("handshake timeout. No response received")
	}

	// After the response is stored, find
	if !a.Handshake {
		return 0, errors.New(fmt.Sprintf("Unable to complete Handshake error:`%s`", a.Error))
	}
	return a.Id, nil
}