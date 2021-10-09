// Post definition
package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id              primitive.ObjectID `bson:"_id" json:"_id"`
	Caption         string             `bson:"caption" json:"caption"`
	ImageUrl        string             `bson:"imageurl" json:"imageurl"`
	PostedTimestamp string             `bson:"postedtimestamp" json:"postedtimestamp"`
	UserId          primitive.ObjectID `bson:"userid" json:"userid"`
}

type Posts []Post

func (post Post) add() {
	Collection := Database.Collection("users")
	fmt.Println("INFO: Inserting user")
	result, errorInInsert := Database.Collection("posts").InsertOne(
		MongoContext, post,
	)

	// Append the post ID into newuser
	newuser := User{}
	errorInDecode := Collection.FindOne(
		MongoContext, bson.M{"_id": post.UserId},
	).Decode(&newuser)
	if errorInDecode != nil {
		log.Println(errorInDecode)
	}
	newuser.Posts = append(newuser.Posts, post.Id)
	log.Println(newuser.Posts)

	// Update user
	Collection.DeleteOne(MongoContext, bson.M{"_id": post.UserId})
	Collection.InsertOne(MongoContext, newuser)

	log.Println(result)
	log.Println(errorInInsert)
}

func (post Post) get() Post {
	Collection := Database.Collection("posts")
	log.Printf("[INFO] Getting post from id %s.\n", post.Id)

	newpost := Post{}
	errorInDecode := Collection.FindOne(
		MongoContext, bson.M{"_id": post.Id},
	).Decode(&newpost)

	if errorInDecode != nil {
		log.Println(errorInDecode)
	}

	return newpost
}
