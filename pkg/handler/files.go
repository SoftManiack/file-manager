package handler

import (
	files "file-manager/dto"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateFile(c *gin.Context) {

	var input files.UpdateFile
	var path string

	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, err := h.services.UpdateFile(input)

	fmt.Println(file)

	fmt.Println(userUid)
	fmt.Println(file.Uid)
	fmt.Println(err)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		_ = h.services.UpdateRecent(userUid, file.Uid)
	}

	if userUid == input.RootUid {
		path = os.Getenv("PATH_FILES") + "/" + input.RootUid + "/"
	} else {
		path = os.Getenv("PATH_FILES") + "/" + userUid + "/" + input.RootUid + "/"
	}

	err = os.Rename(path+input.OldName, path+input.Name)

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

	saveFile, err := h.services.CreateFile(newFile)

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

func (h *Handler) MoveTrashFile(c *gin.Context) {

	fmt.Println("delete file")
	var input files.UidFile

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userUid, err := getUserUid(c)

	fmt.Println(input)
	fmt.Println(userUid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.MoveTrashFile(userUid, input.Uid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
	})
}

func (h *Handler) CreateTextFile(c *gin.Context) {

	var newFile files.NewFile
	var input files.NewTextFile

	//var saveFile files.File
	var path string

	if input.Name != "" {
		newErrorResponse(c, http.StatusBadRequest, "имя не должно быть пустым")
		return
	}

	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		return
	}

	if input.RootUid == userUid {
		path = os.Getenv("PATH_FILES") + "/" + input.RootUid
	} else {
		path = os.Getenv("PATH_FILES") + "/" + userUid + "/" + input.RootUid
	}

	newFile.Size = 0
	newFile.Name = input.Name
	newFile.Path = path
	newFile.RootUid = input.RootUid
	newFile.Data = []byte(input.Text)

	fmt.Println(newFile)

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

	saveFile, err := h.services.CreateFile(newFile)

	// создать текстовый файл

	fmt.Println(path + input.Name + ".txt")

	if err == nil {
		file, err := os.Create(path + "/" + input.Name + ".txt")

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			os.Exit(1)
		}

		defer file.Close()
		file.WriteString(input.Text)

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

func (h *Handler) UpdateTextFile(c *gin.Context) {

	var input files.UpdateTextFile
	var path string

	//var saveFile files.File

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	if input.RootUid == userUid {
		path = os.Getenv("PATH_FILES") + "/" + input.RootUid
	} else {
		path = os.Getenv("PATH_FILES") + "/" + userUid + "/" + input.RootUid
	}

	err = h.services.UpdateTextFile(input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(path + "/" + input.Name + ".txt")

	if err == nil {

		err = os.WriteFile(path+"/"+input.Name+".txt", []byte(input.Text), 0644)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			os.Exit(1)
		}

	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data:    nil,
	})
}

func (h *Handler) CreateTable(c *gin.Context) {

}

func (h *Handler) CopyFile(c *gin.Context) {

	var input files.CopyFile
	var pathFile string
	var pathTargetDir string

	//var saveFile files.File

	fmt.Println(input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	err = h.services.CopyFile(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if input.RootUid == userUid {
		pathFile = os.Getenv("PATH_FILES") + "/" + input.RootUid
	} else {
		pathFile = os.Getenv("PATH_FILES") + "/" + userUid + "/" + input.RootUid
	}

	if input.RootDirUid == userUid {
		pathTargetDir = os.Getenv("PATH_FILES") + "/" + input.RootDirUid
	} else {
		pathTargetDir = os.Getenv("PATH_FILES") + "/" + userUid + "/" + input.RootDirUid
	}

	fmt.Println("1")
	fmt.Println(pathTargetDir)
	fmt.Println(pathFile + "/" + input.Name)

	cpCmd := exec.Command("cp", "-rf", pathFile+"/"+input.Name, pathTargetDir)

	err = cpCmd.Run()

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data:    nil,
	})
	// скопиров стар файл и вставить в новую директорию
}

func (h *Handler) MoverFile(c *gin.Context) {

}

// Создать текст файл вирнуть из директории
// Изменить текстовый файл
//
