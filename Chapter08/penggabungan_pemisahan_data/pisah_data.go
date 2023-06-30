package main

import "fmt"

func main() {
    data := []int{1, 2, 3, 4, 5, 6}

    // Pemisahan data menjadi dua subset
    splitIndex := 3
    subset1 := data[:splitIndex]
    subset2 := data[splitIndex:]

    fmt.Println("Subset 1:", subset1)
    fmt.Println("Subset 2:", subset2)
}
