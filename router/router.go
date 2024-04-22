package router

import (
	"LoginAPI/helper"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	var router = mux.NewRouter()
	router.HandleFunc("/user/loginAPI", helper.LoginAPI).Methods("POST")
	return router
}
