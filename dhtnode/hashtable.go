package dhtnode

import (
	"sync"
	"strconv"
)

const numLocks = 4096

type HashTable struct{
	data []int
	lock sync.Mutex
	locks []sync.Mutex
}

func (ls *HashTable) Init() {
	ls.data = make([]int, keyRange, 1<<32)
	ls.locks = make([]sync.Mutex, numLocks)
}

func (ls *HashTable) Put(key, value int) (int, int) {
	ls.locks[key % numLocks].Lock()
	defer ls.locks[key % numLocks].Unlock()
	//ls.lock.Lock()
	//defer ls.lock.Unlock()
	v := ls.data[key]
	if v != 1<<32 {
		return v, 1
	} else {
		ls.data[key] = value
		return value, 2
	}
}

func (ls *HashTable) Get(key int) (int, int) {
	ls.locks[key % numLocks].Lock()
	defer ls.locks[key % numLocks].Unlock()
	//ls.lock.Lock()
	//defer ls.lock.Unlock()
	value := ls.data[key]
	if value != 1<<32 {
		return value, 2
	} else {
		return value, 0
	}
}

func (ls *HashTable) String() string {
	//ls.lock.Lock()
	//defer ls.lock.Unlock()
	var s string
	for k,v := range ls.data {
		s += "( " + strconv.Itoa(k) + " , " + strconv.Itoa(v) + ")\n"
	}
	return s
}

func (ls *HashTable) Clear() {
	//ls.lock.Lock()
	//defer ls.lock.Unlock()
	ls.data = make([]int, keyRange, 1<<32)
}

func (ls *HashTable) Size() int {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	return len(ls.data)
}