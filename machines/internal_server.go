package machines

import (
	"net/http"
	"golang.org/x/net/websocket"
	"encoding/json"
	"time"
	"errors"
	"fmt"
	"github.com/Extremus-io/gopy/log"
)

const _MACHINE_INT_HANDSHAKE_TIMEOUT = time.Second * 10

func SetupServer() {
	http.Handle("/_ah/machine/connect", websocket.Handler(handleWebsocket))
}

func handleWebsocket(ws *websocket.Conn) {
	log.Debug("slave connect requested")

	// for writing direct objects into ws use these
	decoder := json.NewDecoder(ws)
	encoder := json.NewEncoder(ws)

	// complete handshake and get MachineConfig
	mc, err := wsHandshake(ws, decoder)

	// if raise error if handshake was unsuccessful
	if err != nil {
		log.Errorf("slave connect failed error:`%v`", err)
		encoder.Encode(map[string]interface{}{"handshake":false, "error":err.Error()})

		return
	}
	// register machine and send id to complete handshake
	m, err := NewMachineFromWs(mc, ws)
	if err != nil {
		log.Errorf("slave connect failed error:`%v`", err)
		encoder.Encode(map[string]interface{}{"handshake":false, "error":err.Error()})

		return
	}

	defer DeleteMachine(m.Id)

	// finish the handshake step
	encoder.Encode(map[string]interface{}{
		"handshake":true,
		"id":m.Id,
	})
	log.Infof("slave connect success. hostname:%s Id:%d", mc.Hostname, m.Id)

	var t = make(map[string]interface{})
	// copy data from ws to buffer until ws is closed or encountered error
	for err := decoder.Decode(&t); err == nil; err = decoder.Decode(&t) {
		log.Criticalf("YTI -- got message from machine %v", t)
	}

	log.Infof("slave disconnected. hostname:%s Id:%d", mc.Hostname, m.Id)
}

func wsHandshake(ws *websocket.Conn, decoder *json.Decoder) (MachineConfig, error) {
	c := make(chan interface{})
	var err error
	mc := MachineConfig{}
	go func() {
		defer close(c)
		err = decoder.Decode(&mc)
	}()

	select {
	case _, _ = <-c:
		break
	case <-time.After(_MACHINE_INT_HANDSHAKE_TIMEOUT):
		err = errors.New(fmt.Sprintf("error Hand shake timed out. timeout value `%v`", _MACHINE_INT_HANDSHAKE_TIMEOUT))
	}
	mc.ConnectAt = time.Now()
	return mc, err
}