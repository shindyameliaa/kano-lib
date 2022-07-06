package crud

import (
	"context"
    "kanolibrary/backend/mongo"
	"kanolibrary/backend/util"
	"kanolibrary/backend/models"
)

func Insert(collection string, query models.Books) (string,error) {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	_, err = db.Collection(collection).InsertOne(ctx, query)
	util.ErrorChecker(err)
	return "Inserted Successfully", err
}