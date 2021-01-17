package webpages

import (
	"fmt"
	"ps5/request"
	"strings"
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
		fmt.Printf(debugStillOutOfStock, "Smyths")
	} else {
		return true
	}

	return false
}
