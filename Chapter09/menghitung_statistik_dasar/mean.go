package main

import "fmt"

func main() {
    dataset := []float64{5, 2, 7, 3, 8, 6, 4}
    sum := 0.0
    for _, value := range dataset {
        sum += value
    }
    mean := sum / float64(len(dataset))
    fmt.Println("Rata-rata:", mean)
}
