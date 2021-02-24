package utils

import "sync"

//safe map
type SMap struct {
	m sync.Mutex
	v map[string]interface{}
}

func NewCMap() *SMap {
	return &SMap{
		v: make(map[string]interface{}),
	}
}

func (sm *SMap) Set(key string, value interface{}) {
	sm.m.Lock()
	sm.v[key] = value
	sm.m.Unlock()
}

func (sm *SMap) Get(key string) interface{} {
	sm.m.Lock()
	val := sm.v[key]
	sm.m.Unlock()
	return val
}

func (sm *SMap) Has(key string) bool {
	sm.m.Lock()
	_, ok := sm.v[key]
	sm.m.Unlock()
	return ok
}

func (sm *SMap) Delete(key string) {
	sm.m.Lock()
	delete(sm.v, key)
	sm.m.Unlock()
}

func (sm *SMap) Size() int {
	sm.m.Lock()
	size := len(sm.v)
	sm.m.Unlock()
	return size
}

func (sm *SMap) Clear() {
	sm.m.Lock()
	sm.v = make(map[string]interface{})
	sm.m.Unlock()
}

func (sm *SMap) Keys() []string {
	sm.m.Lock()

	keys := []string{}
	for k := range sm.v {
		keys = append(keys, k)
	}
	sm.m.Unlock()
	return keys
}

func (sm *SMap) Values() []interface{} {
	sm.m.Lock()
	items := []interface{}{}
	for _, v := range sm.v {
		items = append(items, v)
	}
	sm.m.Unlock()
	return items
}
