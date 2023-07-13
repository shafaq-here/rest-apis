package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	/*The StrictSlash function in the mux package controls the behavior of the router when it comes to trailing slashes. When StrictSlash is set to true, the router will redirect requests that end with a trailing slash to the same path without the trailing slash. For example, a request to /articles/ will be redirected to /articles. When StrictSlash is set to false, the router will not redirect requests with trailing slashes.*/
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
