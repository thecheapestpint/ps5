package main

import (
	"ps5/request"
	"ps5/webpages"
)

const (
	smythsUrl = "https://www.smythstoys.com/uk/en-gb/video-games-and-tablets/playstation-5/playstation-5-consoles/playstation-5-console/p/191259"
	amazonUrl = "https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452/ref=sr_1_5?dchild=1&keywords=ps5&qid=1607945045&sr=8-5&th=1"
	currysUrl = "https://www.currys.co.uk/gbuk/sony-ps5/console-gaming/consoles/634_4783_32541_49_ba00013671-bv00313579/xx-criteria.html"
	argosUrl  = "https://www.argos.co.uk/product/8349000"
)

func main() {
	println("Starting ps5 notification script")

	inStock := webpages.SmythScrape(smythsUrl)
	if inStock {
		request.OpenBrowser(smythsUrl)
	}

	inStock = webpages.Amazon(amazonUrl)
	if inStock {
		request.OpenBrowser(amazonUrl)
	}

	inStock = webpages.Currys(currysUrl)
	if inStock {
		request.OpenBrowser(smythsUrl)
	}

	inStock = webpages.Argos(argosUrl)
	if inStock {
		request.OpenBrowser(argosUrl)
	}

}
