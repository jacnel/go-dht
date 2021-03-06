package dhtnode

import (
	"sync"
	"strconv"
)

const numLocks = keyRange

type HashTable struct{
	data []int
	lock sync.Mutex
	// look into Actor model
	locks []sync.Mutex
}

func (table *HashTable) Init() {
	table.data = make([]int, keyRange)
	for i := range table.data {
		table.data[i] = 1<<32
	}
	table.locks = make([]sync.Mutex, numLocks)
}

func (table *HashTable) Put(key, value int) (int, int) {
	table.locks[key % numLocks].Lock()
	defer table.locks[key % numLocks].Unlock()
	v := table.data[key]
	if v != 1<<32 {
		return v, 1
	} else {
		table.data[key] = value
		return value, 2
	}
}

func (table *HashTable) Get(key int) (int, int) {
	table.locks[key % numLocks].Lock()
	defer table.locks[key % numLocks].Unlock()
	value := table.data[key]
	if value != 1<<32 {
		return value, 2
	} else {
		return value, 0
	}
}

func (table *HashTable) String() string {
	for i := range table.locks {
		table.locks[i].Lock()
	}
	var s string
	for k,v := range table.data {
		s += "( " + strconv.Itoa(k) + " , " + strconv.Itoa(v) + ")\n"
	}
	for i := range table.locks {
		table.locks[i].Unlock()
	}
	return s
}

func (table *HashTable) Clear() {
	for i := range table.locks {
		table.locks[i].Lock()
	}
	table.data = make([]int, keyRange)
	for i := range table.data {
		table.data[i] = 1<<32
	}
	for i := range table.locks {
		table.locks[i].Unlock()
	}
}

func (table *HashTable) Size() int {
	for i := range table.locks {
		table.locks[i].Lock()
	}
	size := 0
	for _, v := range table.data {
		if v != 1<<32 {
			size++
		}
	}
	for i := range table.locks {
		table.locks[i].Unlock()
	}
	return size
}