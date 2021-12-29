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
		
		switch args[1] {
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

	os.Exit(22)
	
	url := "https://www.programmableweb.com/category/all/mashups?order=created&page="
	page := 2
	
	scrape(url, pages)
	
}

func action_plot() {
	
}