package main

import (
	"github.com/gorilla/mux"
	"github.com/b1a9id/go-todo-mvc/src/controller"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	controller.InitDb()
	router.HandleFunc("/brands", controller.GetBrands).Methods("GET")
	router.HandleFunc("/brands/{id}", controller.GetBrand).Methods("GET")
	router.HandleFunc("/brands", controller.CreateBrand).Methods("POST")
	router.HandleFunc("/brands/{id}", controller.UpdateBrand).Methods("PUT")
	router.HandleFunc("/brands/{id}", controller.DeleteBrand).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
