package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/erice5005/trailerrandomizer/requests"
	"github.com/erice5005/trailerrandomizer/models"
	ott "github.com/erice5005/trailerrandomizer/ott-api"
	// "net/url"
	// "log"
	"os"
	// "encoding/json"
	// "math/rand"
	"time"
)

type Server struct {
	rx *requests.RClient
	lastRequest time.Time
	requestsPerSecond int
}

func main() {

	sx := &Server{
		rx: requests.NewClient("ott-details.p.rapidapi.com", os.Getenv("api-key")),
		lastRequest: time.Now(),
		requestsPerSecond: 1,
	}

	router := mux.NewRouter()

	router.HandleFunc("/trailer", sx.GetTrailer).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}

func (s *Server) GetTrailer(w http.ResponseWriter, r *http.Request) {
	expectedParams := []string{
		"genre",
	}

	parameters := make(map[string]string)

	for _, qx := range expectedParams {
		parameters[qx] = r.URL.Query().Get(qx)
	}

	var trailer models.TrailerResult 
	var err error
	trailer, err = s.GetRandomTrailer(parameters["genre"])
	for len(trailer.TrailerURL) < 1 {
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		}
		time.Sleep(1 * time.Second)
	}



    w.Write([]byte(trailer.TrailerURL[0]))

}

func (s *Server) GetRandomTrailer(genre string) (models.TrailerResult, error) {
	
	var err error

	resp, err := ott.GetResultsByGenre(s.rx, genre, 1)

	if err != nil {
		return models.TrailerResult{}, err
	}

	if len(resp) == 0 {
		return models.TrailerResult{}, nil
	}

	var trailer models.TrailerResult

	targetIndex := 0
	page := 1
	for len(trailer.TrailerURL) < 1 {

		trailer, err = ott.GetAdditionalInfo(s.rx, resp[targetIndex].Imdbid)
		if err != nil {
			return models.TrailerResult{}, err
		}
		time.Sleep(1 * time.Second)

		if targetIndex == len(resp) - 1 {
			page++
			resp, err = ott.GetResultsByGenre(s.rx, genre, page)

			if err != nil {
				return models.TrailerResult{}, err
			}
		}
		targetIndex++
	}

	time.Sleep(1 * time.Second)
	
	

	return trailer, nil
}