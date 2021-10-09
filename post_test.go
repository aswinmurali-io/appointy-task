package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestInsertPost(t *testing.T) {

	jsonValue, _ := json.Marshal(map[string]string{
		"caption":         "test caption",
		"imageurl":        "http://example.com/test.png",
		"postedtimestamp": "232323232434",
		"userid":          "6161315a42a6799fc59f7e30",
	})

	resp, err := http.Post(ApiUrl+"posts/", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	content := strings.Split(sb, " ")

	// Added user ObjectID(
	if content[0] != "Added" {
		t.Error("Failed to add post")
	}
	log.Println(sb)
}

func TestGetPost(t *testing.T) {
	resp, err := http.Get(ApiUrl + "posts/6161b6e17bf8c14019271e83")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	content := strings.Split(sb, ":")

	fmt.Println(content[0])

	if content[0] != `{"_id"` {
		t.Error("Post get format is incorrect")
	}
	log.Println(sb)
}
