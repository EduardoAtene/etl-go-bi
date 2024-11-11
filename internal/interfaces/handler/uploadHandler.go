package handler

import (
	"net/http"

	"github.com/EduardoAtene/etl-go-bi/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	uploadUseCase *usecase.UploadUseCase
}

func NewUploadHandler(uploadUseCase *usecase.UploadUseCase) *UploadHandler {
	return &UploadHandler{
		uploadUseCase: uploadUseCase,
	}
}

func (h *UploadHandler) HandlePRF(c *gin.Context) {
	file, err := c.FormFile("acidentes_prf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "arquivo de acidentes PRF não fornecido"})
		return
	}

	err = h.uploadUseCase.ProcessPRFData(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "dados PRF importados com sucesso"})
}

func (h *UploadHandler) HandleSESMG(c *gin.Context) {
	file, err := c.FormFile("acidentes_sesmg")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "arquivo de acidentes SES-MG não fornecido"})
		return
	}

	err = h.uploadUseCase.ProcessSESMGData(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "dados SES-MG importados com sucesso"})
}
