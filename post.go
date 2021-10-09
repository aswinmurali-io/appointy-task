// Post definition
package main

import (
	"fmt"
)

type Post struct {
	Id              int
	Caption         string
	ImageUrl        string
	PostedTimestamp string
}

type Posts []Post

func (post Post) add() {
	fmt.Println("INFO: Inserting user")
	result, errorInInsert := Database.Collection("posts").InsertOne(
		MongoContext, post,
	)
	fmt.Println(result)
	fmt.Println(errorInInsert)
}

func (user Post) remove() {
	fmt.Println("INFO: Removing user")
	result, errorInDelete := Database.Collection("posts").DeleteOne(
		MongoContext, user,
	)
	fmt.Println(result)
	fmt.Println(errorInDelete)
}
