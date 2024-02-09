package servidor

import (
	"crud-gorm-one/banco"
	"encoding/json"
	"net/http"
)

// Animais é a base de atributos de novos animais
type Animais struct {
	ID       uint32 `gorm:"primaryKey" json:"id"`
	Nome     string `json:"nome"`
	Raca     string `json:"raca"`
	DonoNome string `json:"dononome"`
}

func putAnimal(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Error ro conect database", http.StatusBadRequest)
	}
	defer db.Close()

	var animal Animais

	erro = json.NewDecoder(r.Body).Decode(&animal)
	if erro != nil {
		http.Error(w, "Error to convert animal to struct", http.StatusBadRequest)
		return
	}

}
