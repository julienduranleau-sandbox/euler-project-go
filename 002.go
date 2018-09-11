package main

import (
    "fmt" 
)

func main() {
    const max = 4000000
    sum := 2
    a := 1
    b := 2

    for b <= max {
        c := a + b
        a = b
        b = c

        if c % 2 == 0 {
            sum += c
        }
    }

    fmt.Println("== Euler project - Problem 2 ==")
    fmt.Println("Σ of even fib seq numbers < ",max,"is",sum)
}
