package routes

import (
	"github.com/COOPSPROFI/GoLangProject/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterEventStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/event/", controllers.GetEvents).Methods("GET")
	router.HandleFunc("/event/", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/event/{bookId}", controllers.GetEventById).Methods("GET")
	router.HandleFunc("/event/{bookId}", controllers.UpdateEvent).Methods("PUT")
	router.HandleFunc("/event/{bookId}", controllers.DeleteEvent).Methods("DELETE")
}
