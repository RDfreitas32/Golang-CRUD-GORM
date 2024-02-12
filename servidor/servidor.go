package servidor

import (
	"crud-gorm-one/banco"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Animais é a base de atributos de novos animais
type Animais struct {
	ID   uint32 `gorm:"primaryKey" json:"id"`
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Dono string `json:"dono"`
}

// PutAnimal, usada pra gerar um novo animal
func PutAnimal(w http.ResponseWriter, r *http.Request) {
	//deu erro ao rodar, vamos tentar mostrar a tabela que sera usada e verificar no BD
	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Error ro conect database", http.StatusBadRequest)
		return
	}
	defer banco.Fechar(db)

	var animal Animais

	erro = json.NewDecoder(r.Body).Decode(&animal)
	if erro != nil {
		http.Error(w, "Error to convert animal to struct", http.StatusBadRequest)
		return
	}

	insereAnimal := db.Create(&animal)
	if insereAnimal.Error != nil {
		http.Error(w, "Error to execut Put method: ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("Animal inserted into DB!: %d", animal.ID)
}

// Busca animal
func BuscaAnimal(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Error ro conect database", http.StatusBadRequest)
		return
	}
	//extrair o ID do parâmetro da URL
	paramsEx := mux.Vars(r)
	id, erro := strconv.ParseUint(paramsEx["id"], 10, 32)
	if erro != nil {
		http.Error(w, "invalid Animal ID", http.StatusBadRequest)
		return
	}
	//consultar bd para obter informações do animal
	var animal Animais
	result := db.First(&animal, id)
	if result.Error != nil {
		http.Error(w, "Animal not found", http.StatusNotFound)
		return
	}
	//escreve o detalhes do animal na resposta http
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}
