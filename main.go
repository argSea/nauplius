package api

import (
	"encoding/json"
	"io"
	"net/http"
)

//API Item and response struct, for returning our API data. API items is mutable
type APIResponse struct {
	status string
	code   int
	items  map[string]string
}

//Our API endpoints
type API struct {
	endpoint string
	function func(w http.ResponseWriter) *APIResponse
}

type APIEndpoint struct {
	endpoints []API
}

func (e *APIEndpoint) AddNewEndpoint(end string, function func(w http.ResponseWriter) *APIResponse) {
	a := API{endpoint: end, function: function}
	e.endpoints = append(e.endpoints, a)
}

//Controller struct to handle API requests
type Controller struct {
	APIEndpoint APIEndpoint
}

//ParseURL parses the URL
func (c *Controller) ParseURL(w http.ResponseWriter, url string) string {
	io.WriteString(w, url)

	//TODO: Strip path
	this_func := c.APIEndpoint.FindEndpoint(url, w)

	response, _ := json.Marshal(this_func(w))

	return string(response)
}

func (e *APIEndpoint) FindEndpoint(end string, w http.ResponseWriter) func(w http.ResponseWriter) *APIResponse {
	for _, value := range e.endpoints {
		io.WriteString(w, value.endpoint)
		if value.endpoint == end {
			return value.function
		}
	}

	return e.BaseFunction
}

func (e *APIEndpoint) BaseFunction(w http.ResponseWriter) *APIResponse {
	io.WriteString(w, "Hello World!")

	return &APIResponse{}
}

//Listens and finds the route
type Router struct {
	Controller Controller
}

func (a *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	a.Controller.ParseURL(w, r.URL.Path)
}
