package models

import (
	// "encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"Username"`
	Nama     string        `bson:"Nama"`
	Email    string        `bson:"Email"`
	Password []byte        `bson:"Password"`
	Jk       string        `bson:"Jk"`
	Image    string        `bson:"Image"`
}

type Home struct {
	Nama  string  `bson:"Nama"`
	Index []Users `bson:"Index"`
}

type Session struct {
	Username     string
	LastActivity time.Time
}
