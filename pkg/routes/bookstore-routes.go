package routes

import (
	"github.com/gorilla/mux"
	"github.com/sunday4me/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Method("POST")
	router.HandleFunc("/book/", controllers.GetBook).Method("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Method("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Method("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Method("DELETE")
}
