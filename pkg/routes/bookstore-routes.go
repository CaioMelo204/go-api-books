package routes

import (
	"github.com/gorilla/mux"
	"go-api-1/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book", controllers.PostBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
