package handler

import (
	dto "file-manager/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRecent(c *gin.Context) {

	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	getFiles, err := h.services.GetRecent(userUid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data: struct {
			Files []dto.File `json:"files"`
		}{
			Files: getFiles,
		},
	})
}
