package entity

type DimCondicoes struct {
	ID                    int    `db:"id_condicoes"`
	CondicaoMeteorologica string `db:"condicao_meteorologica"`
	TipoPista             string `db:"tipo_pista"`
	TracadoVia            string `db:"tracado_via"`
	UsoSolo               string `db:"uso_solo"`
	SentidoVia            string `db:"sentido_via"`
}
