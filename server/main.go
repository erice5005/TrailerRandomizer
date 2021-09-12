package main

import (
	"github.com/gorilla/mux"
	"net/http"
	// "github.com/erice5005/trailerrandomizer/requests"
	// "github.com/erice5005/trailerrandomizer/models"
	imdb "github.com/erice5005/trailerrandomizer/imdb"
	// "net/url"
	"log"
	// "os"
	// "encoding/json"
	"math/rand"
	"time"
	"strconv"

	"github.com/joho/godotenv"
)

type Server struct {
	rx *imdb.IMDBClient
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	
	sx := &Server{
		rx: imdb.NewClient(),	
	}

	router := mux.NewRouter()

	router.HandleFunc("/trailer", sx.GetTrailer).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}

func (s *Server) GetTrailer(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())

	// keyword = r.URL.Query().Get("keyword")
	keywords := []string{
		"dramas",
		"romantic",
	}
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("500 - Something bad happened!"))
	// }

	items := make([]*imdb.KeywordItem, 0)
	
	for _, kx := range keywords {
		dataset, err := s.rx.GetByKeywordID(kx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		}
		// log.Printf("Dataset: %v\n", dataset)
		items = append(items, dataset.Items...)
	}

	if len(items) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	candidates := make([]*imdb.KeywordItem, 0)

	for _, ix := range items {
		rx, _ := strconv.Atoi(ix.ImDbRating)
		if rx > 5 {
			log.Printf("ID: %v, Item: %v\n", ix.ID, ix.Title)
			candidates = append(candidates, ix)
		}
	}

	// log.Printf("Candidates: %v\n", candidates)

	if len(candidates) == 0 {
		// if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
		// }
	}

	indx := 0
	if len(candidates) > 1 {
		indx = rand.Intn(len(candidates)-1)
	}

	titleInfo, err := s.rx.GetTitle(candidates[indx].ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}
	
	log.Printf("Title Info: %v\n", titleInfo)
    w.Write([]byte(titleInfo.Trailer.Link))

}