package entity

type DimVeiculo struct {
	ID            int    `db:"id_veiculo"`
	TipoVeiculo   string `db:"tipo_veiculo"`
	Marca         string `db:"marca"`
	AnoFabricacao int    `db:"ano_fabricacao"`
}
