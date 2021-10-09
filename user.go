// User definition
package main

import (
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
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
