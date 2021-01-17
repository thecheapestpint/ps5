package webpages

import (
	"fmt"
	"ps5/request"
	"strings"
)

func Amazon(url string) bool {

	scraper, err := request.New(url)
	if err != nil {
		println(scraperError + err.Error())
	}

	entry := scraper.GetSelector("#availability > span")
	inner, _ := entry.InnerText()
	scraper.Stop()

	if strings.Contains(inner, "Currently unavailable") {
		fmt.Printf(debugStillOutOfStock, "Amazon")
	} else {
		fmt.Printf(nowInStock, url)
		return true
	}
	return false
}
