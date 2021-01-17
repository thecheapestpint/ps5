package webpages

import (
	"fmt"
	"ps5/request"
	"strings"
	"time"
)

func Currys(url string) bool {

	scraper, err := request.New(url)
	if err != nil {
		println(scraperError + err.Error())
		return false
	}

	entry := scraper.GetSelector("#content > div.row > section.col9 > div.row.tsp > div > h4 > p:nth-child(1)")
	text, err := entry.TextContent(); if err != nil {
		println(scraperError + err.Error())
		return false
	}
	scraper.Stop()

	if strings.Contains(text, "No results were found for your search") {
		println(fmt.Sprintf(debugStillOutOfStock, "Currys", time.Now().Format("2006-01-02 15:04:05")))
	} else {
		println(fmt.Sprintf(nowInStock, url))
		return true
	}
	return false
}
