package main

import (
    "fmt"
)

func main() {
    triangleNumber := 1
    foundSum := 0

    for {
        sum := 0
        nDivisions := 0

        for i := 1; i <= triangleNumber; i++ {
            sum += i
        }

        for i := 1; i <= sum; i++ {
            if sum % i == 0 {
                nDivisions++
            }
        }

        if nDivisions > 500 {
            foundSum = sum
            break
        }

        triangleNumber++
    }

    fmt.Println("== Euler project - Problem 12 ==")
    fmt.Println(foundSum)
}
