package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
 *
 * handle the request for /hello route
 */
func helloHandler(w http.ResponseWriter, r *http.Request){
//check path
if r.URL.Path != "/hello" {
	http.Error(w, "4044 not found", http.StatusNotFound)
}

//chech method
if r.Method != "GET" {
	http.Error(w, "method is not supported", http.StatusNotFound)
}

//printout
fmt.Fprintf(w, "hello!")
}

/**
 *
 * handle the request for /form route
 */
 func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseFom() err: %v", err)
	}

	fmt.Fprintf(w, "POST request successful")

	name := r.FormValue("name")
	message := r.FormValue("message")

	fmt.Fprint(w, "Name = %s\n", name)
	fmt.Fprint(w, "message = %s\n", message)
 }

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !=nil {
			log.Fatal(err)
	}

}