package messages

import (
	"gopkg.in/mgo.v2/bson"
)

type Messages struct {
	ID      bson.ObjectId `bson:"_id"`
	FBID    int64         `bson:"_id" json:"id"`
	token   int64         `bson:"token" json:"token"`
	Message string        `bson:message json:"message"`
}
