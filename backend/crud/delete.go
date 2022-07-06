package crud

import (
	"context"


    "kanolibrary/backend/mongo"
	"kanolibrary/backend/util"
)

func Remove(collection string, query map[string]interface{}) (string, error) {
	ctx := context.Background()
    db, err := mongo.Connect()
    util.ErrorChecker(err)

    
    _, err = db.Collection(collection).DeleteOne(ctx, query)
    util.ErrorChecker(err)

    return"Remove success!", err
}