package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/abourn/info344-in-class/zipsvr/handlers"
	"github.com/abourn/info344-in-class/zipsvr/models"
)

const zipsPath = "/zips/"

// import models package

/*
	helloHandler
		what to do when /hello is requested!

	memoryHandler
		heap = general ram
		stack = function in the ram
	main
		mux is the thing that decides which function to call
		mux.HanldeFunc specifies which function at which resource path
*/

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello %s!", name) // replaces %s with name.  This means you can construct a query like loaclhost:4000/hello?name=Adam
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC() // force garbage collector to run

	stats := &runtime.MemStats{} // creating an instance of a structure. initializer is the curly braces. & means I'm creating struct on heap and getting a pointer to it
	runtime.ReadMemStats(stats)
	// now, lets show stats to client as a json!
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats) // create new encoder, pass that to Encode, and then add to response stream
}

func main() {

	addr := os.Getenv("ADDR")
	// if ADDR not supplied, let's default to somethign
	if len(addr) == 0 {
		addr = ":80" // accept communications from any computer
	}

	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatal("error loading zips: %v", err) // exits process. crashes your server. don't do this inside HTTP handler stuff lol
	}
	log.Printf("loaded %d zips", len(zips))

	// let's create a map so that we can get the zip codes for Seattle

	cityIndex := models.ZipIndex{} // because we declared ZipIndex in models package
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)                   // make sure lowercase version
		cityIndex[cityLower] = append(cityIndex[cityLower], z) // ZipIndex is a map from string to slice of Zips, so we access the underlying slice with cityIndex[cityLower], and then appending the pointer of the Zip to that slice
	}

	// fmt.Println("Hello World!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler) // say what func to call when request to /hello
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath,
	}

	mux.Handle(zipsPath, cityHandler)

	fmt.Printf("Server is Listening at http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
