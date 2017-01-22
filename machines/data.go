package machines

import (
	"sync"
	"database/sql"
	"github.com/Extremus-io/gopy/log"
)

var lock = sync.RWMutex{}
var data = make(map[int]*Machine)

func GetMachine(id int) (*Machine, bool) {
	lock.RLock()
	defer lock.RUnlock()
	return getMachine(id)
}
func getMachine(id int) (*Machine, bool) {
	machine, found := data[id]
	return machine, found
}

func GetMachineInfo(id int) (MachineConfig, bool) {
	mc := MachineConfig{}
	row := machine_sel_by_id.QueryRow(id)
	err := row.Scan(&mc.Id, &mc.Hostname, &mc.Group, &mc.Extra, &mc.ConnectAt)
	if err == sql.ErrNoRows {
		return mc, false
	}
	if err != nil {
		log.Debugf("Requested machine id `%d` DB query error occured:%s", id, err.Error())
		return mc, false
	}
	return mc, true
}

func GetAllMachinesInfo() []MachineConfig {
	mcs := []MachineConfig{}
	row, err := machine_sel_all.Query()
	defer row.Close()

	if err != nil {
		panic(err)
	}

	for row.Next() {
		mc := MachineConfig{}
		err := row.Scan(&mc.Id, &mc.Hostname, &mc.Group, &mc.Extra, &mc.ConnectAt)
		if err != nil {
			panic(err)
		}
		mcs = append(mcs, mc)
	}

	return mcs
}