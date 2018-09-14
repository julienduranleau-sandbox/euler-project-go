package main

import (
    "fmt"
)

func main() {
    fmt.Println("This is going to take a while...")

    const size = 20
    nPaths := countPathFor(0,0,size) + 1

    fmt.Println("== Euler project - Problem 14 ==")
    fmt.Println(nPaths)
}

func countPathFor(x int, y int, size int) int {
    count := 1

    if x < size - 1 {
        count += countPathFor(x + 1, y, size)
    }

    if y < size - 1 {
        count += countPathFor(x, y + 1, size)
    }

    return count
}
