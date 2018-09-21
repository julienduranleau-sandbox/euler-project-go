package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	const expo = 1000

	var r big.Int
	r.Exp(big.NewInt(2), big.NewInt(expo), nil)

	str := r.String()
	sum := 0

	for i := 0; i < len(str); i++ {
		n, _ := strconv.Atoi(string(str[i]))
		sum += n
	}

	fmt.Println("== Euler project - Problem 16 ==")
	fmt.Println("Sum of digits from 2^", expo, "is", sum)
}
