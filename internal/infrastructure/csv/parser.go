package csv

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

// Estruturas auxiliares para organizar os dados antes de criar o fato
type AcidenteTemp struct {
	// Dados temporais
	DataCompleta time.Time
	Hora         string
	PeriodoDia   string

	// Dados de localização
	Municipio string
	BR        string
	KM        int
	Latitude  float64
	Longitude float64
	Regional  string
	Delegacia string
	UOP       string

	// Dados do veículo
	TipoVeiculo   string
	Marca         string
	AnoFabricacao int

	// Dados da pessoa
	TipoEnvolvido string
	Idade         int
	Sexo          string
	RacaCor       string
	EstadoFisico  string

	// Dados das condições
	CondicaoMetereologica string
	TipoPista             string
	TracadoVia            string
	UsoSolo               string
	SentidoVia            string

	// Dados do acidente
	CausaAcidente         string
	TipoAcidente          string
	ClassificacaoAcidente string
	CIDCausaMorte         string
	DescCausaMorte        string
	QtdIlesos             int
	QtdFeridosLeves       int
	QtdFeridosGraves      int
	QtdMortos             int
}

// ParseAcidentesPRF parse acidentes da PRF
func ParseAcidentesPRF(file io.Reader) ([]AcidenteTemp, error) {
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Pular cabeçalho
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var acidentes []AcidenteTemp

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		data, err := time.Parse("02/01/2006", record[0])
		if err != nil {
			continue
		}

		km, _ := strconv.Atoi(record[4])
		idade, _ := strconv.Atoi(record[15])
		anoFab, _ := strconv.Atoi(record[12])
		ilesos, _ := strconv.Atoi(record[9])
		feridosLeves, _ := strconv.Atoi(record[7])
		feridosGraves, _ := strconv.Atoi(record[8])
		mortos, _ := strconv.Atoi(record[6])
		latitude, _ := strconv.ParseFloat(record[13], 64)
		longitude, _ := strconv.ParseFloat(record[14], 64)

		acidente := AcidenteTemp{
			DataCompleta:          data,
			Hora:                  record[1],
			PeriodoDia:            record[2],
			Municipio:             record[3],
			BR:                    record[4],
			KM:                    km,
			Latitude:              latitude,
			Longitude:             longitude,
			Regional:              record[16],
			Delegacia:             record[17],
			UOP:                   record[18],
			TipoVeiculo:           record[10],
			Marca:                 record[11],
			AnoFabricacao:         anoFab,
			TipoEnvolvido:         record[14],
			Idade:                 idade,
			Sexo:                  record[19],
			RacaCor:               record[20],
			EstadoFisico:          record[21],
			CondicaoMetereologica: record[22],
			TipoPista:             record[23],
			TracadoVia:            record[24],
			UsoSolo:               record[25],
			SentidoVia:            record[26],
			CausaAcidente:         record[27],
			TipoAcidente:          record[28],
			ClassificacaoAcidente: record[29],
			QtdIlesos:             ilesos,
			QtdFeridosLeves:       feridosLeves,
			QtdFeridosGraves:      feridosGraves,
			QtdMortos:             mortos,
		}

		acidentes = append(acidentes, acidente)
	}

	return acidentes, nil
}

// ParseAcidentesSESMG parse acidentes da SES-MG
func ParseAcidentesSESMG(file io.Reader) ([]AcidenteTemp, error) {
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Pular cabeçalho
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var acidentes []AcidenteTemp

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		data, err := time.Parse("02/01/2006", record[0])
		if err != nil {
			continue
		}

		idade, _ := strconv.Atoi(record[3])

		acidente := AcidenteTemp{
			DataCompleta:          data,
			Municipio:             record[1],
			Idade:                 idade,
			Sexo:                  record[4],
			RacaCor:               record[5],
			CausaAcidente:         record[6],
			CIDCausaMorte:         record[7],
			DescCausaMorte:        record[8],
			QtdMortos:             1,
			QtdIlesos:             0,
			QtdFeridosLeves:       0,
			QtdFeridosGraves:      0,
			ClassificacaoAcidente: "COM VÍTIMAS FATAIS",
		}

		acidentes = append(acidentes, acidente)
	}

	return acidentes, nil
}
