package main

import (
	"net/http"
	"github.com/Extremus-io/gopy/log"
	"time"
)

func init() {
	http.HandleFunc("/machines", logWrap(machineApi))
}

func logWrap(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		init := time.Now()
		h(w, r)
		diff := time.Now() - init
		log.Infof("%s\t%s\t%s", r.Method, r.URL.Path, diff)
	}
}