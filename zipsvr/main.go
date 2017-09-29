package main

import "fmt"
import "net/http"
import "log"
import "runtime"
import "encoding/json"

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
	// fmt.Println("Hello World!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler) // say what func to call when request to /hello
	mux.HandleFunc("/memory", memoryHandler)
	fmt.Printf("Server is Listening at http://localhost:4000\n")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}
