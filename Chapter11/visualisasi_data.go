package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type SalesData struct {
	Date        string
	Product     string
	Category    string
	Quantity    int
	TotalAmount float64
}

func main() {
	// Membaca file sales.csv
	file, err := os.Open("sales.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Menginisialisasi variabel untuk menyimpan data penjualan
	var salesData []SalesData

	// Membaca dan mengubah format data penjualan
	for i, record := range records {
		// Skip baris header
		if i == 0 {
			continue
		}

		quantity, err := strconv.Atoi(record[3])
		if err != nil {
			log.Printf("Error parsing Quantity at row %d: %s\n", i+1, err.Error())
			continue
		}

		totalAmount, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Printf("Error parsing TotalAmount at row %d: %s\n", i+1, err.Error())
			continue
		}

		data := SalesData{
			Date:        record[0],
			Product:     record[1],
			Category:    record[2],
			Quantity:    quantity,
			TotalAmount: totalAmount,
		}

		salesData = append(salesData, data)
	}

	// Menghitung total penjualan per kategori produk
	categorySales := make(map[string]float64)
	for _, data := range salesData {
		categorySales[data.Category] += data.TotalAmount
	}

	// Membuat plot baru
	p := plot.New()
	p.Title.Text = "Total Penjualan per Kategori Produk"
	p.Y.Label.Text = "Total Penjualan"

	// Membuat slice untuk menyimpan nilai dan label kategori
	var values []float64
	var labels []string
	for category, sales := range categorySales {
		values = append(values, sales)
		labels = append(labels, category)
	}

	// Membuat grafik batang
	bars := plotter.Values(values)
	barWidth := vg.Points(50)
	bar, err := plotter.NewBarChart(bars, barWidth)
	if err != nil {
		log.Fatal(err)
	}

	// Menambahkan label ke sumbu X
	p.NominalX(labels...)
	p.Add(bar)

	// Menyimpan plot sebagai file gambar
	err = p.Save(6*vg.Inch, 4*vg.Inch, "grafik_penjualan.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Grafik penjualan berhasil disimpan sebagai 'grafik_penjualan.png'")
}
