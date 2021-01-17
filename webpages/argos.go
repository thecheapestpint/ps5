package webpages

import (
	"fmt"
	"ps5/request"
	"strings"
)

func Argos(url string) bool {

	body, redirected, err := request.RequestPage(url)
	if err != nil {
		println(fmt.Sprintf(fetchPageError, "Argos ") + err.Error())
	}

	if redirected == false {
		println(fmt.Sprintf("Argos is no longer sending us to a OOS page:" + url))
		return true
	}

	if !strings.Contains(body, "Sorry, PlayStation®5 is currently unavailable.") {
		return true
	}
	return false
}
