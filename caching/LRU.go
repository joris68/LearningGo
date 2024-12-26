package cache

type LRUCache interface {
	Get(key string) int
	Put(key string, value int)
	Keys() []string
	Remove(key string) bool
	Clear()
	Capacity() int
	Len() int
}

type Mycache struct {
	capacity int
	keysMap  map[string]int
	tail     ListEntry
	head     ListEntry
}

func (r Mycache) Get(key string) int {

	val, ok := r.keysMap[key]
	if ok {
		return val
	} else {

		return -1
	}
}
