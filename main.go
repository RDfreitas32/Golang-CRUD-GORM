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
	router.HandleFunc("/animais/{id}", servidor.BuscaAnimal).Methods(http.MethodGet)
	router.HandleFunc("/animais/{id}", servidor.AtualizaAnimal).Methods(http.MethodPut)
	fmt.Println("Service On-line")
	log.Fatal(http.ListenAndServe(":8080", router))
}
