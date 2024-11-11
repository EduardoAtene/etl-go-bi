package repository

import (
	"database/sql"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimTempoRepository struct {
	db *sql.DB
}

func NewDimTempoRepository(db *sql.DB) *DimTempoRepository {
	return &DimTempoRepository{db: db}
}

func (repo *DimTempoRepository) Insert(dimTempo *entity.DimTempo) (int, error) {
	query := `INSERT INTO Dim_Tempo (data_completa, ano, mes, dia, dia_semana, hora, periodo_dia)
			  VALUES (?, ?, ?, ?, ?, ?, ?)
			  RETURNING id_tempo`

	var id int
	err := repo.db.QueryRow(query, dimTempo.DataCompleta, dimTempo.Ano, dimTempo.Mes, dimTempo.Dia,
		dimTempo.DiaSemana, dimTempo.Hora, dimTempo.PeriodoDia).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
