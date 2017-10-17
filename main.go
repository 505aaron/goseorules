package main

import (
	"fmt"
	//	"github.com/505aaron/goseorules/hello"
	"github.com/asciimoo/colly"
	"strings"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	c.AllowedDomains = []string{"www.cars.com"}

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML("head title", func(e *colly.HTMLElement) {
		text := e.Text

		if isCars := strings.Contains(text, "Cars"); isCars {
			fmt.Printf("Cars found: %q\n", e.Text)
		}
		// Print link

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://cars.com
	c.Visit("https://www.cars.com/")
}
