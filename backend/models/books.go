package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Books struct {
	_id			primitive.ObjectID `bson: "_id"`
	Title		string `bson: "title" `
	Author 		string `bson: "author" `
	Synopsis	string `bson: "synopsis"`
	Release	 	string `bson: "release" `
}