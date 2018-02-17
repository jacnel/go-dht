package dhtnode

const keyRange int = 10000
const dhtPortNum int = 8403

func check(e error) {
	if e != nil {
		panic(e)
	}
}
