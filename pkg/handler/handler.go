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
			directories.GET("/download", h.DownloadDirectory)
			directories.DELETE("/", h.MoveTrashDirectory)
		}

		files := fm.Group("/files")
		{
			files.GET("/download", h.DownloadFile)
			files.POST("/text/add", h.CreateTextFile)
			files.PATCH("/", h.UpdateFile)
			files.DELETE("/delete", h.MoveTrashFile)
		}
	}

	return router
}

// 1) Скачивание файла #
// 2) Скачивание директории
// 3) Удаление файла
// 4) Удаление директории
// 5) Вывод корзины
// 6) Перемещение файла
// 7) Перемещение диреткории
// 8) Копирование файла
// 9) Копирование директории
// 10) Недавние файлы
// 11) Ссылка на файл
// 12) Сылка на диреткорию
// 13) В контейнер
// 14) Базу разворачиваем на сервере
// 15) ci-cd github
// 16) созд тектс файл
// 17) Редактировать текст файл
