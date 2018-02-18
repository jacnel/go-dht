package main

import (
	dhtnode "dht/dhtnode"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide a portnumber to start the node.")
	}
	arg := os.Args[1]
	var port int
	if strings.Compare(arg, "") != 0 {
		var err error
		port, err = strconv.Atoi(arg)
		if err != nil {
			panic("Invalid portnumber... exiting.")
		}
	}
	dhtNode := dhtnode.DHTNode{}
	dhtNode.StartNode(port, "./node/config.json")
}