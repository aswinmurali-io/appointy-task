package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"strings"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func userGetDetailsPage(response http.ResponseWriter, request *http.Request) {
	_, errorInBody := ioutil.ReadAll(request.Body)
	idStringInHex := strings.Split(request.URL.Path, "/")[2]

	if errorInBody == nil {
		objectId, errorInHex := primitive.ObjectIDFromHex(idStringInHex)
		if errorInHex != nil {
			fmt.Println("ERROR: Unable to get object id from hex value.")
		}
		userBson, errorInJson := User{Id: objectId}.get().DecodeBytes()
		if errorInJson != nil {
			fmt.Println("ERROR: Error in reading json from content")
			return
		}
		fmt.Println(userBson)
		var doc bson.Raw
		bson.Unmarshal(userBson, &doc)
		fmt.Fprintf(response, "%s", doc.String())
		if errorInJson != nil {
			fmt.Fprintf(response, "ERROR: Unable to unmarshal this json")
			fmt.Println(errorInJson)
			return
		}
	} else {
		fmt.Println("ERROR: Error in reading request body.")
		fmt.Println(errorInBody)
	}
}

func userCreatePage(response http.ResponseWriter, request *http.Request) {
	body, errorInBody := ioutil.ReadAll(request.Body)
	if errorInBody == nil {
		keyVal := make(map[string]string)
		errorInUnmarshal := json.Unmarshal(body, &keyVal)

		// Json validation
		if errorInUnmarshal == nil {
			// Name validation
			if keyVal["name"] == "" || keyVal["name"] == " " {
				fmt.Fprintf(response, "ERROR: 'name' key cannot be empty!")
				return
			}

			// Email validation
			_, errorInParsingMail := mail.ParseAddress(keyVal["email"])
			if errorInParsingMail != nil {
				fmt.Fprintf(response, "ERROR: Invalid email address!")
				return
			}

			user := User{
				Id:       primitive.NewObjectID(),
				Name:     keyVal["name"],
				Email:    keyVal["email"],
				Password: keyVal["password"],
			}
			user.add()

			// print the user object they send as response
			json.NewEncoder(response).Encode(user)
		} else {
			fmt.Println("ERROR: Unable to unmarshal json value.")
			fmt.Println(errorInUnmarshal)
		}
	} else {
		fmt.Println("ERROR: Error in reading request body.")
		fmt.Println(errorInBody)
	}

}

// User route
func userPage(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		userGetDetailsPage(response, request)
	case "POST":
		userCreatePage(response, request)
	default:
		fmt.Fprintf(response, "Sorry, only POST and GET methods are supported.")
	}
}
