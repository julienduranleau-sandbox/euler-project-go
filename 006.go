package main

import (
    "fmt"
)

func main() {
    const max = 100
    sumOfSq := 0
    sum := 0

    for i := 1; i <= max; i++ {
        sum += i
        sumOfSq += i*i
    }

    result := sum*sum - sumOfSq

    fmt.Println("== Euler project - Problem 6 ==")
    fmt.Println("Diff between sum of squares and square of sum of first",max,"naturals is ",result)
}
