// User definition
package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}

type Users []User

func (user User) add() {
	fmt.Println("INFO: Inserting user")
	result, errorInInsert := Database.Collection("users").InsertOne(
		MongoContext, user,
	)
	fmt.Println(result)
	fmt.Println(errorInInsert)
}

func (user User) remove() {
	fmt.Println("INFO: Removing user")
	result, errorInDelete := Database.Collection("users").DeleteOne(
		MongoContext, user,
	)
	fmt.Println(result)
	fmt.Println(errorInDelete)
}
