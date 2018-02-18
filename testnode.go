package main

import (
	"fmt"
	dhtclient "dht/dhtclient"
	rand2 "math/rand"
)

func doRandomWork(i int, ch chan int) {
	c := dhtclient.DHTClient{}
	defer c.Close()
	if i % 1 == 0 {
		c.Init("128.180.110.83:8403")
	} else {
		c.Init("128.180.145.134:8403")
	}
	puts := 0
	for j := 0; j < 100; j++ {
		r := rand2.Intn(100)
		if r < 40 {
			_, ok := c.Put(r, i)
			if ok == 2 {
				puts++
			}
		} else {
			c.Get(j)
		}
	}
	ch <- puts
}

func main() {
	ch := make(chan int)
	numClients := 1
	for i := 0; i < numClients; i++ {
		fmt.Println("Spawning new client")
		go doRandomWork(i, ch)
		//go doSameWork(i, ch)
	}
	var puts []int
	for i := 0; i < numClients; i++ {
		puts = append(puts, <-ch)
	}
	fmt.Println(puts)
	totalPuts := 0
	for i := range puts {
		totalPuts += puts[i]
	}
	fmt.Println(totalPuts)
}