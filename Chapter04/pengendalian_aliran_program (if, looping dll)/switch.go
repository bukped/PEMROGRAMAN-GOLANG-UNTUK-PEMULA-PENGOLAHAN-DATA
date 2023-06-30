package main

import "fmt"

func main() {
    hari := "Senin"
    switch hari {
    case "Senin":
        fmt.Println("Hari pertama dalam seminggu")
    case "Selasa":
        fmt.Println("Hari kedua dalam seminggu")
    default:
        fmt.Println("Hari yang lain")
    }
}
