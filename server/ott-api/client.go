package ottapi

import (
	"github.com/erice5005/trailerrandomizer/requests"
	"github.com/erice5005/trailerrandomizer/models"
	"log"
	"encoding/json"
	"net/url"
	"strconv"
)

func GetResultsByGenre(rx *requests.RClient, genre string, page int) ([]models.Result, error) {
	u, err := url.Parse("https://ott-details.p.rapidapi.com/advancedsearch")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	// for pi, px := range parameters {
	// 	// rootUrl += pix +"="+px
	// 	q.Set(pi, px)
	// }
	q.Set("genre", genre)
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()
	resp, err := rx.Execute(u.String(), "GET")

	if err != nil {
		return nil, err
	}

	var Response models.ResultResponse

	err = json.Unmarshal(resp, &Response)
	if err != nil {
		return nil, err
	}

	return Response.Results, nil
}

func GetAdditionalInfo(rx *requests.RClient, id string) (models.TrailerResult, error) {
	u, err :=  url.Parse("https://ott-details.p.rapidapi.com/getadditionalDetails")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("imdbid", id)
	u.RawQuery = q.Encode()
	resp, err := rx.Execute(u.String(), "GET")

	if err != nil {
		return models.TrailerResult{}, err
	}

	var trailer models.TrailerResult
	err = json.Unmarshal(resp, &trailer)

	if err != nil {
		return models.TrailerResult{}, err
	}

	return trailer, nil
}