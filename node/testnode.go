package main

import (
	"fmt"
	dhtclient "dht/dhtclient"
	rand2 "math/rand"
	"os"
	"strconv"
)

func doRandomWork(i, numOps int, ch chan int) {
	c := dhtclient.DHTClient{}
	if i % 1 != 0 {
		c.Init("128.180.110.83:8403")
	} else {
		c.Init("128.180.145.134:8403")
	}
	puts := 0
	for j := 0; j < numOps; j++ {
		if j % 10 == 0 {
			fmt.Println("client:", i, "j:", j, "puts:", puts)
		}
		r := rand2.Intn(100000)
		if r < 40000 {
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
	if len(os.Args) != 3 {
		panic("Enter a number of clients to spin up and the number of operations per client.")
	}
	numClients,_ := strconv.Atoi(os.Args[1])
	numOps,_ := strconv.Atoi(os.Args[2])

	ch := make(chan int)
	fmt.Println("Spawning new clients")
	for i := 0; i < numClients; i++ {
		go doRandomWork(i, numOps, ch)
	}
	fmt.Println("Done")
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