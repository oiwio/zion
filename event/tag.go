package event

import (
	"zion/db"

	"gopkg.in/mgo.v2/bson"
)

type (
	TagEvent struct {
		EventId int           `json:"eventId,omitempty"`
		Tag     *db.Tag       `json:"tag,omitempty"`
		FeedId  bson.ObjectId `json:"feedId,omitempty"`
	}
)
