package handler

import (
	dto "file-manager/dto"
	files "file-manager/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTrash(c *gin.Context) {

	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	getDirectories, getFiles, err := h.services.GetTrash(userUid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data: struct {
			Files       []dto.File        `json:"files"`
			Directories []dto.Directories `json:"directories"`
		}{
			Files:       getFiles,
			Directories: getDirectories,
		},
	})
}

func (h *Handler) DeleteTrashFile(c *gin.Context) {

	var input files.DeleteFile

	userUid, err := getUserUid(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.DeleteTrashFile(input, userUid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data:    nil,
	})

}

func (h *Handler) DeleteTrashDirectory(c *gin.Context) {

}

func (h *Handler) RemoveTrashFile(c *gin.Context) {
	var input files.UidFile

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.RemoveTrashFile(input.Uid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data:    nil,
	})
}
