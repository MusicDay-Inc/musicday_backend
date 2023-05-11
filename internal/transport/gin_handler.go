package transport

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
	user := router.Group("/user", h.authenticateClient)
	{
		// Отправляю ответ вместе с оценкой
		user.GET("/song/:id", h.getSongById)
		user.GET("/album_info/:id", h.getAlbumById)
		user.GET("/album_full/:id", h.getAlbumWitSongsById)

		// deleted
		//user.GET("/subscription_reviews/:id", h.ReviewsOfSubscriptions)

		// TODO
		user.GET("/activity", h.SearchSongs)
		// deleted
		//user.GET("/stories", h.SearchSongs)
	}

	// Профиль человека
	profile := router.Group("/profile", h.authenticateClient)
	{
		// Done
		profile.GET("/:id", h.getUserProfile)
		profile.GET("/subscribers/:id", h.getUserSubscribers)
		profile.GET("/subscriptions/:id", h.getUserSubscriptions)
	}
	// DONE!
	search := router.Group("/search", h.authenticateClient)
	{
		search.GET("/song", h.SearchSongs)
		search.GET("/album", h.SearchAlbums)
		search.GET("/user", h.SearchUsers)
	}
	// DONE!
	action := router.Group("/action", h.authenticateClient)
	{
		action.POST("/review/:id", h.reviewRelease)
		action.POST("/subscribe/:id", h.subscribe)
		action.POST("/username", h.changeUsername)
		action.POST("/nickname", h.changeNickname)
		action.POST("/delete_review/:id", h.deleteReviewById)
		action.POST("/unsubscribe/:id", h.unsubscribe)
		// deleted
		//action.POST("/post_story/:id", h.postStory)
	}
	// DONE!
	reviews := router.Group("/reviews", h.authenticateClient)
	{
		reviews.GET("/to_release/:id", h.ReviewsOfSubscriptions)
	}

	// DONE!
	library := router.Group("/library", h.authenticateClient)
	{
		// все обзоры пользователя
		library.GET("all/:id", h.UserAllReviews)
		// все песни
		library.GET("/songs/:id", h.UserSongReviews)
		// все альбомы
		library.GET("/albums/:id", h.UserAlbumReviews)
	}
	return router
}
