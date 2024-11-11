package repository

import (
	"database/sql"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimLocalizacaoRepository struct {
	db *sql.DB
}

func NewDimLocalizacaoRepository(db *sql.DB) *DimLocalizacaoRepository {
	return &DimLocalizacaoRepository{db: db}
}

func (repo *DimLocalizacaoRepository) Insert(dimLocalizacao *entity.DimLocalizacao) (int, error) {
	query := `INSERT INTO Dim_Localizacao (municipio, br, km, latitude, longitude, regional, delegacia, uop)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)
			  RETURNING id_localizacao`

	var id int
	err := repo.db.QueryRow(query, dimLocalizacao.Municipio, dimLocalizacao.BR, dimLocalizacao.KM,
		dimLocalizacao.Latitude, dimLocalizacao.Longitude, dimLocalizacao.Regional,
		dimLocalizacao.Delegacia, dimLocalizacao.UOP).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
