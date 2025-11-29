package models

import (
	"time"
)

type Piloto struct {
	ID               uint               `gorm:"primarykey"`
	Nome             string             `gorm:"column:nome;size:100;not null"`
	Irating          int                `gorm:"column:irating"`
	CreatedAt        time.Time          `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time          `gorm:"column:updated_at;autoUpdateTime"`
	Carros           []Carro            `gorm:"many2many:piloto_carros_disponiveis;"`
	RestricaoHorario []RestricaoHorario `gorm:"many2many:piloto_restricao_horario;"`
}

type RestricaoHorario struct {
	ID         uint      `gorm:"primarykey"`
	HoraInicio time.Time `gorm:"column:hora_inicio;not null"`
	HoraFim    time.Time `gorm:"column:hora_fim;not null"`
}
