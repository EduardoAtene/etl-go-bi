package repository

import (
	"database/sql"
	"fmt"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type DimPessoaRepository struct {
	db *sql.DB
}

func NewDimPessoaRepository(db *sql.DB) *DimPessoaRepository {
	return &DimPessoaRepository{db: db}
}

func (repo *DimPessoaRepository) Insert(dimPessoa *entity.DimPessoa) (int, error) {
	// Query de inserção sem o RETURNING
	query := `INSERT INTO Dim_Pessoa (tipo_envolvido, idade, sexo, raca_cor, estado_fisico, municipio_residencia)
			  VALUES (?, ?, ?, ?, ?, ?)`

	_, err := repo.db.Exec(query, dimPessoa.TipoEnvolvido, dimPessoa.Idade, dimPessoa.Sexo,
		dimPessoa.RacaCor, dimPessoa.EstadoFisico, dimPessoa.MunicipioResidencia)
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
