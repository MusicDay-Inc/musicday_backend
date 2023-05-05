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
		// TODO
		lib.GET("/search_song", h.SearchSongs)
		lib.GET("/search_album", h.SearchAlbums)
		lib.GET("/search_user", h.SearchUsers)
	}
	//api := router.Group("/api", h.userIdentity)
	api := router.Group("/action", h.authenticateUser)
	{
		api.POST("/sub_to_user/:id", h.subscribe)
		api.POST("/like_story/:id", h.likeStory)
		api.POST("/post_story/:id", h.likeStory)
		// Done
		api.POST("/review/:id", h.reviewRelease)
	}
	return router
}
