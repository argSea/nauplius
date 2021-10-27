package api

import (
	"net/http"
)

//Controller struct to handle API requests
type Controller struct{}

//ParseURL parses the URL
func (c *Controller) ParseURL(url string) string {
	return ("{\"path\":\"" + url + "\"}")
}

//API is our API handler
type API struct {
	Controller Controller
	data       string
	rWriter    http.ResponseWriter
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	a.data = a.Controller.ParseURL(r.URL.Path)
	a.rWriter = w
}

func (a *API) GetHTTPWriter() http.ResponseWriter {
	return a.rWriter
}

func (a *API) GetData() string {
	return a.data
}
