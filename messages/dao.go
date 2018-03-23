package dao

import (

	// Deployed API
	"gopkg.in/mgo.v2"
	"log"

)

type Messages struct (
	Server string
	Database string
)
var db *mgo.Database

const (
	COLLECTION = "messages"
)

func (m *MessagesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Establish a connection to database
func (m *MessagesDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of messages
func (m *MessagesDao) FindAll() ([]message, error) {
	var messages []message
	err := db.C(COLLECTION).Find(bson.M{}).All(&messages)
	return messages, err
}

// Find a message by its id
func (m *MessagesDao) FindById(id string) (message, error) {
	var message message
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&message)
	return message, err
}

// Insert a message into database
func (m *MessagesDao) Insert(message message) error {
	err := db.C(COLLECTION).Insert(&message)
	return err
}

// Delete an existing message
func (m *MessagesDao) Delete(message message) error {
	err := db.C(COLLECTION).Remove(&message)
	return err
}

// Update an existing message
func (m *MessagesDao) Update(message message) error {
	err := db.C(COLLECTION).UpdateId(message.ID, &message)
	return err
}