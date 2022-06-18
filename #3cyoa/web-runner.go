package main

import (
	"net/http"
)

type WebRunner struct {
}

func (wr WebRunner) Start(provider *StoryArcProvider) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":8080", nil)

}

func (wr WebRunner) rootEndpointHandler(provider *StoryArcProvider, w http.ResponseWriter, req *http.Request) {

	arcValues := req.URL.Query()["arc"]
	arcEndpoint := "intro"

	if len(arcValues) > 0 {
		arcEndpoint = arcValues[0]
	}

	_, err := provider.WriteTemplateText(w, arcEndpoint)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
