package routes

import (
	"github.com/gorilla/mux"
	"github.com/mehulrana017/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(routes *mux.Router) {
	routes.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	routes.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	routes.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	routes.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
}
