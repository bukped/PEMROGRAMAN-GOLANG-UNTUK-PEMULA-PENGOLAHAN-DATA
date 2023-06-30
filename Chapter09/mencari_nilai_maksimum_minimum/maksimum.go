package main

import (
    "fmt"
)

func main() {
    dataset := []int{1000, 1200, 1500, 800, 2000, 900}
    max := dataset[0]
    for _, value := range dataset {
        if value > max {
            max = value
        }
    }
    fmt.Println("Nilai Maksimum:", max)
}
