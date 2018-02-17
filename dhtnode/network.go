package dhtnode

import (
	"net"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strings"
	"strconv"
	"io"
)

type Network struct {
	myAddress string
	myNodeID  int
	ip2idMap  map[string]int
	id2ipMap  map[int]string
	listener  net.Listener
}

// public functions
func (network *Network) Init(port int, configFilepath string) {
	network.getNetworkConfig(configFilepath)
	network.setMyAddr(port)
}

func (network *Network) Listen() {
	ln, err := net.Listen("tcp", network.myAddress)
	check(err)
	network.listener = ln
}

func (network *Network) Accept() *net.Conn{
	conn, err := network.listener.Accept()
	if err != nil {
		fmt.Println("Failed to handle incoming connection...")
	} else {
		fmt.Println("Connection established...")
	}
	return &conn
}

func (network *Network) Receive(conn *net.Conn) (int, int, int){
	message := getMessage(conn)
	opcode, key, value := parseClientMessage(message)
	return opcode, key, value
}

func (network *Network) KeyInRange(key int) bool {
	targetNode := network.hashKey(key)
	fmt.Println("Target DHTNode -> ", targetNode)
	if targetNode == network.myNodeID {
		return true
	}
	return false
}

func (network *Network) Send(conn *net.Conn, message string) {
	data := []byte(message)
	var(
		err error
		n int
	)
	n, err = (*conn).Write(data)
	fmt.Println(n, err, message)
	check(err)
	fmt.Println("HERE!")
}

func (network *Network) LetsGoOffNoding(opcode, key, value int) (int, int) {
	targetNode := network.hashKey(key)
	targetAddr := network.id2ipMap[targetNode]
	// set up connection with target node and send info
	conn, err := net.Dial("tcp", targetAddr)
	check(err)
	message := strconv.Itoa(opcode)+";"+strconv.Itoa(key)+";"+strconv.Itoa(value)
	_, err = conn.Write([]byte(message))
	check(err)
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if(err == io.EOF) {
		return 0, 0
	}
	check(err)
	conn.Close()
	fmt.Println("HERE")
	return parseNodeMessage(string(data[:n]))
}
func (network *Network) Close(conn *net.Conn) {
	(*conn).Close()
}

// dependent helper functions
func (network *Network) getNetworkConfig(configFilepath string) {
	config := loadConfig(configFilepath)
	network.ip2idMap = make(map[string]int)
	network.id2ipMap = make(map[int]string)
	for nodeID, ipString := range config["network"] {
		network.ip2idMap[ipString] = nodeID
		network.id2ipMap[nodeID] = ipString
	}
}
func (network *Network) setMyAddr(port int) {
	ifaces, err := net.InterfaceAddrs()
	check(err)
	var addrString string
	found := false
	fmt.Println(network.ip2idMap)
	for i := 0; !found && i < len(ifaces); i++ {
		addrString = strings.Split(ifaces[i].String(), "/")[0]
		addrString = addrString + ":" + strconv.Itoa(port)
		fmt.Println(addrString)
		if val, ok := network.ip2idMap[addrString]; ok {
			fmt.Println(val)
			network.myNodeID = val
			network.myAddress = addrString
			found = true
		}
	}
	if !found {
		panic("Oops, this node's IP is not in the network configuration...")
	}
}
func (network *Network) hashKey(key int) int {
	return key % len(network.ip2idMap)
}

// independent helper functions
func loadConfig(config_file string) map[string][]string {
	data, err := ioutil.ReadFile(config_file)
	check(err)
	var json_obj map[string][]string
	err = json.Unmarshal(data, &json_obj)
	check(err)
	return json_obj
}
func getMessage(conn *net.Conn) string {
	data := make([]byte, 1024)
	var(
		err error
		n int
	)
	n, err = (*conn).Read(data)
	check(err)
	if(err == io.EOF) {
		return "-1;0;0"
	}
	return string(data[:n])
}
func parseClientMessage(s string) (int, int, int) {
	tokens := strings.Split(s, ";")
	vals := make([]int, 3)
	if len(tokens) < 3 {
		return 0, 0, 0
	}
	for i, s := range tokens {
		if i >= 3 {
			break
		}
		tokens[i] = strings.TrimSpace(s)
		vals[i],_ = strconv.Atoi(tokens[i])
	}
	opcode, key, value := vals[0], vals[1], vals[2]
	return opcode, key, value
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


