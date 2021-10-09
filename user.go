// User definition
package main

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       primitive.ObjectID   `bson:"_id" json:"_id"`
	Name     string               `bson:"name" json:"name"`
	Email    string               `bson:"email" json:"email"`
	Password string               `bson:"password" json:"password"`
	Posts    []primitive.ObjectID `bson:"posts" json:"posts"`
}

type Users []User

func (user User) add() {
	log.Printf("[INFO] Inserting user %s.\n", user)
	result, errorInInsert := Database.Collection("users").InsertOne(
		MongoContext, user,
	)
	log.Println(result)
	log.Println(errorInInsert)
}

func (user User) get() *mongo.SingleResult {
	log.Printf("[INFO] Getting user from id %s.\n", user.Id)
	result := Database.Collection("users").FindOne(
		MongoContext, bson.M{"_id": user.Id},
	)
	log.Println(result)
	return result
}

func (user User) listPosts() []primitive.ObjectID {
	Collection := Database.Collection("users")
	log.Printf("[INFO] Getting posts for user id %s.\n", user.Id)

	updatedUser := User{}

	Collection.FindOne(MongoContext, bson.M{"_id": user.Id}).Decode(&updatedUser)

	return updatedUser.Posts
}
