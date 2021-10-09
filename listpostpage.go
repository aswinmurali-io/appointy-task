package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func listPosts(response http.ResponseWriter, request *http.Request) {
	_, errorInBody := ioutil.ReadAll(request.Body)
	if errorInBody == nil {
		idStringInHex := strings.Split(request.URL.Path, "/")[3]
		fmt.Println(idStringInHex)
		objectId, errorInHex := primitive.ObjectIDFromHex(idStringInHex)
		if errorInHex != nil {
			log.Println(errorInHex)
		}
		fmt.Fprintf(response, "%s", User{Id: objectId}.listPosts())
	}
}

func listPostsPage(response http.ResponseWriter, request *http.Request) {
	log.Printf("[INFO] Request %s triggered", request.Method)
	switch request.Method {
	case "GET":
		listPosts(response, request)
	default:
		fmt.Fprintf(response, "Sorry, only GET method is supported.")
	}
}
