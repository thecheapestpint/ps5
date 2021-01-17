package webpages

import (
	"fmt"
	"ps5/request"
	"strings"
	"time"
)

const (
	outOfStockLabel = "out of stock. expected in stock: february 2021"
)

func SmythScrape(url string) bool {

	body, _, err := request.RequestPage(url)
	if err != nil {
		println(fmt.Sprintf(fetchPageError, "smyths ") + err.Error())
	}

	if strings.Contains(strings.ToLower(body), outOfStockLabel) {
		println(fmt.Sprintf(debugStillOutOfStock, "Smyths", time.Now().Format("2006-01-02 15:04:05")))
	} else {
		return true
	}

	return false
}
