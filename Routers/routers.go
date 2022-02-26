package Routers

import (
	"github.com/gorilla/mux"
	"github.com/jay-salunke/TimeAPI/Handlers"
)

func Routers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", Handlers.GetTime).Methods("GET")
	return router
}
