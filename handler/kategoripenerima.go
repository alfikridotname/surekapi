package handler

import (
	"net/http"
	"surekapi/helper"
	"surekapi/kategoripenerima"

	"github.com/gin-gonic/gin"
)

type kategoriPenerimaHandler struct {
	service kategoripenerima.Service
}

func NewKategoriPenerimaHandler(service kategoripenerima.Service) *kategoriPenerimaHandler {
	return &kategoriPenerimaHandler{service}
}

func (h *kategoriPenerimaHandler) FindKategoriPenerima(c *gin.Context) {
	katPenerima, err := h.service.FindKategoriPenerima()
	if err != nil {
		response := helper.APIResponse("Failed to get kategori penerima", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of kategori penerima", http.StatusOK, "success", katPenerima)
	c.JSON(http.StatusOK, response)

}
