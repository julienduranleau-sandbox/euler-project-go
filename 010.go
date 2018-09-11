package main

import (
    "fmt"
)

func main() {
    const target = 2000000
    
    var primes []int = []int{2,3}
    sum := 5
    
    for k := 1; primes[len(primes)-1] < target; k++ {
        targetReached := false

        for _, n := range [2]int{6*k-1, 6*k+1} {
            isPrime := true
            
            for _, p := range primes {
                if n % p == 0 {
                    isPrime = false
                    break
                }
            }

            if isPrime {
                if (n < target) {
                    primes = append(primes, n)
                    sum += n
                } else {
                    targetReached = true
                    break
                }
            }
        }

        if targetReached {
            break
        }
    }

    fmt.Println("== Euler project - Problem 10 ==")
    fmt.Println(sum)
}
