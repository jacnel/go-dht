package dht


import (
     "testing"
     "dht/dhtclient"
     "dht/dhtnode"
     "strconv"
     "fmt"
     rand2 "math/rand"
)

func TestNode1(t *testing.T) {
     n := dhtnode.DHTNode{}
     n.StartNode(8403)
}

func TestNode2(t *testing.T) {
     n := dhtnode.DHTNode{}
     n.StartNode(8404)
}

func TestNode3(t *testing.T) {
     n := dhtnode.DHTNode{}
     n.StartNode(8405)
}

func doRandomWork(i int, ch chan int) {
     c := dhtclient.DHTClient{}
     defer c.Close()
     c.Init("128.180.110.83:840"+strconv.Itoa(i % 3 + 3))
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

func TestClient(t *testing.T) {
     ch := make(chan int)
     numClients := 3
     for i := 0; i < numClients; i++ {
          fmt.Println("Spawing new client")
          go doRandomWork(i, ch)
          //go doSameWork(i, ch)
     }
     var puts []int
     for i := 0; i < numClients; i++ {
          puts = append(puts, <-ch)
     }
     fmt.Println(puts)
}

func doSameWork(i int, ch chan int) {
     c := dhtclient.DHTClient{}
     defer c.Close()
     c.Init("128.180.110.83:840"+strconv.Itoa(i % 3 + 3))
     puts := 0
     for j := 0; j < 1000; j++ {

     }
     ch <- puts
}

