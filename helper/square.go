package helper

import (
	"fmt"
)

func Square() {
    fmt.Println("Bilangan Kuadrat:")
    for i := 1; i <= 10; i++ {
        square := i * i
        fmt.Printf("%d^2 = %d\n", i, square)
    }
}
