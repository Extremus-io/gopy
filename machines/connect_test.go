package machines

import (
	"testing"
	"net/http"
	"time"
)

func TestSetupServer(t *testing.T) {
	SetupServer()
	go func() {
		http.ListenAndServe("localhost:9000", nil)
	}()
	machine_conf := MachineConfig{
		Hostname:"kittuov-lappy",
	}
	m, err := clientConnect(machine_conf)
	if err != nil {
		t.Error("cannot connect to server")
		t.Error(err)
		t.Fail()
	}
	ma, found := GetMachine(m.Id)
	if !found {
		t.Error("client connected but not registered on server")
		t.Fail()
	}
	if c, err := ma.Conf(); (c.Hostname != machine_conf.Hostname && err != nil) {
	t.Error("client connected but not registered correctly")
	t.Fail()
	}
	m.ws.Close()
	time.Sleep(time.Second)
	ma, found = GetMachine(m.Id)
	if found {
		t.Error("machine not de-registered even after disconnecting")
	}
}