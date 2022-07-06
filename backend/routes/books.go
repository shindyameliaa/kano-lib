package routes

import (
	"kanolibrary/backend/crud"
	"kanolibrary/backend/models"
	"kanolibrary/backend/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"encoding/json"
)

func FindAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "aplication/json")

	if r.Method == "GET" {
		result, err := crud.FindAll("books")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func FindByIdBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET"{
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		result, err := crud.Find("books", bson.M{"_id": objId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}


func InsertBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := crud.Insert("books", book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PATCH" {
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		util.ErrorChecker(err)
		result, err := crud.Update("books", bson.M{"_id": objId}, bson.M{"$set": bson.M{"title": book.Title, "author": book.Author, "synopsis": book.Synopsis, "release": book.Release}})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		result, err := crud.Remove("books", bson.M{"_id":objId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}