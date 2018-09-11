package main

import (
    "fmt"
)

func main() {
    const max = 20
    smallest := max
    
    for {
        isValid := true

        for i := 2; i <= max; i++ {
            if smallest % i != 0 {
                isValid = false
                break
            }
        }

        if isValid {
            break
        }

        smallest++
    }

    fmt.Println("== Euler project - Problem 5 ==")
    fmt.Println("Smallest number evenly divisible by 1 to", max,"is",smallest)
}

