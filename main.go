package main

import (
	"fmt"
)


func main() {
    fmt.Println("Bilangan Prima:")
    for i := 2; i <= 100; i++ {
        if IsPrime(i) {
            fmt.Println(i)
        }
    }
}
