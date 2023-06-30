package main

import "fmt"

// Definisikan struct "Person"
type Person struct {
	Name string
	Age  int
}

func main() {
	// Buat instance struct "Person"
	p := Person{
		Name: "John Doe",
		Age:  30,
	}

	// Mengakses properti struct
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
