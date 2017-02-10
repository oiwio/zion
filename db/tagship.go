package db

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Tagship struct {
		TagshipId       bson.ObjectId `json:"tagshipId" bson:"_id"`
		InitiatorId     bson.ObjectId `json:"initiatorId" bson:"initiatorId"`
		InitiatorName   string        `json:"initiatorName,omitempty" bson:"initiatorName,omitempty"`
		InitiatorAvatar string        `json:"initiatorAvatar,omitempty" bson:"initiatorAvatar,omitempty"`
		TagId           bson.ObjectId `json:"tagId" bson:"tagId"`
		TagName         string        `json:"tagName,omitempty" bson:"tagName,omitempty"`
		TagCover        string        `json:"tagcover,omitempty" bson:"tagcover,omitempty"`
		InitiateAt      int64         `json:"initiateAt" bson:"initiateAt"`
	}
)

func FollowTag(s *mgo.Session, initiatorId bson.ObjectId, tagId bson.ObjectId) (*Tagship, error) {
	var (
		err       error
		tagship   *Tagship
		initiator *User
		tag       *Tag
	)

	tagship = new(Tagship)
	tagship.InitiatorId = initiatorId
	tagship.TagId = tagId

	initiator, err = GetUserById(s, initiatorId)
	tag, err = GetTagById(s, tagId)
	if err != nil {
		return nil, err
	}

	tagship.TagshipId = bson.NewObjectId()
	tagship.InitiateAt = time.Now().Unix()
	tagship.InitiatorName = initiator.Nickname
	tagship.InitiatorAvatar = initiator.Avatar
	tagship.TagName = tag.Name
	tagship.TagCover = tag.Cover

	collection := Collection(s, tagship)
	err = collection.Insert(tagship)
	if err != nil {
		return nil, err
	}

	return tagship, err
}

func UnfollowTag(s *mgo.Session, initiatorId bson.ObjectId, tagId bson.ObjectId) error {
	var (
		err     error
		tagship *Tagship
	)

	err = Collection(s, tagship).Remove(
		bson.M{
			"initiatorId": initiatorId,
			"tagId":       tagId,
		})
	return err
}

func GetTagshipsByTagId(s *mgo.Session, tagId bson.ObjectId) ([]*Tagship, error) {
	var (
		err      error
		tagship  *Tagship
		tagships []*Tagship
	)
	tagships = make([]*Tagship, 15)
	err = Collection(s, tagship).Find(
		bson.M{
			"tagId": tagId,
		}).Sort("-initiateAt").All(&tagships)
	return tagships, err
}

func GetTagshipsByInitiatorId(s *mgo.Session, initiatorId bson.ObjectId) ([]*Tagship, error) {
	var (
		err      error
		tagship  *Tagship
		tagships []*Tagship
	)
	tagships = make([]*Tagship, 15)
	err = Collection(s, tagship).Find(
		bson.M{
			"initiatorId": initiatorId,
		}).Sort("-initiateAt").All(&tagships)
	return tagships, err
}
