package repository

import (
	"database/sql"
	"fmt"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimTempoRepository struct {
	db *sql.DB
}

func NewDimTempoRepository(db *sql.DB) *DimTempoRepository {
	return &DimTempoRepository{db: db}
}

func (repo *DimTempoRepository) Insert(dimTempo *entity.DimTempo) (int, error) {
	// Converte a hora para o formato "hh:mm:ss"
	horaFormatada := dimTempo.Hora.Format("15:04:05") // Formato de hora

	query := `INSERT INTO Dim_Tempo (data_completa, ano, mes, dia, dia_semana, hora, periodo_dia)
			  VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := repo.db.Exec(query, dimTempo.DataCompleta, dimTempo.Ano, dimTempo.Mes, dimTempo.Dia,
		dimTempo.DiaSemana, horaFormatada, dimTempo.PeriodoDia)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Obtém o último id inserido
	var id int
	err = repo.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
