package handler

import (
	"file-manager/pkg/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := router.Group("/v1/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		auth.DELETE("/user/:uid", h.userIdentity, h.deleteUser)
	}

	fm := router.Group("/v1/fm", h.userIdentity)
	{
		directories := fm.Group("/directories")
		{
			directories.GET("/:uid", h.GetDirectories)
			directories.POST("/add", h.CreateDirectories)
			directories.PATCH("/", h.UpdateDirectories)
			directories.POST("/:uid/upload", h.UploadFile)
		}

		files := fm.Group("/files")
		{
			files.PATCH("/", h.UpdateFile)
		}
	}

	return router
}
