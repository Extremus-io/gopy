package main

import (
	"net/http"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"github.com/jroimartin/gocui"
	"strconv"
	"encoding/json"
)

const create_user_usage = `
Usage:
	create_user <email> <password> <is_superuser>
`

func exec(cmd ...string) error {
	mainView, err := gui.View("main")
	if err != nil {
		panic(err)
	}
	switch cmd[0] {
	case "create_user":
		if len(cmd) != 4 {
			fmt.Fprint(mainView, color.YellowString(create_user_usage))
		}
		is_superuser, _ := strconv.ParseBool(cmd[3])
		user := struct {
			Email       string `json:"email"`
			Password    string `json:"password"`
			IsSuperUser bool `json:"is_superuser"`
		}{
			Email:cmd[1],
			Password:cmd[2],
			IsSuperUser:is_superuser,
		}
		json.NewEncoder(ws).Encode(map[string]interface{}{
			"type":"create_user",
			"msg":user,
		})
		result := struct {
			Result string `json:"result"`
		}{}
		err = json.NewDecoder(ws).Decode(&result)
		if err != nil {
			fmt.Fprint(mainView, color.RedString("Error:%s\n", err.Error()))
		}
		fmt.Fprint(mainView, color.GreenString("Result:%s\n", result.Result))
		break

	case "\\m":
		extra := ""
		if len(cmd) > 1 {
			extra = cmd[1]
		}
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
