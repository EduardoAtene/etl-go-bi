package repository

import (
	"database/sql"
	"fmt"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimVeiculoRepository struct {
	db *sql.DB
}

func NewDimVeiculoRepository(db *sql.DB) *DimVeiculoRepository {
	return &DimVeiculoRepository{db: db}
}

func (repo *DimVeiculoRepository) Insert(dimVeiculo *entity.DimVeiculo) (int, error) {
	// Query de inserção sem o RETURNING
	query := `INSERT INTO Dim_Veiculo (tipo_veiculo, marca, ano_fabricacao)
			  VALUES (?, ?, ?)`

	_, err := repo.db.Exec(query, dimVeiculo.TipoVeiculo, dimVeiculo.Marca, dimVeiculo.AnoFabricacao)
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
