package main

import (
    "fmt"
)

func main() {
    fmt.Println("This is going to take a while...")

    triangleNumber := 1
    foundSum := 0
    highestDivisions := 0

    for {
        sum := 0
        nDivisions := 0

        for i := 1; i <= triangleNumber; i++ {
            sum += i
        }

        // todo: Use prime factorization method instead of brute force
        for i := 1; i <= sum; i++ {
            if sum % i == 0 {
                nDivisions++
            }
        }

        if nDivisions > 500 {
            foundSum = sum
            break
        } else {
            if nDivisions > highestDivisions {
                highestDivisions = nDivisions
                fmt.Println(highestDivisions,"divisions for",sum)
            }
        }

        triangleNumber++
    }

    fmt.Println("== Euler project - Problem 12 ==")
    fmt.Println(foundSum)
}
