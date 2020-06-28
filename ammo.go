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
	Caliber			string
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
	// should walk through each listing on the left bars so it can be done ammo type specific
	// go through each and grab general info from title if possible
	// grab items from short desc if possible
	// grab items from description if possible
	// calculate CPR

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

		roundRegex, _ := regexp.Compile(`(?i)(\d+) rounds per box`)
		count := roundRegex.FindStringSubmatch(details)
	
		log.Println(count[1])

		velocityRegex, _ := regexp.Compile(`(?i)(\d+ fps)`)
		velocity := velocityRegex.FindStringSubmatch(details)

		log.Println(velocity[1])

		ammo := Ammo{
			Title:			title,
			Price:			price,
			Count:			count[0],
			Velocity:		velocity[0],
			Weight:			weight,
			URL:			url,
			CPR:			"",
		}

		ammoList = append(ammoList, ammo)
	})

	// start scraping
	c.Visit("https://www.velocityammosales.com/collections/9-mm")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(ammoList)
}