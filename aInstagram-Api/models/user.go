package models

import (
	"gopkg.in/mgo.v2/bson"
)

// models tell the datastructure off golang project
//postman will send json to request golang(used Id) but mongo db takes (_id)
//structure to store user data
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	EmailId  string        `json:"emailid" bson:"emailid"`
	Password string        `json:"password" bson:"password"`
}
