package dhtnode

import (
	"fmt"
	"net"
	"strconv"
)

type DHTNode struct {
	hashTable LocalStore
	network   Network
}

func (node DHTNode) StartNode(port int) {
	node.init(port)
	node.Listen()
}

func (node *DHTNode) init(port int) {
	node.hashTable = LocalStore{}
	node.hashTable.Init()
	node.network = Network{}
	node.network.Init(port)
}

func (node *DHTNode) Listen() {
	node.network.Listen()
	for{
		conn := node.network.Accept()
		go node.handleMessages(conn)
	}
}

func (node *DHTNode) handleMessages(conn *net.Conn) {
	var opcode, key, value int
	for {
		opcode, key, value = node.network.Receive(conn)
		ok := -1
		switch opcode {
		case -1:
			fmt.Println("Closing connection...")
			node.network.Close(conn)
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
			node.network.Send(conn, node.hashTable.String())
			fmt.Println("Current state:")
			fmt.Println(node.hashTable.String())
			break
		case 4:
			node.hashTable.Clear()
			node.network.Send(conn, "OK;;")
			fmt.Println("TABLE CLEARED!")
			break
		default:
			node.network.Send(conn, "IGNORED\n")
			fmt.Println("Ignoring message")
		}

		switch ok {
		case 0:
			node.network.Send(conn, "FAIL;"+strconv.Itoa(key)+";"+strconv.Itoa(value)+"\n")
			break
		case 1:
			node.network.Send(conn, "EXISTS;"+strconv.Itoa(key)+";"+strconv.Itoa(value)+"\n")
			break
		case 2:
			node.network.Send(conn, "OK;"+strconv.Itoa(key)+";"+strconv.Itoa(value)+"\n")
			break
		default:
			break
		}
	}
}




