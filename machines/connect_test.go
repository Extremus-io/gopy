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
	machine_conf := MachineInfo{
		Hostname:"kittuov-lappy",
	}
	m, err := SlaveConnect("localhost:9000", machine_conf)
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
	if c, err := ma.Info(); (c.Hostname != machine_conf.Hostname && err != nil) {
		t.Error("client connected but not registered correctly")
		t.Fail()
	}
	m.ws.Close()
	time.Sleep(time.Second)
	ma, found = GetMachine(m.Id)
	if found {
		t.Error("machine not de-registered even after disconnecting")
		t.Fail()
	}
	_, err1 := SlaveConnect("localhost:9000", machine_conf)
	_, err2 := SlaveConnect("localhost:9000", machine_conf)

	if err1 != nil {
		t.Error("connect to server failed")
		t.Fail()
	}
	if err2 == nil {
		t.Error("attempted to make multiple connections. No error was thrown")
		t.Fail()
	}
}