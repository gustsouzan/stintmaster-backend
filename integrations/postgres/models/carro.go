package models

type Carro struct {
	ID     uint   `gorm:"primarykey"`
	Nome   string `gorm:"column:nome;size:100;not null"`
	Classe string `gorm:"column:classe;size:50;not null"`
	Imagem string `gorm:"column:imagem;size:255"`
}
