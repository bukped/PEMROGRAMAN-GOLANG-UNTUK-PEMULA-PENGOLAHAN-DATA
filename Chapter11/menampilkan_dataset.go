package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Sale struct {
	Date         string
	Product      string
	Category     string
	Quantity     int
	TotalAmount  float64
}

func main() {
	// Membuka file CSV
	file, err := os.Open("sales.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Membaca data CSV
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Menginisialisasi slice untuk menyimpan data penjualan
	var sales []Sale

	// Memproses setiap baris data
	for i, row := range rows {
		// Melewati header
		if i == 0 {
			continue
		}

		// Membaca nilai dari setiap kolom
		date := row[0]
		product := row[1]
		category := row[2]
		quantity := row[3]
		totalAmount := row[4]

		// Mengkonversi nilai ke tipe yang sesuai
		quantityInt := 0
		fmt.Sscanf(quantity, "%d", &quantityInt)

		totalAmountFloat := 0.0
		fmt.Sscanf(totalAmount, "%f", &totalAmountFloat)

		// Membuat objek Sale
		sale := Sale{
			Date:         date,
			Product:      product,
			Category:     category,
			Quantity:     quantityInt,
			TotalAmount:  totalAmountFloat,
		}

		// Menambahkan data penjualan ke slice
		sales = append(sales, sale)
	}

	// Menampilkan data penjualan
	for _, sale := range sales {
		fmt.Printf("Date: %s, Product: %s, Category: %s, Quantity: %d, TotalAmount: %.2f\n",
			sale.Date, sale.Product, sale.Category, sale.Quantity, sale.TotalAmount)
	}
}
