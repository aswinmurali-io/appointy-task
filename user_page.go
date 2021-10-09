package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const RequestBodyErrorMsg = "Unable to read the request!"

func userGetDetailsPage(response http.ResponseWriter, request *http.Request) {
	_, errorInBody := ioutil.ReadAll(request.Body)
	idStringInHex := strings.Split(request.URL.Path, "/")[2]

	if errorInBody == nil {
		objectId, errorInHex := primitive.ObjectIDFromHex(idStringInHex)
		if errorInHex != nil {
			fmt.Println(response, "Unable to get object id from hex value.")
			log.Printf("[ERROR] Unable to get object id from hex value %s.\n", idStringInHex)
		}
		userBson, errorInBson := User{Id: objectId}.get().DecodeBytes()
		if errorInBson != nil {
			fmt.Fprintf(response, "Error in reading json from content")
			log.Printf("[ERROR] Error in reading json from content %s.\n", userBson)
			return
		}

		// Output the response
		var doc bson.Raw
		bson.Unmarshal(userBson, &doc)
		fmt.Fprintf(response, "%s", doc.String())
		log.Println(userBson)

		if errorInBson != nil {
			fmt.Fprintf(response, "Unable to unmarshal this bson")
			log.Printf("[ERROR] Unable to unmarshal this bson\n")
			log.Println(errorInBson)
			return
		}
	} else {
		fmt.Fprintf(response, "%s", RequestBodyErrorMsg)
		log.Println("[ERROR] Error in reading request body.")
		log.Println(errorInBody)
	}
}

// Create user using the POST method
func userCreatePage(response http.ResponseWriter, request *http.Request) {
	// Reading request body
	body, errorInBody := ioutil.ReadAll(request.Body)
	if errorInBody == nil {
		// Json validation
		keyVal := make(map[string]string)
		errorInUnmarshal := json.Unmarshal(body, &keyVal)

		if errorInUnmarshal == nil {
			// Name validation
			if keyVal["name"] == "" || keyVal["name"] == " " {
				fmt.Fprintf(response, "Key `name` has empty value!")
				log.Printf("[ERROR] 'name' key cannot be empty! %s.\n", keyVal)
				return
			}

			// Email validation
			_, errorInParsingMail := mail.ParseAddress(keyVal["email"])
			if errorInParsingMail != nil {
				fmt.Fprintf(response, "Invalid email address!")
				log.Printf("[ERROR] Invalid email address! %s.\n", keyVal["email"])
				return
			}

			user := User{
				Id:       primitive.NewObjectID(),
				Name:     keyVal["name"],
				Email:    keyVal["email"],
				Password: keyVal["password"],
			}
			user.add()
			fmt.Fprintf(response, "Added user %s", user.Id)

			// print the user object they send as response
			json.NewEncoder(response).Encode(user)
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
func userPage(response http.ResponseWriter, request *http.Request) {
	log.Printf("[INFO] Request %s triggered", request.Method)
	switch request.Method {
	case "GET":
		userGetDetailsPage(response, request)
	case "POST":
		userCreatePage(response, request)
	default:
		fmt.Fprintf(response, "Sorry, only POST and GET methods are supported.")
	}
}
