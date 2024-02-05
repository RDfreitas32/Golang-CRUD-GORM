package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar retorna um ponteiro de sql
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	//verificar o status da requisição de abertura do bd
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
