package requests

import (
	"net/http"
	"io/ioutil"
)

type RClient struct {
	Headers map[string]string
	RootURL string
}

//ott-details.p.rapidapi.com

func NewClient(host, apiKey string) *RClient {
	headers := map[string]string{
		"x-rapidapi-host": host,
		"x-rapidapi-key": apiKey,
	}
	return &RClient{
		Headers: headers,
	}
}

func (r *RClient) Execute(url, method string) ([]byte, error) {
	var req *http.Request

	switch method {
	case "GET":
		req, _ = http.NewRequest("GET", url, nil)
	}
	for hx, hv := range r.Headers {
		req.Header.Add(hx, hv)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}