package webpages

import (
	"fmt"
	"ps5/request"
	"time"
)

func Game(url string) bool {

	_, redirected, err := request.RequestPage(url)
	if err != nil {
		println(fmt.Sprintf(fetchPageError, "Argos ") + err.Error())
	}

	if redirected == false {
		println(fmt.Sprintf("Game is no longer out of stock" + url))
		return true
	}

	println(fmt.Sprintf(debugStillOutOfStock, "GAME", time.Now().Format("2006-01-02 15:04:05")))
	return false

}
