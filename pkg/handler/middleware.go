package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userUid"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerPars := strings.Split(header, " ")
	if len(headerPars) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userUid, err := h.services.Authorization.ParseToken(headerPars[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userUid", userUid)
}

func getUserUid(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user uid not fount")
		return "", errors.New("user id not found")
	}

	userUid, ok := id.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user uid is of invalid type")
		return "", errors.New("user id not found")
	}

	return userUid, nil
}
