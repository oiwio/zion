package db

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Tag struct {
		TagId       bson.ObjectId `json:"tagId" bson:"_id"`
		Name        string        `json:"name" bson:"name"`
		Cover       string        `json:"cover,omitempty" structs:"cover,omitempty"`
		Description string        `json:"desc,omitempty" structs:"desc,omitempty"`
		CreateUser  bson.ObjectId `json:"createUser,omitempty" bson:"createUser,omitempty"`
		CreateAt    int64         `json:"createAt,omitempty" bson:"createAt,omitempty"`
		UpdateAt    int64         `json:"updateAt,omitempty" bson:"updateAt,omitempty"`
	}
)

func NewTag(s *mgo.Session, tag *Tag) (*Tag, error) {
	var (
		err error
	)

	collection := Collection(s, tag)
	err = collection.Insert(tag)
	if err != nil {
		return nil, err
	}
	return tag, err
}

func GetTagById(s *mgo.Session, tagId bson.ObjectId) (*Tag, error) {
	var (
		err error
	)
	tag := new(Tag)
	err = Collection(s, tag).FindId(tagId).One(tag)
	return tag, err
}
