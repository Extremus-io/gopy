package main

import (
	"testing"
	"time"
	"net/http"
	"github.com/Extremus-io/gopy/machines"
	"encoding/json"
	"os"
	"fmt"
)

func TestApi(t *testing.T) {
	go main()
	time.Sleep(time.Second)
	resp, err := http.Get("http://" + *host + "/api/machines/")
	if err != nil {
		t.Error("unable to send request to server")
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Errorf("http returned status %d", resp.StatusCode)
		t.Fail()
	}
	mcs := []machines.MachineInfo{}
	d := json.NewDecoder(resp.Body)
	err = d.Decode(&mcs)
	if err != nil {
		t.Error("unable to decody body")
		t.Fail()
	}
	if len(mcs) != 0 {
		t.Error("Something connected without. nothing was initiated")
		t.Fail()
	}
	h, _ := os.Hostname()
	mi := machines.MachineInfo{
		Hostname:h,
		Group:"hello/",
	}
	m, err := machines.ClientConnect(*host, mi)
	if err != nil {
		t.Errorf("unable to connect error %s", err.Error())
	}
	resp, err = http.Get("http://" + *host + "/api/machines/")
	if err != nil {
		t.Error("unable to send request to server")
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Errorf("http returned status %d", resp.StatusCode)
		t.Fail()
	}
	d = json.NewDecoder(resp.Body)
	err = d.Decode(&mcs)
	if err != nil {
		t.Error("unable to decody body")
		t.Fail()
	}
	if len(mcs) != 1 {
		t.Error("client connecte but not showing up. nothing was initiated")
		t.Fail()
	}
	inf, _ := m.Info()
	if mcs[0] != inf {
		t.Error("connected something and stored something")
		t.Errorf("connected : %v\n Stored: %v\n", inf, mcs[0])
	}
	u := fmt.Sprintf("http://" + *host + "/api/machines/%d", inf.Id)
	resp, err = http.Get(u)
	if err != nil {
		t.Error("unable to send request to server")
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Errorf("http returned status %d", resp.StatusCode)
		t.Fail()
	}
	d = json.NewDecoder(resp.Body)
	err = d.Decode(&inf)
	if err != nil {
		t.Error("unable to decody body")
		t.Fail()
	}
	i, _ := m.Info()
	if i != inf {
		t.Error("wrong data received with url /api/machine/{id}")
		t.Fail()
	}
}
