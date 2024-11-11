package usecase

import (
	"fmt"
	"mime/multipart"

	"github.com/EduardoAtene/etl-go-bi/internal/infrastructure/csv"
	"github.com/EduardoAtene/etl-go-bi/internal/infrastructure/database"
)

type UploadUseCase struct {
	db *database.MySQLConnection
}

func NewUploadUseCase(db *database.MySQLConnection) *UploadUseCase {
	return &UploadUseCase{
		db: db,
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

	// Processar cada acidente e inserir nas dimensões e fatos
	for _, acidente := range acidentes {
		// TODO: Implementar a lógica de inserção nas tabelas dimensão
		// 1. Inserir/buscar na Dim_Tempo
		// 2. Inserir/buscar na Dim_Localizacao
		// 3. Inserir/buscar na Dim_Veiculo
		// 4. Inserir/buscar na Dim_Pessoa
		// 5. Inserir/buscar na Dim_Condicoes
		// 6. Inserir na tabela Fato_Acidentes
		fmt.Println(acidente)
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

	// Processar cada acidente e inserir nas dimensões e fatos
	for _, acidente := range acidentes {
		// TODO: Implementar a lógica de inserção nas tabelas dimensão
		// Similar ao processo PRF, mas com menos informações
		fmt.Println(acidente)
	}

	return tx.Commit()
}
