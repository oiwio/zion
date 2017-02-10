package db

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Friendship struct {
		FriendshipId    bson.ObjectId `json:"friendshipId" bson:"_id"`
		InitiatorId     bson.ObjectId `json:"initiatorId" bson:"initiatorId"`
		InitiatorName   string        `json:"initiatorName" bson:"initiatorName"`
		InitiatorAvatar string        `json:"initiatorAvatar" bson:"initiatorAvatar"`
		ResponderId     bson.ObjectId `json:"responderId" bson:"responderId"`
		ResponderName   string        `json:"responderName" bson:"responderName"`
		ResponderAvatar string        `json:"responderAvatar" bson:"responderAvatar"`
		InitiateAt      int64         `json:"initiateAt" bson:"initiateAt"`
	}
)

func FollowUser(s *mgo.Session, initiatorId bson.ObjectId, responderId bson.ObjectId) (*Friendship, error) {
	var (
		err        error
		friendship *Friendship
		initiator  *User
		responder  *User
	)

	friendship = new(Friendship)
	friendship.InitiatorId = initiatorId
	friendship.ResponderId = responderId

	initiator, err = GetUserById(s, initiatorId)
	responder, err = GetUserById(s, responderId)
	if err != nil {
		return nil, err
	}

	friendship.FriendshipId = bson.NewObjectId()
	friendship.InitiateAt = time.Now().Unix()
	friendship.InitiatorName = initiator.Nickname
	friendship.InitiatorAvatar = initiator.Avatar
	friendship.ResponderName = responder.Nickname
	friendship.ResponderAvatar = responder.Avatar

	collection := Collection(s, friendship)
	err = collection.Insert(friendship)
	if err != nil {
		return nil, err
	}

	return friendship, err
}

func UnfollowUser(s *mgo.Session, initiatorId bson.ObjectId, responderId bson.ObjectId) error {
	var (
		err        error
		friendship *Friendship
	)

	err = Collection(s, friendship).Remove(
		bson.M{
			"initiatorId": initiatorId,
			"responderId": responderId,
		})
	return err
}

// 获取两人之间的关系，0无任何关系，1被对方关注了，2关注了对方，3互相关注, -1意外情况
func GetRelation(s *mgo.Session, initiatorId bson.ObjectId, responderId bson.ObjectId) int {
	var (
		following bool
		followed  bool
	)

	following = IsFriendshipExist(s, initiatorId, responderId)
	followed = IsFriendshipExist(s, responderId, initiatorId)

	if following && followed {
		return 3
	} else if following && !followed {
		return 2
	} else if !following && followed {
		return 1
	} else if !following && !followed {
		return 0
	}
	return -1
}

func IsFriendshipExist(s *mgo.Session, initiatorId bson.ObjectId, responderId bson.ObjectId) bool {
	var (
		err        error
		friendship *Friendship
	)
	friendship = new(Friendship)
	err = Collection(s, friendship).Find(
		bson.M{
			"initiatorId": initiatorId,
			"responderId": responderId,
		}).One(friendship)
	if err != nil {
		return false
	}
	return true
}
