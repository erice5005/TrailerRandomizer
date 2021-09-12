package ott-api

import (
	"github.com/erice5005/trailerrandomizer/requests"
)

func GetResultsByGenre(rx *requests.RClient, genre string) ([]*models.Result, error) {
	u, err := url.Parse("https://ott-details.p.rapidapi.com/advancedsearch")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	for pi, px := range parameters {
		// rootUrl += pix +"="+px
		q.Set(pi, px)
	}
	u.RawQuery = q.Encode()
	resp, err := s.rx.Execute(u.String(), "GET")

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

func GetAdditionalInfo(rx *requests.RClient, id string) (*models.TrailerResult, error) {
	u, err :=  url.Parse("https://ott-details.p.rapidapi.com/getadditionalDetails")
	if err != nil {
		log.Fatal(err)
	}
	q = u.Query()
	q.Set("imdbid", id)
	u.RawQuery = q.Encode()
	resp, err = s.rx.Execute(u.String(), "GET")

	if err != nil {
		return err
	}

	var trailer models.TrailerResult
	err = json.Unmarshal(resp, &trailer)

	if err != nil {
		return err
	}

	return trailer, nil
}