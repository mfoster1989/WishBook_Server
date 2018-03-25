package messages

import (
	"gopkg.in/mgo.v2/bson"
)

// Messages building documents
type Messages struct {
	ID      bson.ObjectId `bson:"_id"`
	FBID    int64         `bson:"_id" json:"id"`
	Token   int64         `bson:"token" json:"token"`
	Message string        `bson:"message" json:"message"`
}
