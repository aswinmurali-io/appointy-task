// Post definition
package main

import (
	"fmt"
)

type Post struct {
	Id              int    `bson:"_id" json:"_id"`
	Caption         string `bson:"caption" json:"caption"`
	ImageUrl        string `bson:"imageurl" json:"imageurl"`
	PostedTimestamp string `bson:"postedtimestamp" json:"postedtimestamp"`
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
