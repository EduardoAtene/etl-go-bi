package repository

import (
	"database/sql"
	"fmt"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimCondicoesRepository struct {
	db *sql.DB
}

func NewDimCondicoesRepository(db *sql.DB) *DimCondicoesRepository {
	return &DimCondicoesRepository{db: db}
}

func (repo *DimCondicoesRepository) Insert(dimCondicoes *entity.DimCondicoes) (int, error) {
	// Query de inserção sem o RETURNING
	query := `INSERT INTO Dim_Condicoes (condicao_meteorologica, tipo_pista, tracado_via, uso_solo, sentido_via)
			  VALUES (?, ?, ?, ?, ?)`

	_, err := repo.db.Exec(query, dimCondicoes.CondicaoMeteorologica, dimCondicoes.TipoPista,
		dimCondicoes.TracadoVia, dimCondicoes.UsoSolo, dimCondicoes.SentidoVia)
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
