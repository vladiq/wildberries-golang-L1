package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewFloat(1e20)
	b := big.NewFloat(3e30)

	mulOutput := fmt.Sprintf("%v * %v = %v", a, b, big.NewFloat(0).Mul(a, b))
	divOutput := fmt.Sprintf("%v / %v = %v", a, b, big.NewFloat(0).Quo(a, b))
	addOutput := fmt.Sprintf("%v + %v = %v", a, b, big.NewFloat(0).Add(a, b))
	subOutput := fmt.Sprintf("%v - %v = %v", a, b, big.NewFloat(0).Sub(a, b))

	fmt.Println(mulOutput)
	fmt.Println(divOutput)
	fmt.Println(addOutput)
	fmt.Println(subOutput)
}
