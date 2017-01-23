package event

import (
	"zion/db"

	"gopkg.in/mgo.v2/bson"
)

type (
	FeedEvent struct {
		EventId   int           `json:"eventId,omitempty"`
		Comment   *db.Comment   `json:"comment,omitempty"`
		Feed      *db.Feed      `json:"feed,omitempty"`
		FeedId    bson.ObjectId `json:"feedId,omitempty"`
		UserId    bson.ObjectId `json:"userId,omitempty"`
		CommentId bson.ObjectId `json:"commentId,omitempty"`
	}
)
