package main

import (
    "fmt"
)

func main() {
    const nthPrime = 10001 
    var primes []int
    
    i := 2

    for len(primes) < nthPrime {
        isPrime := true

        for _, p := range primes {
            if i % p == 0 {
                isPrime = false
                break
            }
        }

        if isPrime {
            primes = append(primes, i)
        }

        i++
    }

    fmt.Println("== Euler project - Problem 7 ==")
    fmt.Println("prime #",nthPrime,"is", primes[nthPrime - 1])
}
