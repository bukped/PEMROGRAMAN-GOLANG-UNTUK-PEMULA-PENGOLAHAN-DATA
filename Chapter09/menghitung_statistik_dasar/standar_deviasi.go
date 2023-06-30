package main

import (
    "fmt"
    "math"
)

func main() {
    dataset := []float64{5, 2, 7, 3, 8, 6, 4}
    mean := 0.0
    sum := 0.0
    for _, value := range dataset {
        sum += value
    }
    mean = sum / float64(len(dataset))
    variance := 0.0
    for _, value := range dataset {
        variance += math.Pow(value-mean, 2)
    }
    variance = variance / float64(len(dataset))
    stdDev := math.Sqrt(variance)
    fmt.Println("Deviasi Standar:", stdDev)
}
