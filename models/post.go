package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Caption   string        `bson:"caption" json:"caption"`
	Image_URL string        `bson:"image_url" json:"image_url"`
	Timestamp time.Time     `json:"timestamp"`
}
