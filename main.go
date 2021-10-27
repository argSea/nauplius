package api

import (
	"io"
	"net/http"
)

//Controller struct to handle API requests
type Controller struct{}

//API is our API handler
type API struct {
	Controller Controller
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := a.Controller.ParseURL(r.URL.Path)
	io.WriteString(w, data)
}

//ParseURL parses the URL
func (c *Controller) ParseURL(url string) string {
	return ("{\"path\":\"" + url + "\"}")
}
