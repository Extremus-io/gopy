package machines

import (
	"io"
	"math/rand"
	"encoding/json"
	"golang.org/x/net/websocket"
	"time"
)

const MAX_RETRIES = 20

// fill these parameters and use this to make a new config
type MachineConfig struct {
	Hostname  string               `json:"hostname"`
	Threads   int                  `json:"threads"`
	Extra     json.RawMessage      `json:"extra"`
	ConnectAt time.Time            `json:"connected_at"`
}

// functional reader writer interface for communicating with machine
type Machine struct {
	Id            int                      `json:"id"`
	Conf          MachineConfig            `json:"conf"`
	ws            *websocket.Conn
	reader        io.Reader
	writer        io.Writer
}


// Generates new Machine variable and stores it into local map
func NewMachineFromWs(c MachineConfig, ws *websocket.Conn,reader io.Reader) *Machine {
	// lock the data file for safely updating map
	lock.Lock()
	defer lock.Unlock()

	// getting unique id
	var id int
	var found bool
	var retryCount = 0
	for ; ; retryCount++ {
		id = rand.Int()
		_, found = getMachine(id)
		if !found {
			break
		}
		if retryCount >= MAX_RETRIES {
			panic("rand int is not generating unique int entry for id")
		}
	}

	// making a machine and saving it into data
	m := &Machine{
		Id:id,
		Conf:c,
		ws:ws,
		reader:reader,
		writer:ws,
	}
	data[id] = m

	// returning a machine
	return m

}

func DeleteMachine(id int) {
	// lock the data map for concurrency protection
	lock.Lock()
	defer lock.Unlock()

	delete(data, id)
}

func (m *Machine) Read(p []byte) (int, error) {
	return m.reader.Read(p)
}

func (m *Machine) Write(p []byte) (int, error) {
	return m.writer.Write(p)
}
