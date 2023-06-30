package main

import "fmt"

// Fungsi tambah untuk menjumlahkan dua bilangan
func tambah(a, b int) int {
	return a + b
}

// Fungsi hitungLuasPersegi untuk menghitung luas persegi
func hitungLuasPersegi(sisi int) int {
	return sisi * sisi
}

func main() {
	// Pemanggilan fungsi tambah
	hasil := tambah(3, 4)
	fmt.Println("Hasil penjumlahan:", hasil)

	// Pemanggilan fungsi hitungLuasPersegi
	luas := hitungLuasPersegi(5)
	fmt.Println("Luas persegi:", luas)
}
