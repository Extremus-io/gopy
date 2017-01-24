package main

import (
	"net/http"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"github.com/jroimartin/gocui"
)

func exec(cmd string, extra string) error {
	mainView, err := gui.View("main")
	if err != nil {
		panic(err)
	}
	switch cmd {
	case "\\m":
		resp, err := http.Get("http://" + *host + "/api/machines/" + extra)
		if err != nil {
			fmt.Fprint(mainView, color.RedString("error occured while making request\n"))
			fmt.Fprintf(mainView, color.RedString("error: %s\n", err.Error()))
			return nil
		}
		var b []byte
		b, _ = ioutil.ReadAll(resp.Body)
		fmt.Fprint(mainView, color.CyanString("Success: %s\n", b))
	case "exit":
		return gocui.ErrQuit
	}
	return nil

}
