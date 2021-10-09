package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func postGetDetailsPage(response http.ResponseWriter, request *http.Request) {
	_, errorInBody := ioutil.ReadAll(request.Body)
	idStringInHex := strings.Split(request.URL.Path, "/")[2]

	if errorInBody == nil {
		objectId, errorInHex := primitive.ObjectIDFromHex(idStringInHex)
		if errorInHex != nil {
			fmt.Println(response, "Unable to get object id from hex value.")
			log.Printf("[ERROR] Unable to get object id from hex value %s.\n", idStringInHex)
		}
		post := Post{Id: objectId}.get()
		fmt.Println(post)
		json.NewEncoder(response).Encode(post)
	} else {
		fmt.Fprintf(response, "%s", RequestBodyErrorMsg)
		log.Println("[ERROR] Error in reading request body.")
		log.Println(errorInBody)
	}
}

// Create user using the POST method
func postCreatePage(response http.ResponseWriter, request *http.Request) {
	// Reading request body
	body, errorInBody := ioutil.ReadAll(request.Body)
	if errorInBody == nil {
		// Json validation
		keyVal := make(map[string]string)
		errorInUnmarshal := json.Unmarshal(body, &keyVal)

		if errorInUnmarshal == nil {

			userId, errorFromUserId := primitive.ObjectIDFromHex(keyVal["userid"])
			if errorFromUserId != nil {
				log.Println(errorFromUserId)
			}
			post := Post{
				Id:              primitive.NewObjectID(),
				Caption:         keyVal["caption"],
				ImageUrl:        keyVal["imageurl"],
				PostedTimestamp: keyVal["postedtimestamp"],
				UserId:          userId,
			}

			post.add()
			fmt.Fprintf(response, "Added post %s", post.Id)

			// print the post object they send as response
			json.NewEncoder(response).Encode(post)
		} else {
			log.Printf("[ERROR] Unable to unmarshal json value %s.\n", body)
			log.Println(errorInUnmarshal)
			return
		}
	} else {
		fmt.Fprintf(response, "%s", RequestBodyErrorMsg)
		log.Printf("[ERROR] Error in reading request body %s.\n", request.Body)
		log.Println(errorInBody)
	}
}

// User route handler which binds logic for POST and GET method.
func postPage(response http.ResponseWriter, request *http.Request) {
	log.Printf("[INFO] Request %s triggered", request.Method)
	switch request.Method {
	case "GET":
		postGetDetailsPage(response, request)
	case "POST":
		postCreatePage(response, request)
	default:
		fmt.Fprintf(response, "Sorry, only POST and GET methods are supported.")
	}
}
