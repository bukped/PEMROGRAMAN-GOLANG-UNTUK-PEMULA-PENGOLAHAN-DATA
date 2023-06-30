package main

import (
    "fmt"
)

func main() {
    dataset := []int{75, 80, 90, 80, 85, 75, 90}
    frequency := make(map[int]int)
    for _, value := range dataset {
        frequency[value]++
    }
    fmt.Println("Frekuensi Data:")
    for value, count := range frequency {
        fmt.Printf("%d: %d\n", value, count)
    }
}
