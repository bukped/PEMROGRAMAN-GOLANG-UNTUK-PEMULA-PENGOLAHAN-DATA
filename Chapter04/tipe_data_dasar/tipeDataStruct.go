package main

import "fmt"

type Mahasiswa struct {
	Nama   string
	Umur   int
	Alamat string
}

func main() {
	var p Mahasiswa

	p.Nama = "John Doe"
	p.Umur = 25
	p.Alamat = "Jl. Raya 123"

	fmt.Println("Nama:", p.Nama)
	fmt.Println("Umur:", p.Umur)
	fmt.Println("Alamat:", p.Alamat)
}
