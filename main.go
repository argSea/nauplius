package api

import (
	"encoding/json"
	"io"
	"net/http"
)

//API Item and response struct, for returning our API data. API items is mutable
type APIResponse struct {
	status  string
	code    int
	message string
	items   map[string]interface{}
}

func (r *APIResponse) SetStatus(status string) {
	r.status = status
}

func (r *APIResponse) SetCode(code int) {
	r.code = code
}

func (r *APIResponse) SetMessage(message string) {
	r.message = message
}

func (r *APIResponse) SetItems(items map[string]interface{}) {
	r.items = items
}

//Our API endpoints
type API struct {
	endpoint string
	function func(r *APIResponse)
}

type APIEndpoint struct {
	endpoints []API
}

func (e *APIEndpoint) AddNewEndpoint(end string, function func(r *APIResponse)) {
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

	apiResponse := APIResponse{}
	this_func(&apiResponse)
	response, _ := json.Marshal(apiResponse)

	return string(response)
}

func (e *APIEndpoint) FindEndpoint(end string, w http.ResponseWriter) func(r *APIResponse) {
	for _, value := range e.endpoints {
		io.WriteString(w, value.endpoint)
		if value.endpoint == end {
			return value.function
		}
	}

	return e.BaseFunction
}

func (e *APIEndpoint) BaseFunction(r *APIResponse) {
	// io.WriteString(w, "Hello World!")

	items := make(map[string]interface{})

	r.SetStatus("error")
	r.SetCode(404)
	r.SetMessage("API Not Found")
	r.SetItems(items)
}

//Listens and finds the route
type Router struct {
	Controller Controller
}

func (a *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := a.Controller.ParseURL(w, r.URL.Path)

	io.WriteString(w, response)
}
