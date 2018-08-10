package myFinance

import (
	"github.com/gorilla/mux"
	"awesomeProject/controllers"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/accounts").Name("Index").HandlerFunc(controllers.AccountsIndex)
	router.Methods("GET").Path("/accounts/{id}").Name("Show").HandlerFunc(controllers.AccountsShow)

	return router
}
