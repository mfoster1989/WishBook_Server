package messages

import (
	"gopkg.in/mgo.v2/bson"
)

// Messages building documents
type Messages struct {
	ID      bson.ObjectId `bson:"_id"`
	FBID    string        `bson:"_id" json:"id"`
	Token   string        `bson:"token" json:"token"`
	Message string        `bson:"message" json:"message"`
}
