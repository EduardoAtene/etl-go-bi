package entity

type FatoAcidentes struct {
	IDAcidente            int    `db:"id_acidente"`
	IDTempo               int    `db:"id_tempo"`
	IDLocalizacao         int    `db:"id_localizacao"`
	IDVeiculo             int    `db:"id_veiculo"`
	IDPessoa              int    `db:"id_pessoa"`
	IDCondicoes           int    `db:"id_condicoes"`
	FonteDados            string `db:"fonte_dados"`
	CausaAcidente         string `db:"causa_acidente"`
	TipoAcidente          string `db:"tipo_acidente"`
	ClassificacaoAcidente string `db:"classificacao_acidente"`
	CIDCausaMorte         string `db:"cid_causa_morte"`
	DescCausaMorte        string `db:"desc_causa_morte"`
	QtdIlesos             int    `db:"qtd_ilesos"`
	QtdFeridosLeves       int    `db:"qtd_feridos_leves"`
	QtdFeridosGraves      int    `db:"qtd_feridos_graves"`
	QtdMortos             int    `db:"qtd_mortos"`
}
