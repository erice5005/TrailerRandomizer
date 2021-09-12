package models

type ResultResponse struct {
	Page    int `json:"page"`
	Results []Result `json:"results"`
}

type Result struct {
	Imageurl   []string    `json:"imageurl"`
	Genre      []string    `json:"genre"`
	Imdbid     string      `json:"imdbid"`
	Title      string      `json:"title"`
	Imdbrating interface{} `json:"imdbrating"`
	Released   int         `json:"released"`
	Type       string      `json:"type"`
	Synopsis   string      `json:"synopsis,omitempty"`
}

type TrailerResult struct {
	Imdbid string `json:"imdbid"`
	Title  string `json:"title"`
	People []struct {
		Peopleid   string      `json:"peopleid"`
		Characters []string    `json:"characters"`
		Category   string      `json:"category"`
		Job        interface{} `json:"job"`
	} `json:"people"`
	Quotes      []interface{} `json:"quotes"`
	PlotSummary string        `json:"plotSummary"`
	Reviews     []interface{} `json:"reviews"`
	TrailerURL  []string `json:"trailerUrl"`
}