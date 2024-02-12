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

	router.HandleFunc("/animais", servidor.PutAnimal).Methods(http.MethodPost)
	router.HandleFunc("/animais/{id}", servidor.BuscaAnimal).Methods((http.MethodGet))

	fmt.Println("Service On-line")
	log.Fatal(http.ListenAndServe(":8080", router))
}
