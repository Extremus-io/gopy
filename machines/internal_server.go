package machines

import (
	"net/http"
	"golang.org/x/net/websocket"
	"encoding/json"
	"time"
	"errors"
	"fmt"
)

const _MACHINE_INT_HANDSHAKE_TIMEOUT = time.Second * 10

func SetupServer() {
	http.Handle("/_ah/machine/connect", websocket.Handler(handleWebsocket))
}

func handleWebsocket(ws *websocket.Conn) {
	// for writing direct objects into ws use these
	decoder := json.NewDecoder(ws)
	encoder := json.NewEncoder(ws)

	// complete handshake and get MachineConfig
	mc, err := wsHandshake(ws, decoder)

	// if raise error if handshake was unsuccessful
	if err != nil {
		encoder.Encode(map[string]string{"handshake":false, "error":err})
		return
	}

	// register machine and send id to complete handshake
	cl := make(chan error)
	m := NewMachineFromWs(mc, ws, cl)
	defer DeleteMachine(m.Id)

	if err != nil {
		return
	}

	err, ok := <-cl
	if ok {
		encoder.Encode(map[string]string{

		})
	}

}

func wsHandshake(ws *websocket.Conn, decoder json.Decoder) (*MachineConfig, error) {
	c := make(chan interface{})
	var err error
	mc := MachineConfig{}
	go func() {
		defer close(c)
		err = decoder.Decode(&mc)
	}()

	select {
	case <-c:
		break
	case <-time.After(_MACHINE_INT_HANDSHAKE_TIMEOUT):
		err = errors.New(fmt.Sprintf("error Hand shake timed out. timeout value `%v`", _MACHINE_INT_HANDSHAKE_TIMEOUT))
	}

	return mc, err
}