package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Evento struct {
	gorm.Model
	Nome       string         `gorm:"column:nome;size:100;not null"`
	Plataforma string         `gorm:"column:plataforma;size:50;not null"`
	DataEvento time.Time      `gorm:"column:data_evento;not null"`
	Duracao    int            `gorm:"column:duracao;not null"` // duração em minutos
	Classes    pq.StringArray `gorm:"column:classes;type:text[];not null"`
	MinPilotos int            `gorm:"column:min_pilotos;not null"`
	MaxPilotos int            `gorm:"column:max_pilotos;not null"`
	Imagem     string         `gorm:"column:imagem;type:text"`
	CreatedBy  string         `gorm:"column:created_by;size:100;not null"`
	PistaId    uint           `gorm:"column:pista_id;not null"`
	Pista      Pista          `gorm:"foreignKey:PistaId;references:ID"`
}
