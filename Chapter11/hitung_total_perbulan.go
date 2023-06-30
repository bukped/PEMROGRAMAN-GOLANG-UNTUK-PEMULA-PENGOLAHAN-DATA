package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Sale struct {
	Date         time.Time
	Product      string
	Category     string
	Quantity     int
	TotalAmount  float64
}

func main() {
	// Membaca dataset penjualan dari file CSV
	salesData, err := readSalesData("sales.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Menghitung total penjualan per bulan
	salesPerMonth := calculateSalesPerMonth(salesData)

	// Menampilkan total penjualan per bulan
	for month, total := range salesPerMonth {
		fmt.Printf("Bulan: %s, Total Penjualan: %.2f\n", month, total)
	}
}

// Membaca dataset penjualan dari file CSV
func readSalesData(filename string) ([]Sale, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var salesData []Sale
	for _, row := range rows[1:] {
		date, err := time.Parse("2006-01-02", row[0])
		if err != nil {
			log.Println("Invalid date format:", row[0])
			continue
		}

		quantity, err := strconv.Atoi(row[3])
		if err != nil {
			log.Println("Invalid quantity:", row[3])
			continue
		}

		totalAmount, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			log.Println("Invalid total amount:", row[4])
			continue
		}

		sale := Sale{
			Date:        date,
			Product:     row[1],
			Category:    row[2],
			Quantity:    quantity,
			TotalAmount: totalAmount,
		}
		salesData = append(salesData, sale)
	}

	return salesData, nil
}

// Menghitung total penjualan per bulan
func calculateSalesPerMonth(salesData []Sale) map[string]float64 {
	salesPerMonth := make(map[string]float64)

	for _, sale := range salesData {
		month := sale.Date.Format("January 2006")
		salesPerMonth[month] += sale.TotalAmount
	}

	return salesPerMonth
}
