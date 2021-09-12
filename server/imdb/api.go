package imdb

import (
	// "log"
	"os"
	"encoding/json"
	"strconv"
	"errors"
)

type IMDBClient struct {
	publicKey string
	rootUrl string
	lang string	
	requests *RequestTracking
}

type RequestTracking struct {
	requests int
	limit int
}

func (rt *RequestTracking) Inc() {
	rt.requests++
}

func (rt *RequestTracking) IsLocked() bool {
	return rt.requests >= rt.limit
}



func loadOptsForKeys(opts []string) map[string]string {
	keys := make(map[string]string)

	// log.Printf("Env: %v\n", os.Environ())

	for _, ox := range opts {
		// if val, ok := os.LookupEnv(ox); ok {
			keys[ox] = os.Getenv(ox)
			// log.Printf("Key: %v\n", os.Getenv(ox))
		// }
	}

	return keys
}

func NewClient() *IMDBClient {
	keys := loadOptsForKeys([]string{
		"PublicKey",
		"ImdbRootUrl",
		"RequestStarting",
		"RequestLimit",
	})

	rs := keys["RequestStarting"]
	rl := keys["RequestLimit"]

	rsConv, _ := strconv.Atoi(rs)
	rlConv, _ := strconv.Atoi(rl)

	return &IMDBClient{
		publicKey: keys["PublicKey"],
		rootUrl: keys["ImdbRootUrl"],
		lang: "en",
		requests: &RequestTracking{
			requests: rsConv,
			limit: rlConv,
		},
	}
}

func (c *IMDBClient) GetURL(path, subpath string, options []string) string {
	ret := c.rootUrl + "/" + c.lang + "/" + "API" + "/" + path + "/" + c.publicKey
	if subpath != "" {
		ret = ret + "/"+ subpath
	}
	if len(options) > 0 {
		ret = ret + "/"
		for i, opt := range options {
			ret += opt
			if i < len(options) - 1 {
				ret = ret + ","
			}
		}
	}
	

	return ret
}

func (c *IMDBClient) DoRequest(path string, subpath string, opts []string, method string, body []byte) ([]byte, error) {
	if c.requests.IsLocked() {
		return []byte{}, errors.New("maxed out requests")
	}
	req := NewRequest(c.GetURL(path, subpath, opts), method, body)
	data, err := DoRequest(req)
	if err != nil {
		return nil, err
	}
	c.requests.Inc()
	return data, nil
}


type IMDB interface {
	SearchKeyword(string) (*KeywordSearchResponse, error)
	GetByKeywordID(string) (*KeywordElementResponse, error)
	GetReviewForID(string) (*ReviewsResponse, error)
	GetTitle(string) (*TitleResponse, error)
}

func (c *IMDBClient) SearchKeyword(keyword string) (*KeywordSearchResponse, error) {
	data, err := c.DoRequest("SearchKeyword", keyword, []string{}, "GET", []byte{})

	if err != nil {
		return nil, err
	}

	var resp *KeywordSearchResponse

	err = json.Unmarshal(data, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *IMDBClient) GetByKeywordID(id string) (*KeywordElementResponse, error) {
	data, err := c.DoRequest("Keyword", id, []string{}, "GET", []byte{})

	if err != nil {
		return nil, err
	}
	
	var resp *KeywordElementResponse

	err = json.Unmarshal(data, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *IMDBClient) GetReviewForID(id string) (*ReviewsResponse, error) {
	data, err := c.DoRequest("Reviews", id, []string{}, "GET", []byte{})

	if err != nil {
		return nil, err
	}
	
	var resp *ReviewsResponse

	err = json.Unmarshal(data, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *IMDBClient) GetTitle(id string) (*TitleResponse, error) {
	data, err := c.DoRequest("Title", id, []string{"Reviews", "Trailer"}, "GET", []byte{})

	if err != nil {
		return nil, err
	}
	
	var resp *TitleResponse

	err = json.Unmarshal(data, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}