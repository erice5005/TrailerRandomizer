package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/erice5005/trailerrandomizer/requests"
	"github.com/erice5005/trailerrandomizer/models"
	"net/url"
	"log"
	"os"
	"encoding/json"
	"math/rand"
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

	resp, err := ott.GetResultsByGenre(s.rx, parameters["genre"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
   		w.Write([]byte("500 - Something bad happened!"))
	}

	resultIndex := 0
	if len(Response.Results) > 0 {
		resultIndex = rand.Intn(len(Response.Results) - 1)
	}

	targetResult := Response.Results[resultIndex]

	time.Sleep(2 * time.Second)
	// marshed, _ := json.Marshal(targetResult)


	
	trailer, err := ott.GetAdditionalInfo(s.rx, targetResult.Imdbid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
   		w.Write([]byte("500 - Something bad happened!"))
	}


	marshed, _ := json.Marshal(trailer)

    w.Write(marshed)

}