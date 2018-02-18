package main

import (
	"fmt"
	"dht/dhtclient"
	rand2 "math/rand"
	"os"
	"strconv"
	//"time"
	"time"
)

func doRandomWork(i, numOps, keyRange int, opsChan, putsChan chan int, runtimeChan chan float64) {
	c := dhtclient.DHTClient{}
	switch i % 4 {
	case 0:
		c.Init("54.208.29.162:8403")
		break
	case 1:
		c.Init("54.211.127.45:8403")
		break
	case 2:
		c.Init("75.101.226.165:8403")
		break
	case 3:
		c.Init("34.233.120.248:8403")
		break
	default:
		return
	}
	puts := 0
	ops  := 0
	var nanos int64
	for j := 0; j < numOps; j++ {
		if j % 50 == 0 {
			fmt.Println("client:", i, "j:", j, "puts:", puts)
		}
		r := rand2.Intn(keyRange)
		if r < int(float64(keyRange) * .4) {
			start := time.Now()
			_, ok := c.Put(r, i)
			elapsed := time.Since(start)
			nanos += elapsed.Nanoseconds()
			ops++
			if ok == 2 {
				puts++
			}
		} else {
			start := time.Now()
			c.Get(j)
			elapsed := time.Since(start)
			nanos += elapsed.Nanoseconds()
			ops++
		}
	}
	putsChan <- puts
	opsChan <- ops
	runtimeChan <- float64(nanos) / float64(1000000)

	c.Close()
}

func main() {
	if len(os.Args) < 4 {
		panic("Enter a number of clients to spin up, the number of operations per client and the key range.")
	}
	numClients,_ := strconv.Atoi(os.Args[1])
	numOps,_ := strconv.Atoi(os.Args[2])
	keyRange, _ := strconv.Atoi(os.Args[3])
	targetNode := -1
	if len(os.Args) > 4 {
		targetNode,_ = strconv.Atoi(os.Args[4])
	}

	putsChan := make(chan int, numClients)
	opsChan  := make(chan int, numClients)
	runtimeChan := make(chan float64, numClients)

	// Start clients
	fmt.Println("Spawning new clients")
	for i := 0; i < numClients; i++ {
		if targetNode >= 0 {
			go doRandomWork(targetNode, numOps, keyRange, opsChan, putsChan, runtimeChan)
		} else {
			go doRandomWork(i, numOps, keyRange, opsChan, putsChan, runtimeChan)
		}
	}
	fmt.Println("Done")

	// Calculate total successful ops
	totalOps := 0
	ops := make([]int, 0)
	for i := 0; i < numClients; i++ {
		o := <-opsChan
		ops = append(ops, o)
		totalOps += o
	}
	fmt.Printf("Total operations: %dops\n", totalOps)

	// Calculate total runtime
	runtimes := make([]float64, 0)
	throughput := 0.0
	for i := 0; i < numClients; i++ {
		rt := <-runtimeChan
		runtimes = append(runtimes, rt)
		throughput += float64(ops[i]) / rt * 1000
	}
	throughput /= float64(numClients)
	fmt.Printf("Average throughput: %4.2fops\n", throughput)

	latency := 0.0
	for i := 0; i < numClients; i++ {
		latency += runtimes[i] / float64(ops[i])
	}
	latency /= float64(numClients)
	fmt.Printf("Average latency %4.2fms\n", latency)

	// Sanity check
	totalPuts := 0
	for i := 0; i < numClients; i++ {
		totalPuts += <-putsChan
	}

	totalSize := getTotalSize()
	if totalPuts != totalSize {
		fmt.Printf("totalPuts(%d) != totalSize(%d) : [FAILED]\n", totalPuts, totalSize)
	} else {
		fmt.Printf("totalPuts(%d) == totalSize(%d) : [PASSED]\n", totalPuts, totalSize)
	}
}

func getTotalSize() int {
	c := dhtclient.DHTClient{}
	totalSize := 0
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			c.Init("54.208.29.162:8403")
			break
		case 1:
			c.Init("54.211.127.45:8403")
			break
		case 2:
			c.Init("75.101.226.165:8403")
			break
		case 3:
			c.Init("34.233.120.248:8403")
			break
		}
		totalSize += c.Size()
		c.Close()
	}
	return totalSize
}