package handler

import (
	user "file-manager/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var input user.User

	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, "Не верный пароль или почта")
		return
	}

	uid, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data: struct {
			Uid string `json:"uid"`
		}{
			Uid: uid,
		},
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, "Не верный пароль или почта")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)

	if err != nil && err.Error() == "Неверная электронная почта или пароль" {
		newErrorResponse(c, 400, err.Error())
		return
	} else if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	uid, err := h.services.Authorization.ParseToken(token)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: "success",
		Data: struct {
			Uid   string `json:"uid"`
			Token string `json:"token"`
			Email string `json:"email"`
			Size  string `json:"size"`
		}{
			Uid:   uid,
			Token: token,
			Email: input.Email,
			Size:  "0",
		},
	})
}

func (h *Handler) deleteUser(c *gin.Context) {
	uidUser := c.Param("uid")

	err := h.services.DeleteUser(uidUser)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		Message: fmt.Sprintf("Удлаен пользователь с uid %s", uidUser),
	})

}
