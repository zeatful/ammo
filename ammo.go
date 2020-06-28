package main

import (
	"log"
	"encoding/json"	
	"os"
	"regexp"

	"github.com/gocolly/colly"
)

// store ammo information
type Ammo struct {
	Title			string
	Price			string
	Count			string
	Velocity		string
	URL				string
	CPR				string
}

func main() {
	velocity()
}

func velocity() {
	// Instantiate default collector
	c := colly.NewCollector(	
		colly.AllowedDomains("www.velocityammosales.com", "velocityammosales.com"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./velocity_cache"),
	)

	ammoList := make([]Ammo, 0, 200)

	// On every link with the product-grid-item class
	c.OnHTML(`a.product-grid-item`, func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		// check for soldout badge
		soldout := e.ChildText("div.product-grid-image > div.product-grid-image--centered > div.badge > span.badge-label")

		// only use links that aren't soldout
		if(!(len(soldout) > 0)){
			log.Println(link)
			// start scaping the page under the link found
			e.Request.Visit(link)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	// Extract details of the ammo entry
	c.OnHTML(`div.prd-detils`, func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		title := e.ChildText("h1.prd-head")
		price := e.ChildText("span.product-price")
		details := e.ChildText("div.tab.active")

		roundRegex, _ := regexp.Compile(`(\d+) rounds per box`)
		velocityRegex, _ := regexp.Compile(`(\d+ fps)`)
		
		//count := roundRegex.FindStringSubmatch(details)
		//velocity := velocityRegex.FindStringSubmatch(details)

		ammo := Ammo{
			Title:			title,
			Price:			price,
			Count:			"",
			Velocity:		"",
			URL:			url,
			CPR:			"",
		}

		ammoList = append(ammoList, ammo)
	})

	// 	course := Course{
	// 		Title:       title,
	// 		URL:         e.Request.URL.String(),
	// 		Description: e.ChildText("div.content"),
	// 		Creator:     e.ChildText("div.creator-names > span"),
	// 	}
	// 	// Iterate over rows of the table which contains different information
	// 	// about the course
	// 	e.ForEach("table.basic-info-table tr", func(_ int, el *colly.HTMLElement) {
	// 		switch el.ChildText("td:first-child") {
	// 		case "Language":
	// 			course.Language = el.ChildText("td:nth-child(2)")
	// 		case "Level":
	// 			course.Level = el.ChildText("td:nth-child(2)")
	// 		case "Commitment":
	// 			course.Commitment = el.ChildText("td:nth-child(2)")
	// 		case "How To Pass":
	// 			course.HowToPass = el.ChildText("td:nth-child(2)")
	// 		case "User Ratings":
	// 			course.Rating = el.ChildText("td:nth-child(2) div:nth-of-type(2)")
	// 		}
	// 	})
	// 	courses = append(courses, course)
	// })

	// start scraping
	c.Visit("https://www.velocityammosales.com/collections/9-mm")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(ammoList)
}