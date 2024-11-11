package repository

import (
	"database/sql"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimPessoaRepository struct {
	db *sql.DB
}

func NewDimPessoaRepository(db *sql.DB) *DimPessoaRepository {
	return &DimPessoaRepository{db: db}
}

func (repo *DimPessoaRepository) Insert(dimPessoa *entity.DimPessoa) (int, error) {
	query := `INSERT INTO Dim_Pessoa (tipo_envolvido, idade, sexo, raca_cor, estado_fisico, municipio_residencia)
			  VALUES (?, ?, ?, ?, ?, ?)
			  RETURNING id_pessoa`

	var id int
	err := repo.db.QueryRow(query, dimPessoa.TipoEnvolvido, dimPessoa.Idade, dimPessoa.Sexo,
		dimPessoa.RacaCor, dimPessoa.EstadoFisico, dimPessoa.MunicipioResidencia).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
