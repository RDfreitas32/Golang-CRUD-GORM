package main

import (
	"crud-gorm-one/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/animais", servidor.InsereAnimal).Methods(http.MethodPost)
	router.HandleFunc("/animal/{id}", servidor.BuscaAnimal).Methods(http.MethodGet)
	router.HandleFunc("/animais/{id}", servidor.AtualizaAnimal).Methods(http.MethodPut)
	router.HandleFunc("/animais/{id}", servidor.ApagaAnimal).Methods(http.MethodDelete)
	router.HandleFunc("/animais", servidor.BuscaAnimais).Methods(http.MethodGet)

	fmt.Println("Service On-line")
	log.Fatal(http.ListenAndServe(":8080", router))
}
