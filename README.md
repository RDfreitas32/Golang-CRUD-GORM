# Projeto CRUD de Animais com ORM Golang GORM

Este é um projeto de CRUD (Create, Read, Update, Delete) para gerenciar informações de animais.

## Estrutura de Dados

O projeto usa a seguinte estrutura para representar os animais:

```go
type Animais struct {
	ID   uint32 `gorm:"primaryKey" json:"id"`
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Dono string `json:"dono"`
}
