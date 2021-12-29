package main

import (
	"fmt"
	"os"
)

type Entry struct {
	Name		string
	Category	string
	Submitted	string
}

func main() {
	argsWithProg := os.Args[1:]
	fmt.Println(len(argsWithProg))
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
		
	os.Exit(99)
	
}

func action_scrape() {
	url := "https://www.programmableweb.com/category/all/mashups?order=created&page="
	pages := 2
	
	scrapeProgrammableWeb(url, pages)	
}

func action_plot() {
	
}