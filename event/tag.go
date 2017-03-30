package event

import (
	"zion/db"

	"gopkg.in/mgo.v2/bson"
)

type (
	TagEvent struct {
		EventId     int           `json:"eventId,omitempty"`
		Tag         *db.Tag       `json:"tag,omitempty"`
		FeedId      bson.ObjectId `json:"feedId,omitempty"`
		TagIds      []string      `json:"tags,omitempty"`
		AddUser     bson.ObjectId `json:"addUser,omitempty"`
		InitiatorId bson.ObjectId `json:"initiatorId"`
		TagId       bson.ObjectId `json:"tagId"`
	}
)
