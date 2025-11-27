package models

import (
	"time"

	"gorm.io/gorm"
)

type Piloto struct {
	gorm.Model
	Nome      string    `gorm:"column:nome;size:100;not null"`
	Email     string    `gorm:"column:email;size:100;not null"`
	IracingID string    `gorm:"column:iracing_id"`
	Irating   int       `gorm:"column:irating"`
	Imagem    string    `gorm:"column:imagem;size:255"`
	CreatedBy string    `gorm:"column:created_by;size:100;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type PilotoCarrosDisponiveis struct {
	PilotoID uint   `gorm:"column:piloto_id;not null;uniqueIndex:idx_piloto_carro"`
	CarroID  uint   `gorm:"column:carro_id;not null;uniqueIndex:idx_piloto_carro"`
	Piloto   Piloto `gorm:"foreignKey:PilotoID;references:ID"`
	Carro    Carro  `gorm:"foreignKey:CarroID;references:ID"`
}
