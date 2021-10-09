package models

import (
	"gopkg.in/mgo.v2/bson"
)

//structure to store posts created by the user
type Post struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Caption  string        `json:"caption" bson:"caption"`
	ImageURL string        `json:"imageurl" bson:"imageurl"`
	User     *User         `json:"user"`
}
