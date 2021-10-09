package main

import (
	"fmt"
	"net/http"
)

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Endpoint")
}
