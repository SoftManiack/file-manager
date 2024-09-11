package handler

import (
	files "file-manager/dto"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateFile(c *gin.Context) {

	var input files.UpdateFile

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, err := h.services.UpdateFile(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data:    file,
	})
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

	// проверить нет ли файла с таким именем

	filesDir, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range filesDir {
		if file.Name() == newFile.Name && file.IsDir() == false {
			newErrorResponse(c, http.StatusBadRequest, "Файл с таким именем существует")
			return
		}
	}

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

	var input files.DownloadFile
	var path string

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userUid, err := getUserUid(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if input.UidRoot == userUid {
		path = os.Getenv("PATH_FILES") + "/" + userUid + "/" + input.Name
	} else {
		path = os.Getenv("PATH_FILES") + "/" + userUid + "/" + input.UidRoot + "/" + input.Name
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+input.Name)
	c.Header("Content-Type", "application/octet-stream")
	c.File(path)
}

// Создать текст файл вирнуть из директории
// Изменить текстовый файл
//
