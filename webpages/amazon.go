package webpages

import (
	"fmt"
	"ps5/request"
	"strings"
	"time"
)

func Amazon(url string) bool {

	scraper, err := request.New(url)
	if err != nil {
		println(fmt.Sprintf(scraperError + err.Error()))
	}

	entry := scraper.GetSelector("#availability > span")
	inner, _ := entry.InnerText()
	scraper.Stop()

	if strings.Contains(inner, "Currently unavailable") {
		println(fmt.Sprintf(debugStillOutOfStock, "Amazon", time.Now().Format("2006-01-02 15:04:05")))
	} else {
		println(fmt.Sprintf(nowInStock, url))
		return true
	}
	return false
}
