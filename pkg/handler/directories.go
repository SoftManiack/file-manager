package handler

import (
	dto "file-manager/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateDirectories(c *gin.Context) {

	var input dto.NewDirectory
	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newDirectories, err := h.services.CreateDirectory(userUid, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data:    newDirectories,
	})

}

func (h *Handler) GetDirectories(c *gin.Context) {

	uidDir := c.Param("uid")

	getDirectories, getFiles, err := h.services.GetDirectory(uidDir)

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

func (h *Handler) UpdateDirectories(c *gin.Context) {
	var input dto.UpdateDirectory

	fmt.Println(input)
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	directory, err := h.services.UpdateDirectory(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data:    directory,
	})

}

func (h *Handler) DownloadDirectory(c *gin.Context) {

	// найти в базе все id на файлы и папки

	//

	// построить граф
	// создать временую папку
	// переместтиь в нее все файлы и папки
	// сжать
	// скачать
	// удалить

	uid := c.Param("uid")

	err := h.services.DownloadDirectory(uid)

	if err != nil {
		return
	}

}

func (h *Handler) MoveTrashDirectory(c *gin.Context) {

}
