package models

import "gorm.io/gorm"

type Pista struct {
	gorm.Model
	Nome        string `gorm:"column:nome;size:100;not null"`
	Localizacao string `gorm:"column:localizacao;size:100;not null"`
	Imagem      string `gorm:"column:imagem;size:255"`
}
