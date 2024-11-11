package entity

type DimPessoa struct {
	ID                  int    `db:"id_pessoa"`
	TipoEnvolvido       string `db:"tipo_envolvido"`
	Idade               int    `db:"idade"`
	Sexo                string `db:"sexo"`
	RacaCor             string `db:"raca_cor"`
	EstadoFisico        string `db:"estado_fisico"`
	MunicipioResidencia string `db:"municipio_residencia"`
}
