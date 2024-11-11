package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/go-sql-driver/mysql"
)

type Acidente struct {
	SentidoVia            string
	CondicaoMetereologica string
	TipoPista             string
	TracadoVia            string
	UsoSolo               string
	Pessoas               int
	Mortos                int
	FeridosLeves          int
	FeridosGraves         int
	Ilesos                int
	Ignorados             int
	Feridos               int
	Veiculos              int
	Latitude              float64
	Longitude             float64
	Regional              string
	Delegacia             string
	Uop                   string
}

type Pessoa struct {
	DataObito           string
	DataNascimento      string
	Idade               int
	Sexo                string
	RacaCor             string
	MunicipioResidencia string
	MunicipioOcorrencia string
	CidCausaBasica      string
	DescCidCausaBasica  string
}

func main() {
	dsn := "root:sua_senha@tcp(localhost:3306)/DW_Acidentes?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	acidentes, err := lerExcelAcidentes("acidentes.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	pessoas, err := lerExcelPessoas("pessoas.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	err = processarDados(db, acidentes, pessoas)
	if err != nil {
		log.Fatal(err)
	}
}

func lerExcelAcidentes(arquivo string) ([]Acidente, error) {
	f, err := excelize.OpenFile(arquivo)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var acidentes []Acidente
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		pessoas, _ := strconv.Atoi(row[5])
		mortos, _ := strconv.Atoi(row[6])
		feridosLeves, _ := strconv.Atoi(row[7])
		feridosGraves, _ := strconv.Atoi(row[8])
		ilesos, _ := strconv.Atoi(row[9])
		ignorados, _ := strconv.Atoi(row[10])
		feridos, _ := strconv.Atoi(row[11])
		veiculos, _ := strconv.Atoi(row[12])
		latitude, _ := strconv.ParseFloat(row[13], 64)
		longitude, _ := strconv.ParseFloat(row[14], 64)

		acidente := Acidente{
			SentidoVia:            row[0],
			CondicaoMetereologica: row[1],
			TipoPista:             row[2],
			TracadoVia:            row[3],
			UsoSolo:               row[4],
			Pessoas:               pessoas,
			Mortos:                mortos,
			FeridosLeves:          feridosLeves,
			FeridosGraves:         feridosGraves,
			Ilesos:                ilesos,
			Ignorados:             ignorados,
			Feridos:               feridos,
			Veiculos:              veiculos,
			Latitude:              latitude,
			Longitude:             longitude,
			Regional:              row[15],
			Delegacia:             row[16],
			Uop:                   row[17],
		}
		acidentes = append(acidentes, acidente)
	}
	return acidentes, nil
}

func lerExcelPessoas(arquivo string) ([]Pessoa, error) {
	f, err := excelize.OpenFile(arquivo)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var pessoas []Pessoa
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		idade, _ := strconv.Atoi(row[2])
		pessoa := Pessoa{
			DataObito:           row[0],
			DataNascimento:      row[1],
			Idade:               idade,
			Sexo:                row[3],
			RacaCor:             row[4],
			MunicipioResidencia: row[5],
			MunicipioOcorrencia: row[6],
			CidCausaBasica:      row[7],
			DescCidCausaBasica:  row[8],
		}
		pessoas = append(pessoas, pessoa)
	}
	return pessoas, nil
}

func processarDados(db *sql.DB, acidentes []Acidente, pessoas []Pessoa) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, acidente := range acidentes {
		var idCondicoes, idLocalizacao, idTempo int64

		stmtCondicoes, err := tx.Prepare(`
            INSERT INTO Dim_Condicoes (condicao_meteorologica, tipo_pista, tracado_via, uso_solo, sentido_via) 
            VALUES (?, ?, ?, ?, ?)`)
		if err != nil {
			return err
		}
		res, err := stmtCondicoes.Exec(
			acidente.CondicaoMetereologica,
			acidente.TipoPista,
			acidente.TracadoVia,
			acidente.UsoSolo,
			acidente.SentidoVia,
		)
		if err != nil {
			return err
		}
		idCondicoes, _ = res.LastInsertId()
		stmtCondicoes.Close()

		stmtLocalizacao, err := tx.Prepare(`
            INSERT INTO Dim_Localizacao (latitude, longitude, regional, delegacia, uop) 
            VALUES (?, ?, ?, ?, ?)`)
		if err != nil {
			return err
		}
		res, err = stmtLocalizacao.Exec(
			acidente.Latitude,
			acidente.Longitude,
			acidente.Regional,
			acidente.Delegacia,
			acidente.Uop,
		)
		if err != nil {
			return err
		}
		idLocalizacao, _ = res.LastInsertId()
		stmtLocalizacao.Close()

		now := time.Now()
		stmtTempo, err := tx.Prepare(`
            INSERT INTO Dim_Tempo (data_completa, dia, mes, ano, dia_semana, hora, periodo_dia) 
            VALUES (?, ?, ?, ?, ?, ?, ?)`)
		if err != nil {
			return err
		}
		res, err = stmtTempo.Exec(
			now,
			now.Day(),
			now.Month(),
			now.Year(),
			now.Weekday().String(),
			now.Format("15:04:05"),
			getPeriodoDia(now.Hour()),
		)
		if err != nil {
			return err
		}
		idTempo, _ = res.LastInsertId()
		stmtTempo.Close()

		stmtFato, err := tx.Prepare(`
            INSERT INTO Fato_Acidentes (id_tempo, id_localizacao, id_condicoes, qtd_ilesos, qtd_feridos_leves, 
            qtd_feridos_graves, qtd_mortos) 
            VALUES (?, ?, ?, ?, ?, ?, ?)`)
		if err != nil {
			return err
		}
		_, err = stmtFato.Exec(
			idTempo,
			idLocalizacao,
			idCondicoes,
			acidente.Ilesos,
			acidente.FeridosLeves,
			acidente.FeridosGraves,
			acidente.Mortos,
		)
		if err != nil {
			return err
		}
		stmtFato.Close()
	}

	for _, pessoa := range pessoas {
		stmtPessoa, err := tx.Prepare(`
            INSERT INTO Dim_Pessoa (tipo_envolvido, estado_fisico, idade, sexo, raca_cor) 
            VALUES (?, ?, ?, ?, ?)`)
		if err != nil {
			return err
		}
		_, err = stmtPessoa.Exec(
			"Não Informado",
			"Não Informado",
			pessoa.Idade,
			pessoa.Sexo,
			pessoa.RacaCor,
		)
		if err != nil {
			return err
		}
		stmtPessoa.Close()
	}

	return tx.Commit()
}

func getPeriodoDia(hora int) string {
	switch {
	case hora >= 5 && hora < 12:
		return "Manhã"
	case hora >= 12 && hora < 18:
		return "Tarde"
	case hora >= 18 && hora < 24:
		return "Noite"
	default:
		return "Madrugada"
	}
}
