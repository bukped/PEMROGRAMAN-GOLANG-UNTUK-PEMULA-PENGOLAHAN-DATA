package main

import "fmt"

// Definisikan struct "Rectangle"
type Rectangle struct {
	Width  float64
	Height float64
}

// Metode "CalculateArea" untuk struct "Rectangle"
func (r Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

func main() {
	// Buat instance struct "Rectangle"
	rect := Rectangle{
		Width:  5,
		Height: 3,
	}

	// Panggil metode "CalculateArea"
	area := rect.CalculateArea()

	fmt.Println("Area:", area)
}
