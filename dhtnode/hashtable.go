package dhtnode

import (
	"sync"
	"strconv"
)

const numLocks = 50

type HashTable struct{
	data map[int]int
	locks []sync.Mutex
}

func (ls *HashTable) Init() {
	ls.data = make(map[int]int)
	ls.locks = make([]sync.Mutex, numLocks)
}

func (ls *HashTable) Put(key, value int) (int, int) {
	ls.locks[key % numLocks].Lock()
	defer ls.locks[key % numLocks].Unlock()
	v, exists := ls.data[key]
	if exists {
		return v, 1
	} else {
		ls.data[key] = value
		return value, 2
	}
}

func (ls *HashTable) Get(key int) (int, int) {
	ls.locks[key % numLocks].Lock()
	defer ls.locks[key % numLocks].Unlock()
	value, exists := ls.data[key]
	if exists {
		return value, 2
	} else {
		return value, 0
	}
}

func (ls *HashTable) String() string {
	var s string
	for k,v := range ls.data {
		s += "( " + strconv.Itoa(k) + " , " + strconv.Itoa(v) + ")\n"
	}
	return s
}

func (ls *HashTable) Clear() {
	ls.data = make(map[int]int)
}

func (ls *HashTable) Size() int {
	return len(ls.data)
}