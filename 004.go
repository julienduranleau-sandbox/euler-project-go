package main

import (
    "fmt"
    "strconv"
    "math"
)

func main() {
    const nDigits = 3
    result := 0

    min := math.Pow(10, nDigits - 1)
    max := math.Pow(10, nDigits)

    for i := min; i < max; i++ {
        for j := min; j < max; j++ {
            n := int(i * j)
            nStr := strconv.Itoa(int(i * j))
            if nStr == reverse(nStr) && n > result {
                result = n
            }
        }
    }

    fmt.Println("== Euler project - Problem 4 ==")
    fmt.Println("Largest palindrome from the product of",nDigits,"is",result)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
