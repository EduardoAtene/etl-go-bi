package repository

import (
	"database/sql"
	"fmt"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type FatoAcidentesRepository struct {
	db *sql.DB
}

func NewFatoAcidentesRepository(db *sql.DB) *FatoAcidentesRepository {
	return &FatoAcidentesRepository{db: db}
}

// Método para validar se a chave estrangeira é válida ou deve ser nula
func (repo *FatoAcidentesRepository) getNullableValue(value int) interface{} {
	if value == 0 {
		return nil
	}
	return value
}

// Método para realizar a inserção na tabela Fato_Acidentes
func (repo *FatoAcidentesRepository) insertFatoAcidente(fato *entity.FatoAcidentes) (int, error) {
	// Prepara a query de inserção
	query := `INSERT INTO Fato_Acidentes (
		id_tempo, id_localizacao, id_veiculo, id_pessoa, id_condicoes, fonte_dados,
		causa_acidente, tipo_acidente, classificacao_acidente, cid_causa_morte,
		desc_causa_morte, qtd_ilesos, qtd_feridos_leves, qtd_feridos_graves, qtd_mortos
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Definir valores das chaves estrangeiras, tratando como NULL quando necessário
	idVeiculo := repo.getNullableValue(fato.IDVeiculo)
	idPessoa := repo.getNullableValue(fato.IDPessoa)
	idCondicoes := repo.getNullableValue(fato.IDCondicoes)
	idLocalizacao := repo.getNullableValue(fato.IDLocalizacao)
	idTempo := repo.getNullableValue(fato.IDTempo)

	// Inserir os dados na tabela Fato_Acidentes
	result, err := repo.db.Exec(query, idTempo, idLocalizacao, idVeiculo, idPessoa,
		idCondicoes, fato.FonteDados, fato.CausaAcidente, fato.TipoAcidente,
		fato.ClassificacaoAcidente, fato.CIDCausaMorte, fato.DescCausaMorte,
		fato.QtdIlesos, fato.QtdFeridosLeves, fato.QtdFeridosGraves, fato.QtdMortos)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Recuperar o ID do registro inserido
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Método para inserir e retornar o id do acidente
func (repo *FatoAcidentesRepository) Insert(fato *entity.FatoAcidentes) (int, error) {
	// Inserir o fato do acidente e retornar o id
	id, err := repo.insertFatoAcidente(fato)
	if err != nil {
		return 0, err
	}
	return id, nil
}
