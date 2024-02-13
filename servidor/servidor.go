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
func InsereAnimal(w http.ResponseWriter, r *http.Request) {
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

func AtualizaAnimal(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao abrir banco de dados: ", http.StatusInternalServerError)
	}

	// Extrair o ID do parâmetro da URL
	paramEx := mux.Vars(r)
	id, erro := strconv.ParseUint(paramEx["id"], 10, 32)
	if erro != nil {
		http.Error(w, "ID Inválido: ", http.StatusBadRequest)
		return
	}
	// Decodificar o JSON da solicitação para obter os novos dados do animal
	var animalAtualizado Animais

	erro = json.NewDecoder(r.Body).Decode(&animalAtualizado)
	if erro != nil {
		http.Error(w, "Erro ao ler o corpo da requisição: ", http.StatusBadRequest)
		return
	}

	// Consultar o banco de dados para obter o animal existente
	var animalCadastrado Animais
	resultado := db.First(&animalCadastrado, id)
	if resultado.Error != nil {
		http.Error(w, "Animal não encontrado", http.StatusNotFound)
		return
	}
	// Atualizar as informações do animal existente com os novos dados
	animalCadastrado.Nome = animalAtualizado.Nome
	animalCadastrado.Raca = animalAtualizado.Raca
	animalCadastrado.Dono = animalAtualizado.Dono

	//Salvar as alterações no banco de dados:
	resultado = db.Save(&animalCadastrado)
	if resultado.Error != nil {
		http.Error(w, "Erro ao atualizar: ", http.StatusInternalServerError)
		return
	}
	// Escrever uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Animal Atualizado com Sucesso"))
}
