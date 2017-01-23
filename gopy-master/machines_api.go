package main

import (
	"net/http"
	"regexp"
	"strconv"
	"github.com/Extremus-io/gopy/machines"
	"encoding/json"
	"github.com/Extremus-io/gopy/log"
)

var machine_url = regexp.MustCompile("^/api/machines/([0-9]+)$")

func machineApi(w http.ResponseWriter, r *http.Request) {

	// Extracting id (if any) from the request url
	// following pattern will match only if the url is requesting for specific machine
	match := machine_url.FindAllStringSubmatch(r.URL.Path, -1)
	var id = -1
	log.Verbosef("machine_url matched following %v", match)
	if len(match) == 1 {
		_id, _ := strconv.ParseInt(match[0][1], 10, strconv.IntSize)
		id = int(_id)
	}

	switch r.Method {
	case http.MethodGet:
		if id >= 0 {
			// It means that the request is for specific id
			mach, found := machines.GetMachineInfo(id)
			if !found {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Not Found!!"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mach)
			return
		} else {
			// it means the request if for all the machines
			mach := machines.GetAllMachinesInfo()
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(mach)
			if err != nil {
				log.Criticalf("Unable to encode machine info reason %v", err)
			}
		}
	case http.MethodPost:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Not yet implemented"))
	}
}