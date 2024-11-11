package csv

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
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

func ParseAcidentesPRF(file io.Reader) ([]entity.FatoAcidentes, error) {
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Pular cabeçalho
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var fatos []entity.FatoAcidentes

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Convertendo valores numéricos
		ilesos, _ := strconv.Atoi(record[9])
		feridosLeves, _ := strconv.Atoi(record[7])
		feridosGraves, _ := strconv.Atoi(record[8])
		mortos, _ := strconv.Atoi(record[6])
		latitude, _ := strconv.ParseFloat(record[13], 64)
		longitude, _ := strconv.ParseFloat(record[14], 64)

		// Criando o fato
		fato := entity.FatoAcidentes{
			FonteDados:       "PRF",
			QtdIlesos:        ilesos,
			QtdFeridosLeves:  feridosLeves,
			QtdFeridosGraves: feridosGraves,
			QtdMortos:        mortos,
		}

		// Os IDs (IDTempo, IDLocalizacao, etc.) devem ser preenchidos após inserir nas dimensões
		// Aqui você precisará implementar a lógica para:
		// 1. Inserir/buscar na Dim_Tempo
		// 2. Inserir/buscar na Dim_Localizacao
		// 3. Inserir/buscar na Dim_Veiculo
		// 4. Inserir/buscar na Dim_Pessoa
		// 5. Inserir/buscar na Dim_Condicoes

		fatos = append(fatos, fato)
	}

	return fatos, nil
}

func ParseAcidentesSESMG(file io.Reader) ([]entity.FatoAcidentes, error) {
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Pular cabeçalho
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var fatos []entity.FatoAcidentes

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Convertendo a data do óbito
		dataObito, err := time.Parse("02/01/2006", record[0])
		if err != nil {
			continue // Pula registros com data inválida
		}

		// Para SES-MG, consideramos apenas óbitos
		fato := entity.FatoAcidentes{
			FonteDados:       "SES-MG",
			CIDCausaMorte:    record[7],
			DescCausaMorte:   record[8],
			QtdMortos:        1, // Cada registro representa um óbito
			QtdIlesos:        0,
			QtdFeridosLeves:  0,
			QtdFeridosGraves: 0,
		}

		// Os IDs (IDTempo, IDLocalizacao, etc.) devem ser preenchidos após inserir nas dimensões
		// Similar ao processo da PRF, mas com menos informações disponíveis

		fatos = append(fatos, fato)
	}

	return fatos, nil
}
