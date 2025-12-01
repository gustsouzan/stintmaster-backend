package models

type Corrida struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	EventoID     uint           `gorm:"column:evento_id;not null"`
	Pilotos      []Piloto       `gorm:"foreignKey:CorridaID;references:ID"`
	StintPilotos []StintPilotos `gorm:"foreignKey:CorridaID;references:ID"`
}

type StintPilotos struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	CorridaID    uint    `gorm:"column:corrida_id;not null"`
	PilotoID     uint    `gorm:"column:piloto_id;not null"`
	HoraInicio   string  `gorm:"column:hora_inicio;not null"`
	HoraFim      string  `gorm:"column:hora_fim;not null"`
	Piloto       Piloto  `gorm:"foreignKey:PilotoID;references:ID"`
	ConsumoMedio float64 `gorm:"column:consumo_medio;not null"`
	ConsumoReal  float64 `gorm:"column:consumo_real;not null"`
	TempoVolta   float64 `gorm:"column:tempo_volta;not null"`
}
