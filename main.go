package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gocolly/colly"
)

type Entry struct {
	Name		string
	Category	string
	Submitted	string
}

func main() {
	var entries []Entry
	entries = make([]Entry, 0)

	var counter = 0
	var counterMax = 259

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	// Create another collector to scrape entry details
	// detailCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	
	c.OnHTML("table.views-table tr", func(e *colly.HTMLElement) {
		entry := Entry{
			Name:		e.ChildText("td.views-field-title"),
			Category:	e.ChildText("td.views-field-field-article-primary-category"),
			Submitted:	e.ChildText("td.views-field-created"),
		}
		entries = append(entries, entry)
		counter++
	})


	c.OnScraped(func(r *colly.Response) {
		fmt.Print("finished scraping", r.Request.URL)
		fmt.Println("    Counter:", counter)
	})

	for i := 0; i < counterMax; i++ {
		c.Visit("https://www.programmableweb.com/category/all/mashups?order=created&page=" + strconv.Itoa(i))
	}

	entries = entries[1:]
	
	entriesJson, err := json.Marshal(&entries)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		//fmt.Printf("\nMarshalled data: %s\n", entriesJson)
	}

	ioutil.WriteFile("./output.json", entriesJson, 0777)
	fmt.Println("Data saved")

}