package dhtnode

import (
	"sync"
	"strconv"
)

type LocalStore struct{
	data map[int]int
	lock sync.Mutex
}

func (ls *LocalStore) Init() {
	ls.data = make(map[int]int)
}

func (ls *LocalStore) Put(key, value int) (int, int) {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	v, exists := ls.data[key]
	if exists {
		return v, 1
	} else {
		ls.data[key] = value
		return value, 2
	}
}

func (ls *LocalStore) Get(key int) (int, int) {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	value, exists := ls.data[key]
	if exists {
		return value, 2
	} else {
		return value, 0
	}
}

func (ls *LocalStore) String() string {
	var s string
	for k,v := range ls.data {
		s += "( " + strconv.Itoa(k) + " , " + strconv.Itoa(v) + ")\n"
	}
	return s
}

func (ls *LocalStore) Clear() {
	ls.data = make(map[int]int)
}