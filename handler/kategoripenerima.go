package handler

import (
	"net/http"
	"surekapi/helper"
	"surekapi/recipient_category"

	"github.com/gin-gonic/gin"
)

type recipientCategoryHandler struct {
	service recipient_category.Service
}

func NewRecipientCategoryHandler(service recipient_category.Service) *recipientCategoryHandler {
	return &recipientCategoryHandler{service}
}

func (h *recipientCategoryHandler) GetRecipientCategory(c *gin.Context) {
	recipientCategory, err := h.service.GetRecipientCategory()
	if err != nil {
		response := helper.APIResponse("Failed to get recipient category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of recipient category", http.StatusOK, "success", recipientCategory)
	c.JSON(http.StatusOK, response)

}
