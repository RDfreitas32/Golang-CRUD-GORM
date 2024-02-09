package banco

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Conectar retorna um ponteiro de sql
func Conectar() (*gorm.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := gorm.Open(mysql.Open(stringConexao), &gorm.Config{})
	if erro != nil {
		return nil, erro
	}
	return db, nil
}

// Fechar a conexão com bd
func Fechar(db *gorm.DB) error {
	sqlDB, erro := db.DB()
	if erro != nil {
		return erro
	}

	return sqlDB.Close()
}
