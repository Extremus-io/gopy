package main

import (
	"os/exec"
	"encoding/json"
	"sync"
	"os"
	"io"
	"fmt"
)

type Cmd struct {
	Args []string `json:"args"`
}

func main() {
	c := exec.Command("python", "-u", "gopy.py")
	i, _ := c.StdoutPipe()

	in := json.NewDecoder(i)
	c.Start()
	var p Cmd
	w := sync.WaitGroup{}
	for err := in.Decode(&p); err == nil; err = in.Decode(&p) {
		fmt.Print(p)
		w.Add(1)
		go func() {
			cm := exec.Command("python", p.Args...)
			k, _ := cm.StdoutPipe()
			cm.Start()
			io.Copy(os.Stdout, k)
			w.Done()
		}()
	}
	c.Wait()
	w.Wait()
}