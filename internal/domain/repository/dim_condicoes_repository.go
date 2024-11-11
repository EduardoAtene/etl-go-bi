package repository

import (
	"database/sql"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimCondicoesRepository struct {
	db *sql.DB
}

func NewDimCondicoesRepository(db *sql.DB) *DimCondicoesRepository {
	return &DimCondicoesRepository{db: db}
}

func (repo *DimCondicoesRepository) Insert(dimCondicoes *entity.DimCondicoes) (int, error) {
	query := `INSERT INTO Dim_Condicoes (condicao_meteorologica, tipo_pista, tracado_via, uso_solo, sentido_via)
			  VALUES (?, ?, ?, ?, ?)
			  RETURNING id_condicoes`

	var id int
	err := repo.db.QueryRow(query, dimCondicoes.CondicaoMeteorologica, dimCondicoes.TipoPista,
		dimCondicoes.TracadoVia, dimCondicoes.UsoSolo, dimCondicoes.SentidoVia).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
