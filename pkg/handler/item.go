package handler

import (
	item "file-manager"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetItems(c *gin.Context) {

	rootuid := c.Param("uid")
	time.Sleep(8 * time.Second)

	if rootuid == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	items, err := h.services.GetItems(rootuid)

	fmt.Println(len(items))
	if len(items) == 1 {
		items = []item.Item{}
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) CreateDirectory(c *gin.Context) {

	var input item.NewDirectory
	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.services.CreateDirectory(input, userUid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) CreateTextFile(c *gin.Context) {

	var input item.NewFile
	userUid, err := getUserUid(c)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	item, err := h.services.CreateTextFile(input, userUid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) Rename(c *gin.Context) {

	var input item.Rename
	uid := c.Param("uid")

	if err := c.BindJSON(&input); err != nil || uid == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	err := h.services.Rename(uid, input.Name)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) Delete(c *gin.Context) {

	uid := c.Param("uid")
	userUid, err := getUserUid(c)

	if err != nil {
		return
	}

	if uid == userUid || uid == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	err = h.services.Delete(uid, userUid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
