package main

import (
	"fmt"
	"os"
)

type Entry struct {
	Name		string	`json:"Name"`
	Category	string	`json:"Category"`
	Submitted	string	`json:"Submitted"`
}

func main() {
	args := os.Args[1:]
	
	if len(args) == 1 {
		switch args[0] {
			case "scrape":
				action_scrape()
			case "plot":
				action_plot()
		}
	} else {
		os.Exit(1)
	}
}

func action_scrape() {
	url := "https://www.programmableweb.com/category/all/mashups?order=created&page="
	pages := 2
	
	fmt.Println("Scrapign data from https://www.programmableweb.com/category/all/mashups")
	scrapeProgrammableWeb(url, pages)	
	fmt.Println("Finished")
}

func action_plot() {
	file := "./mashups.json"

	fmt.Println("Plotting... ")
	plotEntriesFromJson(file)
	fmt.Println("Finished")
}