package db

import (
	"errors"
	"fmt"

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

func IsTagExist(s *mgo.Session, tagName string) (*Tag, error) {
	var (
		err error
	)

	tag := new(Tag)
	err = Collection(s, tag).Find(bson.M{"name": tagName}).One(tag)
	if err == nil {
		return tag, err
	}
	return nil, errors.New(fmt.Sprintf("Can not found tag with tagName. %v", tagName))
}

func AddTags(s *mgo.Session, feedId bson.ObjectId, tags []FeedTag) error {
	var (
		err  error
		feed *Feed
	)
	update := bson.M{"$addToSet": bson.M{"tags": bson.M{"$each": tags}}}
	_, err = Collection(s, feed).Upsert(bson.M{"_id": feedId}, update)
	return err
}

func RemoveTag(s *mgo.Session, feedId bson.ObjectId, tagId bson.ObjectId) error {
	var (
		err  error
		feed *Feed
	)
	update := bson.M{"$pull": bson.M{"tags": bson.M{"tagId": tagId}}}
	err = Collection(s, feed).Update(bson.M{"_id": feedId}, update)
	return err
}

func GetFeedsByTagId(s *mgo.Session, tagId bson.ObjectId, timestamp int64) ([]*Feed, error) {
	var (
		err   error
		feed  *Feed
		feeds []*Feed
	)
	feeds = make([]*Feed, 30)
	err = Collection(s, feed).Find(
		bson.M{
			"tags.tagId": tagId,
			"updateAt": bson.M{
				"$lt": timestamp,
			},
		}).Limit(30).Sort("-updateAt").All(&feeds)
	return feeds, err
}

func FuzzySearchByTagName(s *mgo.Session, tagName string) ([]*Tag, error) {
	var (
		err  error
		tag  *Tag
		tags []*Tag
	)
	tags = make([]*Tag, 30)
	err = Collection(s, tag).Find(
		bson.M{
			"name": bson.M{"$regex": bson.RegEx{tagName, "i"}},
		}).Limit(30).All(&tags)
	return tags, err
}
