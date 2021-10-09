package main

import (
	"fmt"
	"net/http"
)

func postPage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Endpoint")
}
