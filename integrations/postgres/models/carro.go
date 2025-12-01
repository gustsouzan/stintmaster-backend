package models

import "time"

type Carro struct {
	ID        uint      `gorm:"primarykey"`
	Nome      string    `gorm:"column:nome;size:100;not null"`
	Classe    string    `gorm:"column:classe;size:50;not null"`
	Imagem    string    `gorm:"column:imagem;size:255"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}
