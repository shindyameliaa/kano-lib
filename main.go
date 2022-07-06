package main


import (
	"fmt"
	"kanolibrary/backend/routes"
	"net/http"
// 	"kanolibrary/backend/crud"
 )

type Books struct {
	Title string `bson:"title"`
	Author string `bson:"author"`
	Synopsis string `bson:"synopsis"`
	Release string `bson:"release"`
}

func main() {
	// crud.Insert("books", Books{Title: "Hujan", Author: "Tere Liye", Synopsis: "Hujan merupakan . .", Release: "2016"})
	// crud.Find("books", Books{})
	// crud.Update()
	// crud.Remove()

	http.HandleFunc("/books", routes.FindAll)
	http.HandleFunc("/books/find", routes.FindByIdBook)
	http.HandleFunc("/books/insert", routes.InsertBook)
	http.HandleFunc("/books/update", routes.UpdateBook)
	http.HandleFunc("/books/delete", routes.DeleteBook)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}