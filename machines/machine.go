package machines

import (
	"io"
	"math/rand"
	"encoding/json"
	"golang.org/x/net/websocket"
	"time"
	"github.com/Extremus-io/gopy/log"
	"errors"
)

const MAX_RETRIES = 20

// fill these parameters and use this to make a new config
type MachineInfo struct {
	Id        int                  `json:"id"`
	Hostname  string               `json:"hostname"`
	Extra     json.RawMessage      `json:"extra"`
	Group     string               `json:"group"`
	ConnectAt time.Time            `json:"connected_at"`
}

// functional reader writer interface for communicating with machine
type Machine struct {
	Id     int                      `json:"id"`
	ws     *websocket.Conn
	reader io.Reader
	writer io.Writer
}


// Generates new Machine variable and stores it into local map
func NewMachineFromWs(c MachineInfo, ws *websocket.Conn) (*Machine, error) {
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
			log.Critical("rand int is not generating unique int entry for id")
			return nil, errors.New("ID generation failed. Exceed max retries")
		}
	}
	_, err := machine_ins.Exec(id, c.Hostname, []byte(c.Extra), c.Group, c.ConnectAt)
	if err != nil {
		log.Critical("Failed to store config to db")
		return nil, err
	}

	// making a machine and saving it into data
	m := &Machine{
		Id:id,
		ws:ws,
		reader:ws,
		writer:ws,
	}
	data[id] = m

	// returning a machine
	return m, nil
}

func DeleteMachine(id int) {
	// lock the data map for concurrency protection
	lock.Lock()
	defer lock.Unlock()
	delete(data, id)
	machine_del_by_id.Exec(id)
	log.Verbosef("machine id %d deleted successfully", id)
}
func (m *Machine) Info() (MachineInfo, error) {
	row := machine_sel_by_id.QueryRow(m.Id)
	c := MachineInfo{}
	err := row.Scan(&c.Id, &c.Hostname, &c.Extra, &c.Group, &c.ConnectAt)
	return c, err
}

func (m *Machine) Read(p []byte) (int, error) {
	return m.reader.Read(p)
}

func (m *Machine) Write(p []byte) (int, error) {
	return m.writer.Write(p)
}
