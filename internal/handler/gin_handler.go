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

	auth := router.Group("/auth")
	{
		auth.POST("/start", h.start)
		auth.POST("/sign_up", h.signUp)
	}
	user := router.Group("/user", h.authenticateUser)
	{
		// Отправляю ответ вместе с оценкой
		user.GET("/song/:id", h.getSongById)
		user.GET("/album_info/:id", h.getAlbumById)
		user.GET("/album_full/:id", h.getAlbumWitSongsById)
		user.GET("/subscription_reviews/:id", h.getSongById)

		// TODO
		user.GET("/activity", h.SearchSongs)
		user.GET("/stories", h.SearchSongs)

		// TODO
		// Профиль человека
		profile := router.Group("/profile")
		{
			profile.GET("/", h.SearchSongs)
			profile.GET("/:id", h.SearchSongs)
		}

	}
	// DONE!
	search := router.Group("/search", h.authenticateUser)
	{
		search.GET("/song", h.SearchSongs)
		search.GET("/album", h.SearchAlbums)
		search.GET("/user", h.SearchUsers)
	}
	action := router.Group("/action", h.authenticateUser)
	{
		// Done
		action.POST("/review/:id", h.reviewRelease)
		action.POST("/subscribe/:id", h.subscribe)
		action.POST("/username", h.changeUsername)
		action.POST("/nickname", h.changeNickname)

		// TODO
		action.POST("/post_story/:id", h.postStory)
	}
	// DONE!
	reviews := router.Group("/reviews", h.authenticateUser)
	{
		reviews.GET("/to_release/:id", h.ReviewsOfSubscribers)
	}

	// Песни человека с его оценками
	// Альбомы человека с его оценками
	library := router.Group("/library", h.authenticateUser)
	{
		// TODO
		// все для профиля
		library.GET("all/:id", h.SearchUsers)
		// все песни
		library.GET("/songs/:id", h.SearchSongs)
		// все альбомы
		library.GET("/albums/:id", h.SearchAlbums)
	}
	return router
}
