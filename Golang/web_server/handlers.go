package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello World!")
}

func HandleHome(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Home Page!")
}
