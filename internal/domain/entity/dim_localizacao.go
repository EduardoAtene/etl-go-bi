package entity

type DimLocalizacao struct {
	ID        int     `db:"id_localizacao"`
	Municipio string  `db:"municipio"`
	BR        string  `db:"br"`
	KM        int     `db:"km"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
	Regional  string  `db:"regional"`
	Delegacia string  `db:"delegacia"`
	UOP       string  `db:"uop"`
}
