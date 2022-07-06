package crud

import (
	"context"
    "fmt"
    "kanolibrary/backend/mongo"
	"kanolibrary/backend/util"
)

func Update(collection string, query map[string]interface{}, update map [string]interface{}) (string, error) {
	ctx := context.Background()
    db, err := mongo.Connect()
    util.ErrorChecker(err)

    fmt.Println(update)
    _, err = db.Collection("books").UpdateOne(ctx, query, update)
    util.ErrorChecker(err)

    return"Update success!", err
}