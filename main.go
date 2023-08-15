package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"luhn/algoritm"
)

func main() {
	newRouter := mux.NewRouter()
	newRouter.HandleFunc("/", algoritm.GetData).Methods("POST")
	logrus.Fatal(http.ListenAndServe(":10000", newRouter))
}
