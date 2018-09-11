package main

import (
    "fmt"
)

func main() {
    const target = 1000
    answer := 0

    for a := 0; a < target; a++ {
        for b := a; b < target; b++ {
            for c := b; c < target; c++ {
                if a*a + b*b == c*c {
                    // is pythagorean triplet
                    if a+b+c == target {
                        answer = a*b*c
                    }
                }
            }
        }
    }

    fmt.Println("== Euler project - Problem 9 ==")
    fmt.Println(answer)
}
