package main

import (
    "fmt"
)

func main() {
    dataset := []int{1000, 1200, 1500, 800, 2000, 900}
    min := dataset[0]
    for _, value := range dataset {
        if value < min {
            min = value
        }
    }
    fmt.Println("Nilai Minimum:", min)
}
