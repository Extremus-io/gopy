package main

import (
	"net/http"
	"github.com/Extremus-io/gopy/log"
	"time"
)

func init() {
	http.HandleFunc("/api/machines/", cros(logWrap(machineApi)))
}

func cros(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		h(w, r)
	}

}

func logWrap(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		init := time.Now()
		h(w, r)
		diff := time.Since(init)
		log.Infof("%s\t%s\t%s", r.Method, r.URL.Path, diff)
	}
}