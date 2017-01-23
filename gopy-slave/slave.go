package main

import (
	"github.com/Extremus-io/gopy/machines"
	"flag"
	"os"
	"github.com/Extremus-io/gopy/log"
	"io"
)

var defHostName, _ = os.Hostname()
var server_host = flag.String("host", "127.0.0.1:8765", "Server address of the master. " +
	"Note bad format can make the process go into an infinite loop")
var hostname = flag.String("hname", defHostName, "used to identify this machine")
var group = flag.String("group", "new_group", "used to identify specific group " +
	"of machines. eg:generic_crawler, post_processor")

func main() {
	flag.Parse()
	mi := machines.MachineInfo{
		Hostname:*hostname,
		Group:*group,
	}
	log.Infof("Connecting to server %s", *server_host)
	m, err := machines.ClientConnect(*server_host, mi)
	if err != nil {
		panic(err)
	}
	log.Infof("Successfully connected hostname:%s id:%d", *hostname, m.Id)
	io.Copy(os.Stdin, m)
	if err != nil {
		log.Errorf("Exiting process. error:%s", err.Error())
	}
}

