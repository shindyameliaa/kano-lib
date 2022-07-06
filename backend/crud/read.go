package crud

import (
	"context"
	"encoding/json"
	"fmt"
	"kanolibrary/backend/models"
	"kanolibrary/backend/mongo"
	"kanolibrary/backend/util"

	"go.mongodb.org/mongo-driver/bson"
)

type books struct {
	Title, Author, Synopsis, Release string
}

func Find(collection string, query map[string]interface{}) ([]byte, error) {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	csr, err := db.Collection(collection).Find(ctx, query)
	util.ErrorChecker(err)
	defer csr.Close(ctx)

	result := make([]books, 0)
	for csr.Next(ctx) {
		var row books
		err := csr.Decode(&row)
		util.ErrorChecker(err)

		result = append(result, row)
	}

	data, err := json.Marshal(result)
	util.ErrorChecker(err)
	return data, err

	var i = 0
	for {
		fmt.Println("Title      :", result[i].Title)
		fmt.Println("Author     :", result[i].Author)
		fmt.Println("Synopsis   :", result[i].Synopsis)
		fmt.Println("Release    :", result[i].Release)
		fmt.Println("")
		i++
	}
}

func FindAll(collection string) ([]byte, error) {

	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	csr, err := db.Collection(collection).Find(ctx, bson.M{})
	util.ErrorChecker(err)
	defer csr.Close(ctx)

	result := make([]models.Books, 0)
	for csr.Next(ctx) {
		var row models.Books
		err := csr.Decode(&row)
		util.ErrorChecker(err)

		result = append(result, row)
	}

	// if len(result) > 0 {
	//     fmt.Println("Name   :", result[0].Name)
	//     fmt.Println("Author :", result[0].Author)
	// }

	data, err := json.Marshal(result)
	util.ErrorChecker(err)
	return data, err
}
