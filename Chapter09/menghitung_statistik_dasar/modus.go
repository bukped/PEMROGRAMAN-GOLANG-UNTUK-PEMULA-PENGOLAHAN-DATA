package main

import (
    "fmt"
)

func main() {
    dataset := []int{5, 2, 7, 3, 8, 6, 4, 5, 7}
    frequency := make(map[int]int)
    for _, value := range dataset {
        frequency[value]++
    }
    maxFrequency := 0
    modes := []int{}
    for value, count := range frequency {
        if count > maxFrequency {
            maxFrequency = count
            modes = []int{value}
        } else if count == maxFrequency {
            modes = append(modes, value)
        }
    }
    fmt.Println("Modus:", modes)
}
