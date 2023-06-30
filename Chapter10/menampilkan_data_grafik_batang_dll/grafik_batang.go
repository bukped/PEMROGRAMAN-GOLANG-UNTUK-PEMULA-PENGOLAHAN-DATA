package main

import (
	"fmt"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	var err error
	// Data
	data := []struct {
		Label string
		Value float64
	}{
		{"A", 40},
		{"B", 25},
		{"C", 35},
		{"D", 50},
	}

	// Membuat plot baru
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	// Membuat slice untuk menyimpan nilai dan label
	var values []float64
	var labels []string
	for _, d := range data {
		values = append(values, d.Value)
		labels = append(labels, d.Label)
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
	err = p.Save(4*vg.Inch, 4*vg.Inch, "grafik_batang.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Grafik batang berhasil disimpan sebagai 'grafik_batang.png'")
}
