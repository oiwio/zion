package event

import "gopkg.in/mgo.v2/bson"

type (
	FriendEvent struct {
		EventId     int           `json:"eventId"`
		InitiatorId bson.ObjectId `json:"initiatorId"`
		ResponderId bson.ObjectId `json:"responderId"`
	}
)
