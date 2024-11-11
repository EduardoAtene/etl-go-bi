package repository

import (
	"database/sql"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
)

type FatoAcidentesRepository struct {
	db *sql.DB
}

func NewFatoAcidentesRepository(db *sql.DB) *FatoAcidentesRepository {
	return &FatoAcidentesRepository{db: db}
}

func (repo *FatoAcidentesRepository) Insert(fato *entity.FatoAcidentes) (int, error) {
	query := `INSERT INTO Fato_Acidentes (
		id_tempo, id_localizacao, id_veiculo, id_pessoa, id_condicoes, fonte_dados,
		causa_acidente, tipo_acidente, classificacao_acidente, cid_causa_morte,
		desc_causa_morte, qtd_ilesos, qtd_feridos_leves, qtd_feridos_graves, qtd_mortos
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id_acidente`

	var id int
	err := repo.db.QueryRow(query, fato.IDTempo, fato.IDLocalizacao, fato.IDVeiculo, fato.IDPessoa,
		fato.IDCondicoes, fato.FonteDados, fato.CausaAcidente, fato.TipoAcidente,
		fato.ClassificacaoAcidente, fato.CIDCausaMorte, fato.DescCausaMorte,
		fato.QtdIlesos, fato.QtdFeridosLeves, fato.QtdFeridosGraves, fato.QtdMortos).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
