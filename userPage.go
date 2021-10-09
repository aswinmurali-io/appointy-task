package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// User route
func userPage(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		body, errorInBody := ioutil.ReadAll(request.Body)
		if errorInBody == nil {
			keyVal := make(map[string]string)
			errorInUnmarshal := json.Unmarshal(body, &keyVal)
			if errorInUnmarshal == nil {
				id, errorInId := strconv.Atoi(keyVal["Id"])
				if errorInId == nil {
					user := User{
						Id:       id,
						Name:     keyVal["Name"],
						Email:    keyVal["Email"],
						Password: keyVal["Password"],
					}
					user.add()
					json.NewEncoder(response).Encode(user)
				} else {
					fmt.Println("ERROR: Unable to get user id from id value in form.")
					fmt.Println(errorInId)
				}
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
				Id:       0,
				Name:     "name",
				Email:    "example@mail.com",
				Password: "password",
			},
		})
	}

}
