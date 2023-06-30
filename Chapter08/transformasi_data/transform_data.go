package main

import "fmt"

func main() {
    data := []int{1, 2, 3, 4, 5}

    // Mengalikan setiap elemen data dengan faktor 2
    transformedData := make([]int, len(data))
    for i, value := range data {
        transformedData[i] = value * 2
    }

    fmt.Println("Data setelah transformasi:", transformedData)
}
