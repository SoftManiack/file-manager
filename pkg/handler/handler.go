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
		AllowHeaders:     []string{"Authorization, User-Agent, Accept, Origin, Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := router.Group("/v1/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	fm := router.Group("/v1/fm", h.userIdentity)
	{
		items := fm.Group("/items")
		{
			items.GET("/:uid", h.GetItems)
			items.POST("/directory", h.CreateDirectory)
			items.POST("/file", h.CreateTextFile)
			items.PUT("/:uid", h.Rename)
			/*items.DELETE("/:uid", h.Delete)
			items.POST("/upload", h.Upload)*/
		}

		/*recent := router.Group("/recent")
		{
			recent.GET("/")
		}

		trash := router.Group("/trash")
		{
			trash.GET("/")
		}*/
	}

	return router
}
