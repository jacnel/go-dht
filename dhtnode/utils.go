package dhtnode

const keyRange int = 10000

func check(e error) {
	if e != nil {
		panic(e)
	}
}
