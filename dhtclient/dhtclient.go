package dhtclient

import (
	"net"
	"strconv"
	"strings"
	"io"
)

type DHTClient struct {
	targetAddr string
	dhtConn    *net.Conn
}

func (client *DHTClient) Init(addr string) {
	client.targetAddr = addr
	conn, err := net.Dial("tcp", client.targetAddr)
	client.dhtConn = &conn
	if(err != nil) {
		panic("Could not connect to " + client.targetAddr)
	}
}

func (client *DHTClient) Get(key int) (int, int) {
	_, err := (*client.dhtConn).Write([]byte("1;"+strconv.Itoa(key)+";\n"))
	if err != nil {
		panic(err)
	}
	data := make([]byte, 1024)
	n, err := (*client.dhtConn).Read(data)
	if err != nil {
		if err == io.EOF {
			return 0, 0
		}
	}
	return parseNodeMessage(string(data[:n]))
}

func (client *DHTClient) Put(key, value int) (int, int) {
	n, err := (*client.dhtConn).Write([]byte("2;"+strconv.Itoa(key)+";"+strconv.Itoa(value)+"\n"))
	if err != nil {
		panic(err)
	}
	data := make([]byte, 1024)
	n, err = (*client.dhtConn).Read(data)
	if err != nil {
		if err == io.EOF {
			return 0, 0
		}
	}
	return parseNodeMessage(string(data[:n]))
}

func (client *DHTClient) Close() {
	(*client.dhtConn).Close()
}

func parseNodeMessage(s string) (int, int) {
	tokens := strings.Split(s, ";")
	vals := make([]int, 3)
	if len(tokens) < 3 {
		return 0, 0
	}
	for i, tok := range tokens {
		if i >= 3 {
			break
		}
		tokens[i] = strings.TrimSpace(tok)
		if strings.Compare(tok, "OK") == 0 {
			vals[i] = 2
		} else if strings.Compare(tok, "EXISTS") == 0 {
			vals[i] = 1
		} else if strings.Compare(tok, "FAIL") == 0 {
			vals[i] = 0
		} else {
			vals[i],_ = strconv.Atoi(tokens[i])
		}
	}
	ok := vals[0]
	value := vals[2]
	return value, ok
}