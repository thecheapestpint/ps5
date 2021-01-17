package webpages

import (
	"fmt"
	"ps5/request"
	"strings"
)

func Currys(url string) bool {

	scraper, err := request.New(url)
	if err != nil {
		println(scraperError + err.Error())
	}

	entry := scraper.GetSelector("#content > div.row > section.col9 > div.row.tsp > div > h4 > p:nth-child(1)")
	text, _ := entry.TextContent()
	scraper.Stop()

	if strings.Contains(text, "No results were found for your search") {
		fmt.Printf(debugStillOutOfStock, "Currys")
	} else {
		fmt.Printf(nowInStock, url)
		return true
	}
	return false
}
