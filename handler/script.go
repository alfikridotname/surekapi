package handler

import (
	"net/http"
	"surekapi/helper"
	"surekapi/script"

	"github.com/gin-gonic/gin"
)

type scriptHandler struct {
	service script.Service
}

func NewScriptHandler(service script.Service) *scriptHandler {
	return &scriptHandler{service}
}

func (h *scriptHandler) GetScripts(c *gin.Context) {
	scripts, err := h.service.GetScripts()
	if err != nil {
		response := helper.APIResponse("Failed to get script", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := script.FormatMultipleScript(scripts)
	response := helper.APIResponse("List of script", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
