package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type PlotData struct {
	Year	float64
	Amount	float64
}

func plotEntriesFromJson(file string) {
	entries := loadEntriesFromJson(file)
	data := processEntries(entries)
	mashup_points := plotData(data)
	fmt.Println("Points[{X, Y}]: ", mashup_points)


	// init plot
	p := plot.New()
	p.Title.Text = "Anzahl neuer Mashups pro Jahr"
	p.X.Label.Text = "Jahr"
	p.Y.Label.Text = "Anzahl"
	p.Add(plotter.NewGrid())
	
	// scatter points
	scatter, err := plotter.NewScatter(mashup_points)
	if err != nil {
		panic(err)
	}
	scatter.Shape = draw.CircleGlyph{}
	scatter.Radius = vg.Points(3)
	
	// line graph
	line, err := plotter.NewLine(mashup_points)
	if err != nil {
		panic(err)
	}
	line.LineStyle.Width = vg.Points(0.5)
	line.LineStyle.Dashes = []vg.Length{vg.Points(3), vg.Points(3)}
	//line.LineStyle.Color = color.RGBA{B: 255, A: 255}
	
	// add points and graphs to plot
	p.Add(scatter, line)

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "./plot.png"); 
	err != nil {
		panic(err)
	}
}

func loadEntriesFromJson(file string) []Entry {
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var entries []Entry
	json.Unmarshal(byteValue, &entries)

	return entries
}

func processEntries(entries []Entry) []PlotData {
	var data []PlotData
	data = make([]PlotData, 0)

	counter := 0
	var year string
	year_old := ""
	for _, e := range entries {

		date := strings.Split(e.Submitted, ".")

		if (len(date) == 3) {	
			year = strings.Split(e.Submitted, ".")[2]
		} else {
			continue;
		}

		if (year_old != year) {
			plot_year, err := strconv.ParseFloat(year_old, 64)
			if err != nil {
				plot_year = 0
			}

			data = append(data, PlotData{
				Year: plot_year,
				Amount: float64(counter),
			})

			counter = 0
			year_old = year
		}

		counter++
	}

	plot_year, err := strconv.ParseFloat(year_old, 64)
	if err != nil {
		plot_year = 0
	}

	data = append(data, PlotData{
		Year: plot_year,
		Amount: float64(counter),
	})

	return data[1:]
}

func plotData(data []PlotData) plotter.XYs {
	pts := make(plotter.XYs, len(data))

	for i, e := range data {
		pts[i].X = e.Year
		pts[i].Y = e.Amount
	}

	return pts
}