package handler

import (
	"net/http"
	"surekapi/helper"
	"surekapi/naskahdinas"

	"github.com/gin-gonic/gin"
)

type naskahDinasHandler struct {
	service naskahdinas.Service
}

func NewNaskahDinasHandler(service naskahdinas.Service) *naskahDinasHandler {
	return &naskahDinasHandler{service}
}

func (h *naskahDinasHandler) FindNaskahDinas(c *gin.Context) {
	naskahDinas, err := h.service.FindNaskahDinas()
	if err != nil {
		response := helper.APIResponse("Failed to get naskah dinas", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := naskahdinas.FormatMultipleNaskahDinas(naskahDinas)
	response := helper.APIResponse("List of naskah dinas", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
