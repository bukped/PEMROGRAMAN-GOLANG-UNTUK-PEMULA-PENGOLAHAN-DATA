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
	data := plotter.Values{2, 3, 5, 6, 4, 8}

	// Membuat plot baru
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	// Menambahkan grafik batang ke plot
	bar, err := plotter.NewBarChart(data, vg.Points(50))
	if err != nil {
		log.Fatal(err)
	}
	p.Add(bar)

	// Menyimpan plot sebagai file gambar
	err = p.Save(4*vg.Inch, 4*vg.Inch, "grafik.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Grafik berhasil disimpan sebagai 'grafik.png'")
}
