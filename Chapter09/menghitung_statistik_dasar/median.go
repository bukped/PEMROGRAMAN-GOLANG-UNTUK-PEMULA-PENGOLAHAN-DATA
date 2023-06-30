package main

import (
    "fmt"
    "sort"
)

func main() {
    dataset := []float64{5, 2, 7, 3, 8, 6, 4}
    sort.Float64s(dataset)
    length := len(dataset)
    median := 0.0
    if length%2 == 0 {
        median = (dataset[length/2-1] + dataset[length/2]) / 2
    } else {
        median = dataset[length/2]
    }
    fmt.Println("Median:", median)
}
