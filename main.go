package main

import (
	"fmt"
)

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println("Bilangan Prima:")
    for i := 2; i <= 100; i++ {
        if isPrime(i) {
            fmt.Println(i)
        }
    }
}
