package main

import (
    "fmt" 
)

func main() {
    const max = 600851475143

    currentMax := max
    factor := 2
    
    for currentMax > factor {
        if currentMax % factor == 0 {
            currentMax /= factor
            factor = 2
        } else {
            factor++
        }
    }
    

    fmt.Println("== Euler project - Problem 3 ==")
    fmt.Println("Largest prime factor of",max,"is",factor)
}
