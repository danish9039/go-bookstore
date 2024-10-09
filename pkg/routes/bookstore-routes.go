package routes

import (
    "github.com/gorilla/mux"
	"github.com/danish9039/go-bookstore/pkg/controllers"

)

var RegisterBookstoreRoutes = func (router *mux.Router)  {

	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.Getbook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}






