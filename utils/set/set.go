package set

import "sync"

type HashSet struct {
	m map[interface{}]bool
	sync.RWMutex
}

func NewSetFromUint64(data []uint64) *HashSet {
	set := &HashSet{m: make(map[interface{}]bool)}
	for _, v := range data {
		set.Add(v)
	}
	return set
}

func (set *HashSet) Add(item interface{}) (b bool) {
	set.Lock()
	defer set.Unlock()
	if !set.m[item] {
		set.m[item] = true
		return true
	}
	return false
}

func (set *HashSet) Contains(item interface{}) bool {
	set.RLock()
	defer set.RUnlock()
	return set.m[item]
}
