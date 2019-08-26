package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Contact encapsulates a contact in sampark
type Contact struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name,omitempty"`
	Email string        `json:"email" bson:"email"`
}
