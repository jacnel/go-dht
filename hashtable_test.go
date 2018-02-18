package dht

import (
     "testing"
     "dht/dhtclient"
     "fmt"
     rand2 "math/rand"
     "dht/dhtnode"
)

func TestNode1(t *testing.T) {
     n := dhtnode.DHTNode{}
     n.StartNode(8403, "../node/config.json")
}
//
//func TestNode2(t *testing.T) {
//     n := dhtnode.DHTNode{}
//     n.StartNode(8404)
//}
//
//func TestNode3(t *testing.T) {
//     n := dhtnode.DHTNode{}
//     n.StartNode(8405)
//}

func doRandomWork(c dhtclient.DHTClient, i int, ch chan int) {
     if i % 1 == 0 {
          c.Init("128.180.110.83:8403")
     } else {
          c.Init("128.180.145.134:8403")
     }
     puts := 0
     for j := 0; j < 10000; j++ {
          r := rand2.Intn(10000)
          if r < 4000 {
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

func TestClient(t *testing.T) {
     ch := make(chan int)
     numClients := 1
     for i := 0; i < numClients; i++ {
          fmt.Println("Spawning new client")
          c := dhtclient.DHTClient{}
          go doRandomWork(c, i, ch)
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

