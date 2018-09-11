package main

import (
    "fmt" 
)

func main() {
    const max = 1000
    multiplesOf := [...]int{3, 5}
    total := 0

    for i := 3; i < max; i++ {
        for _, n := range multiplesOf {
            if i % n == 0 {
                total += i
                break
            }
        }
    }

    fmt.Println("== Euler project - Problem 1 ==")
    fmt.Println("Σ of ℕ multiples of", multiplesOf, "under", max, "is:", total)
}
