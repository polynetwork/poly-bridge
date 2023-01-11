package set

import (
	"bytes"
	"fmt"
	"sync"
)

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

func (set *HashSet) Remove(inter interface{}) {

	set.Lock()

	defer set.Unlock()

	delete(set.m, inter)

}

func (set *HashSet) Clear() {

	set.Lock()

	defer set.Unlock()

	set.m = make(map[interface{}]bool)

}

func (set *HashSet) Len() int {

	return len(set.m)

}
func (set *HashSet) IsSame(other *HashSet) bool {

	if other == nil {

		return false

	}

	if set.Len() != other.Len() {

		return false

	}

	for k, _ := range set.m {

		if !other.Contains(k) {

			return false

		}

	}

	return true

}

func (set *HashSet) String() string {

	var buf bytes.Buffer

	buf.WriteString("set{")

	first := true

	for k, _ := range set.m {

		if first {

			first = false

		} else {

			buf.WriteString(" ")

		}

		buf.WriteString(fmt.Sprintf("%v", k))

	}

	buf.WriteString("}")

	return buf.String()

}
