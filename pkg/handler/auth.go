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

	fmt.Println(input.Name)
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(input)
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
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)

	fmt.Println(token)

	if err.Error() == "неверная электронная почта или пароль" {
		newErrorResponse(c, 401, err.Error())
		return
	} else if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	uid, _ := h.services.ParseToken(token)
	fmt.Println(token)

	fmt.Println("err")

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
