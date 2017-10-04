package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/abourn/info344-in-class/zipsvr/models"
)

// create handler to return zip codes for given city

// a http handler only has two parameters so we are going to create a handler struct

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

// parameter on left is called 'receiver parameter'
// basically way that Go does 'this'
// this is how you get access to the Index when you are restricted by having the two parameters 'w' and 'r'
func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// URL we want to support: /zips/city-name
	cityName := r.URL.Path[len(ch.PathPrefix):] // get city-name from URL
	cityName = strings.ToLower(cityName)        // case insensitive search

	// if user didn't suppy a city name...
	if len(cityName) == 0 {
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		return // stop request and exit
	}
	w.Header().Add(accessControlAllowOrigin, "*")
	w.Header().Add(headerContentType, contentTypeJSON)
	zips := ch.Index[cityName] // Zip index, so we can use the city name as a key (this is our map)
	json.NewEncoder(w).Encode(zips)
}
