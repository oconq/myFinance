package controllers

import (
	"net/http"
	"encoding/json"
	"awesomeProject/models"
	"github.com/gorilla/mux"
	"strconv"
	"log"
	"fmt"
)

func AccountsIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entering AccountsIndex...")
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllAccounts())
}

func AccountsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	account := models.FindAccountById(id)

	json.NewEncoder(w).Encode(account)
}
