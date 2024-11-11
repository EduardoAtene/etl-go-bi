package usecase

import (
	"mime/multipart"
	"time"

	"github.com/EduardoAtene/etl-go-bi/internal/domain/entity"
	"github.com/EduardoAtene/etl-go-bi/internal/domain/repository"
	"github.com/EduardoAtene/etl-go-bi/internal/infrastructure/csv"
	"github.com/EduardoAtene/etl-go-bi/internal/infrastructure/database"
)

type UploadUseCase struct {
	db                 *database.MySQLConnection
	dimTempoRepo       *repository.DimTempoRepository
	dimLocalizacaoRepo *repository.DimLocalizacaoRepository
	dimVeiculoRepo     *repository.DimVeiculoRepository
	dimPessoaRepo      *repository.DimPessoaRepository
	dimCondicoesRepo   *repository.DimCondicoesRepository
	fatoAcidentesRepo  *repository.FatoAcidentesRepository
}

func NewUploadUseCase(db *database.MySQLConnection) *UploadUseCase {
	return &UploadUseCase{
		db:                 db,
		dimTempoRepo:       repository.NewDimTempoRepository(db.Conn),
		dimLocalizacaoRepo: repository.NewDimLocalizacaoRepository(db.Conn),
		dimVeiculoRepo:     repository.NewDimVeiculoRepository(db.Conn),
		dimPessoaRepo:      repository.NewDimPessoaRepository(db.Conn),
		dimCondicoesRepo:   repository.NewDimCondicoesRepository(db.Conn),
		fatoAcidentesRepo:  repository.NewFatoAcidentesRepository(db.Conn),
	}
}

func (u *UploadUseCase) ProcessPRFData(file *multipart.FileHeader) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	acidentes, err := csv.ParseAcidentesPRF(fileReader)
	if err != nil {
		return err
	}

	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, acidente := range acidentes {
		// var hora time.Time
		hora, _ := time.Parse("15:04", acidente.Hora)
		// 1. Inserir/buscar na Dim_Tempo
		dimTempo := &entity.DimTempo{
			DataCompleta: acidente.DataCompleta,
			Hora:         hora,
			PeriodoDia:   acidente.PeriodoDia,
			DiaSemana:    acidente.DataCompleta.Weekday().String(),
			Dia:          acidente.DataCompleta.Day(),
			Mes:          int(acidente.DataCompleta.Month()),
			Ano:          acidente.DataCompleta.Year(),
		}
		idTempo, err := u.dimTempoRepo.Insert(dimTempo)
		if err != nil {
			return err
		}

		// 2. Inserir/buscar na Dim_Localizacao
		dimLocalizacao := &entity.DimLocalizacao{
			Municipio: acidente.Municipio,
			BR:        acidente.BR,
			KM:        acidente.KM,
			Latitude:  acidente.Latitude,
			Longitude: acidente.Longitude,
			Regional:  acidente.Regional,
			Delegacia: acidente.Delegacia,
			UOP:       acidente.UOP,
		}
		idLocalizacao, err := u.dimLocalizacaoRepo.Insert(dimLocalizacao)
		if err != nil {
			return err
		}

		// 3. Inserir/buscar na Dim_Veiculo
		dimVeiculo := &entity.DimVeiculo{
			TipoVeiculo:   acidente.TipoVeiculo,
			Marca:         acidente.Marca,
			AnoFabricacao: acidente.AnoFabricacao,
		}
		idVeiculo, err := u.dimVeiculoRepo.Insert(dimVeiculo)
		if err != nil {
			return err
		}

		// 4. Inserir/buscar na Dim_Pessoa
		dimPessoa := &entity.DimPessoa{
			TipoEnvolvido: acidente.TipoEnvolvido,
			Idade:         acidente.Idade,
			Sexo:          acidente.Sexo,
			RacaCor:       acidente.RacaCor,
			EstadoFisico:  acidente.EstadoFisico,
		}
		idPessoa, err := u.dimPessoaRepo.Insert(dimPessoa)
		if err != nil {
			return err
		}

		// 5. Inserir/buscar na Dim_Condicoes
		dimCondicoes := &entity.DimCondicoes{
			CondicaoMeteorologica: acidente.CondicaoMetereologica,
			TipoPista:             acidente.TipoPista,
			TracadoVia:            acidente.TracadoVia,
			UsoSolo:               acidente.UsoSolo,
			SentidoVia:            acidente.SentidoVia,
		}
		idCondicoes, err := u.dimCondicoesRepo.Insert(dimCondicoes)
		if err != nil {
			return err
		}

		// 6. Inserir na tabela Fato_Acidentes
		fatoAcidente := &entity.FatoAcidentes{
			IDTempo:               idTempo,
			IDLocalizacao:         idLocalizacao,
			IDVeiculo:             idVeiculo,
			IDPessoa:              idPessoa,
			IDCondicoes:           idCondicoes,
			FonteDados:            "PRF",
			CausaAcidente:         acidente.CausaAcidente,
			TipoAcidente:          acidente.TipoAcidente,
			ClassificacaoAcidente: acidente.ClassificacaoAcidente,
			QtdIlesos:             acidente.QtdIlesos,
			QtdFeridosLeves:       acidente.QtdFeridosLeves,
			QtdFeridosGraves:      acidente.QtdFeridosGraves,
			QtdMortos:             acidente.QtdMortos,
		}
		_, err = u.fatoAcidentesRepo.Insert(fatoAcidente)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (u *UploadUseCase) ProcessSESMGData(file *multipart.FileHeader) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	acidentes, err := csv.ParseAcidentesSESMG(fileReader)
	if err != nil {
		return err
	}

	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, acidente := range acidentes {
		// 1. Inserir/buscar na Dim_Tempo
		dimTempo := &entity.DimTempo{
			DataCompleta: acidente.DataCompleta,
			DiaSemana:    acidente.DataCompleta.Weekday().String(),
			Dia:          acidente.DataCompleta.Day(),
			Mes:          int(acidente.DataCompleta.Month()),
			Ano:          acidente.DataCompleta.Year(),
		}
		idTempo, err := u.dimTempoRepo.Insert(dimTempo)
		if err != nil {
			return err
		}

		// 2. Inserir/buscar na Dim_Localizacao
		dimLocalizacao := &entity.DimLocalizacao{
			Municipio: acidente.Municipio,
		}
		idLocalizacao, err := u.dimLocalizacaoRepo.Insert(dimLocalizacao)
		if err != nil {
			return err
		}

		// 3. Inserir/buscar na Dim_Pessoa
		dimPessoa := &entity.DimPessoa{
			Idade:   acidente.Idade,
			Sexo:    acidente.Sexo,
			RacaCor: acidente.RacaCor,
		}
		idPessoa, err := u.dimPessoaRepo.Insert(dimPessoa)
		if err != nil {
			return err
		}

		// 4. Inserir na tabela Fato_Acidentes
		fatoAcidente := &entity.FatoAcidentes{
			IDTempo:               idTempo,
			IDLocalizacao:         idLocalizacao,
			IDPessoa:              idPessoa,
			FonteDados:            "SES-MG",
			CausaAcidente:         acidente.CausaAcidente,
			ClassificacaoAcidente: acidente.ClassificacaoAcidente,
			CIDCausaMorte:         acidente.CIDCausaMorte,
			DescCausaMorte:        acidente.DescCausaMorte,
			QtdMortos:             1,
			QtdIlesos:             0,
			QtdFeridosLeves:       0,
			QtdFeridosGraves:      0,
		}
		_, err = u.fatoAcidentesRepo.Insert(fatoAcidente)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
