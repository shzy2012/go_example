package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Results []string `json:"results"`
	Nlg     string   `json:"nlg"`
}

func NewResponse(results []string, nlg string) *Response {
	return &Response{
		Results: results,
		Nlg:     nlg,
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		keys := []string{"go", "C#"}
		res := NewResponse(keys, "i am from server")
		bts, _ := json.Marshal(res)

		w.Header().Set("Context-Type", "application/json")
		w.Write(bts)
	})
	http.ListenAndServe(":8080", nil)
}
