package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User route
func userPage(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		body, errorInBody := ioutil.ReadAll(request.Body)
		if errorInBody == nil {
			keyVal := make(map[string]string)
			errorInUnmarshal := json.Unmarshal(body, &keyVal)

			// Json validation
			if errorInUnmarshal == nil {
				// Name validation
				if keyVal["name"] == "" || keyVal["name"] == " " {
					fmt.Fprintf(response, "ERROR: 'name' key cannot be empty!")
					break
				}

				// Email validation
				_, errorInParsingMail := mail.ParseAddress(keyVal["email"])
				if errorInParsingMail != nil {
					fmt.Fprintf(response, "ERROR: Invalid email address!")
					break
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

	default:
		fmt.Fprintf(response, "Sorry, only POST methods are supported.")
		fmt.Fprintf(response, "Expecting syntax:-")
		json.NewEncoder(response).Encode(Users{
			User{
				Name:     "name",
				Email:    "example@mail.com",
				Password: "password",
			},
		})
	}

}
