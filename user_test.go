package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

const ApiUrl = "http://localhost:5000/"

func TestInsertUser(t *testing.T) {

	jsonValue, _ := json.Marshal(map[string]string{
		"name":     "Aswin Murali",
		"email":    "example@mail.com",
		"password": "123445",
	})

	resp, err := http.Post(ApiUrl+"users/", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	content := strings.Split(sb, " ")

	// Added post ObjectID(....
	if content[0] != "Added" {
		t.Error("Failed to add post")
	}
	log.Println(sb)
}

func TestGetUser(t *testing.T) {
	resp, err := http.Get(ApiUrl + "users/6161315a42a6799fc59f7e30")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	content := strings.Split(sb, " ")

	// {"_id": {"$oid":"6161315a42a679....
	if content[0] != `{"_id":` {
		t.Error("User get format is incorrect")
	}
	log.Println(sb)
}

func TestListPostsFromUser(t *testing.T) {
	resp, err := http.Get(ApiUrl + "posts/users/6161315a42a6799fc59f7e30")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	content := strings.Split(sb, " ")

	// [ObjectID("6161b2ad7bf8c14019271e82") .... ]
	if content[0] != `[ObjectID("6161b2ad7bf8c14019271e82")` {
		t.Error("List of posts seems different. Did the logic change? Test data must remain constant")
	}
	log.Println(sb)
}
