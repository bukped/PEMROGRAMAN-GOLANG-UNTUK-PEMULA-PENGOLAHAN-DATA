package main

import "fmt"

func main() {
    data := []int{10, 20, 30, 40, 50}

  // Pemfilteran data dengan kondisi nilai lebih besar dari 30
    filteredData := make([]int, 0)
    for _, value := range data {
        if value > 30 {
            filteredData = append(filteredData, value)
        }
    }

    fmt.Println("Data setelah pemfilteran:", filteredData)
}
