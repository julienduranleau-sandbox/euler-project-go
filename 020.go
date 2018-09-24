package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	N := 100
	product := new(big.Int)
	product.SetString("1", 10)

	for i := N; i > 1; i-- {
		bigI := big.NewInt(int64(i))
		product.Mul(product, bigI)
	}

	digits := product.String()
	sum := 0

	for i := 0; i < len(digits); i++ {
		n, _ := strconv.Atoi(string(digits[i]))
		sum += n
	}

	fmt.Println("== Euler project - Problem 19 ==")
	fmt.Println("Sum of digits of factorial", N, "is", sum)
}
