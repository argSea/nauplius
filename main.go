package api

import (
	"io"
	"net/http"
)

//Controller struct to handle API requests
type Controller struct{}

//ParseURL parses the URL
func (c *Controller) ParseURL(url string) string {
	return ("{\"path\":\"" + url + "\"}")
}

//Our API endpoints
type API struct {
	endpoint string
	function func(w http.ResponseWriter)
}

type APIEndpoint struct {
	endpoints []API
}

func (e *APIEndpoint) AddNewEndpoint(end string, function func(w http.ResponseWriter)) {
	a := API{endpoint: end, function: function}
	e.endpoints = append(e.endpoints, a)
}

func (e *APIEndpoint) FindEndpoint(end string) func(w http.ResponseWriter) {
	for _, value := range e.endpoints {
		if value.endpoint == end {
			return value.function
		}
	}

	return e.BaseFunction
}

func (e *APIEndpoint) BaseFunction(w http.ResponseWriter) {
	io.WriteString(w, "Hello World!")
}

//Listens and finds the route
type Router struct {
	Controller Controller
	endpoints  APIEndpoint
}

func (a *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := a.Controller.ParseURL(r.URL.Path)
	this_func := a.endpoints.FindEndpoint(data)

	this_func(w)
}
