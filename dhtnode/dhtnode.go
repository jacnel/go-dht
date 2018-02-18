package dhtnode

import (
	"fmt"
	"strconv"
)

type DHTNode struct {
	hashTable HashTable
	network   Network
}

func (node DHTNode) StartNode(port int, configFilepath string) {
	node.init(port, configFilepath)
	node.Listen()
}

func (node *DHTNode) init(port int, configFilepath string) {
	node.hashTable = HashTable{}
	node.hashTable.Init()
	node.network = Network{}
	node.network.Init(port, configFilepath)
}

func (node *DHTNode) Listen() {
	node.network.Listen()
	for{
		node.network.Accept()
		go node.handleMessages()
	}
}

func (node *DHTNode) handleMessages() {
	var opcode, key, value int
	for {
		opcode, key, value = node.network.Receive()
		ok := -1
		switch opcode {
		case -1:
			fmt.Println("Closing connection...", node.network.conn)
			node.network.Send("")
			node.network.Close()
			return
		case 1:
			if node.network.KeyInRange(key){
				value, ok = node.hashTable.Get(key)
			} else {
				value, ok = node.network.LetsGoOffNoding(opcode, key, value)
			}
			fmt.Println("Get: (", key, ")")
			break
		case 2:
			if node.network.KeyInRange(key) {
				value, ok = node.hashTable.Put(key, value)
			} else {
				value, ok = node.network.LetsGoOffNoding(opcode, key, value)
			}
			fmt.Println("Put: (", key, ", ", value, ")")
			break
		case 3:
			node.network.Send(node.hashTable.String())
			fmt.Println("Current state:")
			fmt.Println(node.hashTable.String())
			break
		case 4:
			node.network.Send(strconv.Itoa(node.hashTable.Size()) + "\n")
			fmt.Println("Size: ", node.hashTable.Size())
			break
		case 5:
			node.hashTable.Clear()
			ok = 2
			fmt.Println("TABLE CLEARED!")
			break
		default:
			fmt.Println("default switch handle....",opcode)
			node.network.Send("IGNORED\n")
			fmt.Println("Ignoring message")
		}

		fmt.Println(ok, strconv.Itoa(key),strconv.Itoa(value))

		switch ok {
		case 0:
			node.network.Send("FAIL;"+strconv.Itoa(key)+";"+strconv.Itoa(value)+"\n")
			break
		case 1:
			node.network.Send("EXISTS;"+strconv.Itoa(key)+";"+strconv.Itoa(value)+"\n")
			break
		case 2:
			node.network.Send( "OK;"+strconv.Itoa(key)+";"+strconv.Itoa(value)+"\n")
			break
		default:
			break
		}
	}
}




