package handler

import (
	files "file-manager/dto"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateFile(c *gin.Context) {

	/* var input files.NewTextFile
	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, err := h.services.CreateTextFile(userUid, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data:    file,
	}) */

}

func (h *Handler) UpdateFile(c *gin.Context) {

}

func (h *Handler) UploadFile(c *gin.Context) {

	var newFile files.NewFile
	//var saveFile files.File
	var path string
	var extension string

	rootDir := c.Param("uid")
	file, _ := c.FormFile("file")
	userUid, err := getUserUid(c)
	extension = strings.Split(file.Filename, ".")[1]

	if err != nil {
		return
	}

	if rootDir == userUid {
		path = os.Getenv("PATH_FILES") + "/" + rootDir + "/"
	} else {
		path = os.Getenv("PATH_FILES") + "/" + userUid + "/" + rootDir + "/"
	}

	newFile.Size = int64(file.Size)
	newFile.Name = file.Filename
	newFile.Path = path
	newFile.RootUid = rootDir
	newFile.Data = []byte{}

	if extension == "jpg" || extension == "png" || extension == "jpeg" || extension == "txt" {
		form, _ := c.MultipartForm()
		fileReader, _ := form.File["file"][0].Open()
		newFile.Data, _ = io.ReadAll(fileReader)
	}

	saveFile, err := h.services.UploadFile(newFile)

	if err == nil {
		c.SaveUploadedFile(file, path+file.Filename)
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data:    saveFile,
	})
}

func (h *Handler) DownloadFile(c *gin.Context) {

}

// Создать текст файл вирнуть из директории
// Изменить текстовый файл
//
