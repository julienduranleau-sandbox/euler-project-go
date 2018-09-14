package main

import (
    "fmt"
)

func main() {

    const max = 1000000
    resultNSteps := 0
    result := 0

    for i := 2; i < max; i++ {
        n := i
        nSteps := 0

        for n != 1 {
            nSteps++

            if n % 2 == 0 {
                n = n / 2
            } else {
                n = n * 3 + 1
            }
        }

        if nSteps > resultNSteps {
            resultNSteps = nSteps
            result = i
        }
    }

    fmt.Println("== Euler project - Problem 14 ==")
    fmt.Println("Number with highest n steps for Collatz seq <",max,"is",result,"with",resultNSteps,"steps")
}
