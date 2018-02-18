package dhtnode

const keyRange int = 100000

func check(e error) {
	if e != nil {
		panic(e)
	}
}
