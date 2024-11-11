package main

import (
	"log"

	"github.com/EduardoAtene/etl-go-bi/internal/config"
	"github.com/EduardoAtene/etl-go-bi/internal/infrastructure/database"
	"github.com/EduardoAtene/etl-go-bi/internal/interfaces/handler"
	"github.com/EduardoAtene/etl-go-bi/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.NewMySQLConnection(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	uploadUseCase := usecase.NewUploadUseCase(db)
	uploadHandler := handler.NewUploadHandler(uploadUseCase)

	api := router.Group("/api/v1")
	{
		api.POST("/upload", uploadHandler.Handle)
	}

	log.Fatal(router.Run(":8080"))
}
