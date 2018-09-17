package main

import (
	"fmt"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Whoa, Go is neat!")
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "About page")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)

	http.ListenAndServe(":8000", nil)
}
