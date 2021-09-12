package imdb

import (
	"testing"
	"github.com/joho/godotenv"
)


func Test_GetURL(t *testing.T) {
	t.Parallel()

	cx := NewClient()

	tests := []struct{
		Name string
	}{
		{
			Name: "Happy Path: Correct compile",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Logf("Out: %v\n", cx.GetURL("path", "subpath", []string{}))
		})
	}
}

func Test_APIRequest(t *testing.T) {
	t.Parallel()

	err := godotenv.Load()
	if err != nil {
		t.Error(err)
	}

	tests := []struct{
		Name string
		// Client *IMDBClient
		// Path string
		Expression string
		ID string
		Options []string
	}{
		{
			Name: "Happy Path: Get By Keyword",
			Expression: "Comedy",
			ID: "",
			Options: []string{},
		},
	}

	for _, tt := range tests {

		cx := NewClient()

		t.Run(tt.Name, func(t *testing.T) {
			dataset, err := cx.SearchKeyword(tt.Expression)
			if err != nil  {
				t.Error(err)
			}
			if len(dataset.Results) < 1 {
				t.Error("no results")
			}
		})
	}
}