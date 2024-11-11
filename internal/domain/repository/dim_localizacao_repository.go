package repository

import (
	"database/sql"
	"fmt"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimLocalizacaoRepository struct {
	db *sql.DB
}

func NewDimLocalizacaoRepository(db *sql.DB) *DimLocalizacaoRepository {
	return &DimLocalizacaoRepository{db: db}
}

func (repo *DimLocalizacaoRepository) Insert(dimLocalizacao *entity.DimLocalizacao) (int, error) {
	// Query de inserção sem o RETURNING
	query := `INSERT INTO Dim_Localizacao (municipio, br, km, latitude, longitude, regional, delegacia, uop)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := repo.db.Exec(query, dimLocalizacao.Municipio, dimLocalizacao.BR, dimLocalizacao.KM,
		dimLocalizacao.Latitude, dimLocalizacao.Longitude, dimLocalizacao.Regional,
		dimLocalizacao.Delegacia, dimLocalizacao.UOP)
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
