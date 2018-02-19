package main

import (
	"fmt"
	"dht/dhtclient"
	rand2 "math/rand"
	"os"
	"strconv"
	"time"
)

func doRandomWork(i, targetNode, numOps, keyRange int, ops, puts *[]int, runtimes *[]float64, done chan bool) {
	c := dhtclient.DHTClient{}
	switch targetNode % 4 {
	case -1:
		c.Init("127.0.0.1:8403")
		break
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
	fmt.Println("Target node ->", targetNode % 4)
	myPuts := 0
	myOps := 0
	var nanos int64
	now := time.Now()
	rand2.Seed(time.Since(now).Nanoseconds())
	for j := 0; j < numOps; j++ {
		if j % 50 == 0 {
			fmt.Printf(".")
		}
		r := rand2.Intn(keyRange)
		if r < int(float64(keyRange) * .4) {
			start := time.Now()
			_, ok := c.Put(r, i)
			elapsed := time.Since(start)
			nanos += elapsed.Nanoseconds()
			myOps++
			if ok == 2 {
				myPuts++
			}
		} else {
			start := time.Now()
			c.Get(r)
			elapsed := time.Since(start)
			nanos += elapsed.Nanoseconds()
			myOps++
		}
	}
	(*ops)[i] = myOps
	(*puts)[i] = myPuts
	(*runtimes)[i] = float64(nanos) / 1000000

	done <- true
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

	runtimes := make([]float64, numClients)
	ops := make([]int, numClients)
	puts := make([]int, numClients)
	done := make(chan bool, numClients)

	// Start clients
	fmt.Println("Spawning new clients")
	for i := 0; i < numClients; i++ {
		if targetNode >= 0 {
			go doRandomWork(i, targetNode, numOps, keyRange, &ops, &puts, &runtimes, done)
		} else {
			go doRandomWork(i, i, numOps, keyRange, &ops, &puts, &runtimes, done)
		}
	}
	fmt.Println("Done")

	// Wait for go routines to finish
	for i := 0; i < numClients; i++ {
		<- done
	}

	// Calculate total successful ops
	totalOps := 0
	for i := 0; i < numClients; i++ {
		o := ops[i]
		ops = append(ops, o)
		totalOps += o
	}
	fmt.Printf("Total operations: %d ops\n", totalOps)

	// Calculate total runtime
	throughputs := make([]float64, 4)
	for i := 0; i < numClients; i++ {
		rt := runtimes[i]
		runtimes = append(runtimes, rt)
		throughputs[i % 4] += float64(ops[i]) / rt * 1000
	}

	totalThroughput := 0.0
	for _, t := range throughputs {
		totalThroughput += t
	}

	fmt.Printf("System throughput: %4.2f ops/s\n", totalThroughput)

	latency := 0.0
	for i := 0; i < numClients; i++ {
		latency += runtimes[i] / float64(ops[i])
	}
	latency /= float64(numClients)
	fmt.Printf("Average latency %4.2f ms\n", latency)

	// Sanity check
	totalPuts := 0
	for i := 0; i < numClients; i++ {
		totalPuts += puts[i]
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
	}
	c.Close()
	return totalSize
}