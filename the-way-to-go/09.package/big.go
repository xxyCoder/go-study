package main

import (
	"fmt"
	"math"
	"math/big"
)

func bigp() {
	im := big.NewInt(math.MaxInt64)
	in := im
	ip := big.NewInt(1)
	ip.Mul(im, in)
	fmt.Println(ip)
}
