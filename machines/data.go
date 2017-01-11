package machines

import "sync"

var lock = sync.RWMutex{}
var data map[int]*Machine

func GetMachine(id int) (*Machine, bool) {
	lock.RLock()
	defer lock.RUnlock()
	return getMachine(id)
}
func getMachine(id int) (*Machine, bool) {
	machine, found := data[id]
	return machine, found
}
