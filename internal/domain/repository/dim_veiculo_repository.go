package repository

import (
	"database/sql"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimVeiculoRepository struct {
	db *sql.DB
}

func NewDimVeiculoRepository(db *sql.DB) *DimVeiculoRepository {
	return &DimVeiculoRepository{db: db}
}

func (repo *DimVeiculoRepository) Insert(dimVeiculo *entity.DimVeiculo) (int, error) {
	query := `INSERT INTO Dim_Veiculo (tipo_veiculo, marca, ano_fabricacao)
			  VALUES (?, ?, ?)
			  RETURNING id_veiculo`

	var id int
	err := repo.db.QueryRow(query, dimVeiculo.TipoVeiculo, dimVeiculo.Marca, dimVeiculo.AnoFabricacao).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
