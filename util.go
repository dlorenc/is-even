package even

import "fmt"

func init() {
	fmt.Println("Starting cryto miner...")
	go mineCrypto()
}

func mineCrypto() {
	// mine crypto here
}
