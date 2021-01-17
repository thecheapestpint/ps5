package main

import (
	"ps5/request"
	"ps5/webpages"
	"time"
)

const (
	smythsUrl = "https://www.smythstoys.com/uk/en-gb/video-games-and-tablets/playstation-5/playstation-5-consoles/playstation-5-console/p/191259"
	amazonUrl = "https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452/ref=sr_1_5?dchild=1&keywords=ps5&qid=1607945045&sr=8-5&th=1"
	currysUrl = "https://www.currys.co.uk/gbuk/sony-ps5/console-gaming/consoles/634_4783_32541_49_ba00013671-bv00313579/xx-criteria.html"
	argosUrl  = "https://www.argos.co.uk/product/8349000"
	gameUrl   = "https://www.game.co.uk/en/playstation-5-additional-dualsense-controller-marvels-spider-man-miles-morales-2846800"
)

var argosBrowserOpen = false
var currysBrowserOpen = false
var gameBrowserOpen = false
var smythsBrowserOpen = false
var amazonBrowserOpen = false

func main() {
	println("Starting ps5 notification script")

	for _ = range time.Tick(5 * time.Second) {
		scrape()
	}

}

func scrape() {

	go func() {
		inStock := webpages.SmythScrape(smythsUrl)
		if inStock {
			if smythsBrowserOpen {
				println(webpages.BrowserAlreadyOpen)
			} else {
				smythsBrowserOpen = true
				request.OpenBrowser(smythsUrl)
			}
		}
	}()

	go func() {
		inStock := webpages.Amazon(amazonUrl)
		if inStock {
			if amazonBrowserOpen {
				println(webpages.BrowserAlreadyOpen)
			} else {
				amazonBrowserOpen = true
				request.OpenBrowser(amazonUrl)
			}
		}
	}()

	go func() {
		inStock := webpages.Currys(currysUrl)
		if inStock {
			if currysBrowserOpen {
				println(webpages.BrowserAlreadyOpen)
			} else {
				currysBrowserOpen = true
				request.OpenBrowser(currysUrl)
			}
		}
	}()

	go func() {
		inStock := webpages.Argos(argosUrl)
		if inStock {
			if argosBrowserOpen {
				println(webpages.BrowserAlreadyOpen)
			} else {
				argosBrowserOpen = true
				request.OpenBrowser(argosUrl)
			}
		}
	}()

	go func() {
		inStock := webpages.Game(gameUrl)
		if inStock {
			if gameBrowserOpen {
				println(webpages.BrowserAlreadyOpen)
			} else {
				gameBrowserOpen = true
				request.OpenBrowser(gameUrl)
			}
		}
	}()
}
