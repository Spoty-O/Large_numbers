package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	bit_arr := [...]int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
	for i := 0; i < 10; i++ {
		fmt.Println("-----------------------------------------------------------------------------")
		brute_force(random_key(possible_keys(bit_arr[i])))
	}
}

func brute_force(key *big.Int) {
	possible_key := big.NewInt(0)
	start_time := time.Now()
	for key.Cmp(possible_key) != 0 {
		possible_key.Add(possible_key, big.NewInt(1))
	}
	fmt.Printf("Founded key: 0x%x; time: %s\n", possible_key, time.Now().Sub(start_time))
}

func random_key(max *big.Int) *big.Int {
	n, err := rand.Int(rand.Reader, max.Sub(max, big.NewInt(1)))
	if err != nil {
		println(err)
	}
	fmt.Printf("Random key: 0х%x\n", n)
	return n
}

func possible_keys(bit int64) *big.Int {
	var bigBit = big.NewInt(bit)
	bigBit.Exp(big.NewInt(2), bigBit, nil)
	fmt.Printf("Keys for %d-bit: %d\n", bit, bigBit)
	return bigBit
}
