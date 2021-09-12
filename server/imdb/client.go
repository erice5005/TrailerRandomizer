package imdb


import (
	"net/http"
	"log"
	"io/ioutil"
)


func NewRequest(url, method string, body []byte) *http.Request {
	log.Printf(": %v\n", url)

	var req *http.Request

	switch method {
	case "GET":
		req, _ = http.NewRequest("GET", url, nil)
	}

	return req
}


func DoRequest(c *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(c)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

