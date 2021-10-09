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

	if sb != `{"_id": {"$oid":"6161315a42a6799fc59f7e30"},"name": "Aswin Murali","email": "test@gmail.com","password": "123456","posts": [{"$oid":"6161b2ad7bf8c14019271e82"},{"$oid":"6161b6e17bf8c14019271e83"}]}` {
		t.Error("User get format is incorrect")
	}
	log.Println(sb)
}
