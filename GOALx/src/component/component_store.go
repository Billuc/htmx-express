package component

import "sync"

type ComponentStore struct {
	Components []string
}

var componentStore *ComponentStore
var lock = &sync.Mutex{}

func GetComponentStore() *ComponentStore {
	if componentStore == nil {
		lock.Lock()
		defer lock.Unlock()

		if componentStore == nil {
			componentStore = &ComponentStore{}
		}
	}

	return componentStore
}
