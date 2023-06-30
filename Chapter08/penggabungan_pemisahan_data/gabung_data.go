package main

import "fmt"

func main() {
    data1 := []int{1, 2, 3}
    data2 := []int{4, 5, 6}
    // Penggabungan data1 dan data2
    combinedData := append(data1, data2...)
    fmt.Println("Data setelah penggabungan:", combinedData)
}
