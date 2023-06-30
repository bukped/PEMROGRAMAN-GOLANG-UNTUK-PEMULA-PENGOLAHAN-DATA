package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// Membaca dataset penjualan dari file CSV
	file, err := os.Open("sales.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Membuat map untuk menyimpan total penjualan per kategori produk
	categorySales := make(map[string]float64)

	// Menghitung total penjualan per kategori produk
	for _, row := range data[1:] {
		category := row[2]
		amount := parseAmount(row[4])

		categorySales[category] += amount
	}

	// Menampilkan hasil analisis
	fmt.Println("Analisis Penjualan Berdasarkan Kategori Produk:")
	for category, sales := range categorySales {
		fmt.Printf("Kategori: %s, Total Penjualan: %.2f\n", category, sales)
	}
}

// Fungsi untuk mengubah string menjadi float64
func parseAmount(amount string) float64 {
	var value float64
	_, err := fmt.Sscanf(amount, "%f", &value)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
