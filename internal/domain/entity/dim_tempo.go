package entity

import "time"

type DimTempo struct {
	ID           int       `db:"id_tempo"`
	DataCompleta time.Time `db:"data_completa"`
	Ano          int       `db:"ano"`
	Mes          int       `db:"mes"`
	Dia          int       `db:"dia"`
	DiaSemana    string    `db:"dia_semana"`
	Hora         time.Time `db:"hora"`
	PeriodoDia   string    `db:"periodo_dia"`
}
