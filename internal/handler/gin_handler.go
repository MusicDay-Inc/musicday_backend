package handler

import (
	"github.com/gin-gonic/gin"
	"server/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO прикрутить хедер авторизации
	auth := router.Group("/auth")
	{
		auth.POST("/start", h.start)
		auth.POST("/sign_up", h.signUp)
	}
	lib := router.Group("/library", h.authenticateUser)
	{
		// Отправляю ответ вместе с оценкой
		lib.GET("/song/:id", h.getSongById)
		lib.GET("/album_info/:id", h.getAlbumById)
		lib.GET("/album_full/:id", h.getAlbumWitSongsById)

		// Отправляю без оценки
		lib.GET("/song/search", h.SearchSongs)
		lib.GET("/album/search", h.SearchAlbums)
		lib.GET("/user/search", h.SearchUsers)
	}
	//api := router.Group("/api", h.userIdentity)
	api := router.Group("/api", h.authenticateUser)
	{
		api.POST("/user/:id", h.subscribe)
		api.POST("/story/:id", h.likeStory)
		api.POST("/song/:id", h.reviewSong)
		api.POST("/album/:id", h.reviewAlbum)
	}
	return router
}
