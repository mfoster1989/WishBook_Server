package messages

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "wishbookserver"

// DOCNAME the name of the document
const DOCNAME = "messages"

// GetMessages returns the list of Messages
func (r Repository) GetMessages() Messages {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Messages{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

// AddMessages inserts a message in the DB
func (r Repository) AddMessage(messages Messages) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	messages.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(messages)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Updatemessage updates an message in the DB (not used for now)
func (r Repository) UpdateMessage(messages Messages) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	session.DB(DBNAME).C(DOCNAME).UpdateId(messages.ID, messages)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Deletemessage deletes an message (not used for now)
func (r Repository) DeleteMessage(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	// Write status
	return "OK"
}