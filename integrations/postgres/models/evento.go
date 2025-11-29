package models

import (
	"github.com/lib/pq"
)

type Evento struct {
	ID         uint           `gorm:"primarykey"`
	Duracao    int            `gorm:"column:duracao;not null"` // duração em minutos
	Classes    pq.StringArray `gorm:"column:classes;type:text[];not null"`
	MinPilotos int            `gorm:"column:min_pilotos;not null"`
	MaxPilotos int            `gorm:"column:max_pilotos;not null"`
	PistaId    uint           `gorm:"column:pista_id;not null"`
	Pista      Pista          `gorm:"foreignKey:PistaId;references:ID"`
}

type EventoPilotosInscritos struct {
	EventoID uint   `gorm:"column:evento_id;not null;uniqueIndex:idx_evento_piloto"`
	PilotoID uint   `gorm:"column:piloto_id;not null;uniqueIndex:idx_evento_piloto"`
	Evento   Evento `gorm:"foreignKey:EventoID;references:ID"`
	Piloto   Piloto `gorm:"foreignKey:PilotoID;references:ID"`
}
