package csv

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/helpers"
)

// Estruturas auxiliares para organizar os dados antes de criar o fato
type AcidenteTemp struct {
	// Dados temporais
	DataCompleta time.Time
	DiaSemana    string
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

func ParseAcidentesPRF(file io.Reader) ([]AcidenteTemp, error) {
	reader := csv.NewReader(file)
	reader.Comma = ';'       // O arquivo CSV está separado por tabulação
	reader.LazyQuotes = true // Permite que o parser ignore problemas com aspas não fechadas

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

		// Parse do campo de data
		data, err := time.Parse("2006-01-02", record[2]) // "2024-01-01" no formato de data
		if err != nil {
			continue
		}

		// Parse de outros campos
		km, _ := strconv.Atoi(record[7])                   // 'km' está no índice 7
		idade, _ := strconv.Atoi(record[19])               // 'idade' está no índice 19
		anoFab, _ := strconv.Atoi(record[21])              // 'ano_fabricacao_veiculo' está no índice 21
		ilesos, _ := strconv.Atoi(record[25])              // 'ilesos' está no índice 25
		feridosLeves, _ := strconv.Atoi(record[26])        // 'feridos_leves' está no índice 26
		feridosGraves, _ := strconv.Atoi(record[27])       // 'feridos_graves' está no índice 27
		mortos, _ := strconv.Atoi(record[28])              // 'mortos' está no índice 28
		latitude, _ := strconv.ParseFloat(record[31], 64)  // 'latitude' está no índice 31
		longitude, _ := strconv.ParseFloat(record[32], 64) // 'longitude' está no índice 32

		acidente := AcidenteTemp{
			DataCompleta:          data,
			DiaSemana:             helpers.RemoveSpecialChars(record[3]),  // dia_semana
			Hora:                  helpers.RemoveSpecialChars(record[4]),  // horario
			PeriodoDia:            helpers.RemoveSpecialChars(record[10]), // fase_di
			BR:                    helpers.RemoveSpecialChars(record[5]),  // br
			KM:                    km,
			Municipio:             helpers.RemoveSpecialChars(record[7]),  // municipio
			CausaAcidente:         helpers.RemoveSpecialChars(record[8]),  // causa_acidente
			TipoAcidente:          helpers.RemoveSpecialChars(record[9]),  // tipo_acidente
			ClassificacaoAcidente: helpers.RemoveSpecialChars(record[10]), // classificacao_acidente
			SentidoVia:            helpers.RemoveSpecialChars(record[11]), // sentido_via
			CondicaoMetereologica: helpers.RemoveSpecialChars(record[12]), // condicao_metereologica
			TipoPista:             helpers.RemoveSpecialChars(record[13]), // tipo_pista
			TracadoVia:            helpers.RemoveSpecialChars(record[14]), // tracado_via
			UsoSolo:               helpers.RemoveSpecialChars(record[15]), // uso_solo
			TipoVeiculo:           helpers.RemoveSpecialChars(record[17]), // tipo_veiculo
			Marca:                 helpers.RemoveSpecialChars(record[18]), // marca
			AnoFabricacao:         anoFab,
			TipoEnvolvido:         helpers.RemoveSpecialChars(record[20]), // tipo_envolvido
			EstadoFisico:          helpers.RemoveSpecialChars(record[21]), // estado_fisico
			Idade:                 idade,
			Sexo:                  helpers.RemoveSpecialChars(record[22]), // sexo
			QtdIlesos:             ilesos,
			QtdFeridosLeves:       feridosLeves,
			QtdFeridosGraves:      feridosGraves,
			QtdMortos:             mortos,
			Latitude:              latitude,
			Longitude:             longitude,
			Regional:              helpers.RemoveSpecialChars(record[29]), // regional
			Delegacia:             helpers.RemoveSpecialChars(record[30]), // delegacia
			UOP:                   helpers.RemoveSpecialChars(record[31]), // uop
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
			Municipio:             record[6],
			Idade:                 idade,
			Sexo:                  record[3],
			RacaCor:               record[4],
			CausaAcidente:         record[8],
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
