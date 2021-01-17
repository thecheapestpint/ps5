package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RequestPage(url string) (string, bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", false, err
		}

		respUrl := resp.Request.URL.String()
		a := respUrl != url
		println(a)

		return string(bodyBytes), resp.Request.URL.String() != url, nil
	} else {
		return "", false, fmt.Errorf("%s", "Error fetching request")
	}
}
