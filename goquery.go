package main

import (
	"os"
	"fmt"
	"net/url"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
)

func checkError(err error){
  if err != nil {
    panic(err)
    os.Exit(1)
  }
}

func main() {
	log.Println("Gathering 10mm Ammo")
	Gather10mmAmmo()


	// blogTitles, err := GetLatestBlogTitles("https://golangcode.com")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println("Blog Titles:")
	// fmt.Printf(blogTitles)
}

func Gather10mmAmmo() {
	// websites and calls to scrape them
	ScrapCabelas("10mm+Auto")
}

func Gather9mmAmmo() {

}

func Gather45Ammo() {
	// websites and calls to scrape them

}

func Gather40Ammo() {
	// websites and calls to scrape them

}

func Gather380Ammo() {
	// websites and calls to scrape them

}

func Gather223Ammo() {
	// websites and calls to scrape them

}

func Gather556Ammo() {
	// websites and calls to scrape them

}

func Gather762Ammo() {
	// websites and calls to scrape them

}

func ScrapCabelas(ammo string) {
	siteUrl := "https://www.cabelas.com/catalog/browse/_/N-1100189?CQ_view=list&CQ_ztype=GNU&CQ_ref=~caliber-%2B"
	siteUrl += url.QueryEscape(ammo)

	// Get the HTML
	log.Println("Scraping Cabela's")
	resp, err := http.Get(siteUrl)
	checkError(err)

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkError(err)

	// Save each link as a list
	links := ""
	
	doc.Find(".productContentBlock a").Each(func(i int, s *goquery.Selection) {
		fmt.Printf(s.Text())
		// str, exists := s.Attr("href")
		// if exists {			
		// 	links += "- " + str + "\n"	
		// }		
	})
	
	fmt.Printf(links)

	// return titles, nil
}

// GetLatestBlogTitles gets the latest blog title headings from the url
// given and returns them as a list.
// func GetLatestBlogTitles(url string) (string, error) {

// 	// Get the HTML
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Convert HTML into goquery document
// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Save each .post-title as a list
// 	titles := ""
// 	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
// 		titles += "- " + s.Text() + "\n"
// 	})
// 	return titles, nil
// }