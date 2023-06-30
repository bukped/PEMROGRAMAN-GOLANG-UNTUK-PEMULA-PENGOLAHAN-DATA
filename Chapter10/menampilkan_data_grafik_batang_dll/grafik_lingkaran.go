package main

import (
	"fmt"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
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

	// Membuat diagram lingkaran menggunakan pustaka plotter
	pie, err := plotter.NewPie(values)
	if err != nil {
		log.Fatal(err)
	}
	pie.Radius = 0.5

	// Menambahkan label ke diagram lingkaran
	pie.Labels = labels

	// Menambahkan diagram lingkaran ke plot
	p.Add(pie)

	// Membuat gambar plot
	img := vgimg.New(4*vg.Inch, 4*vg.Inch)
	dc := vgimg.NewDrawer(img)
	p.Draw(dc)

	// Menyimpan gambar plot sebagai file
	err = vgimg.PngCanvas("diagram_lingkaran.png", img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Diagram lingkaran berhasil disimpan sebagai 'diagram_lingkaran.png'")
}
