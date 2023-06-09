package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core"
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
		user.GET("/activity", h.getUserActivityFeed)
		// deleted
		//user.GET("/subscription_reviews/:id", h.ReviewsOfSubscriptions)
		//user.GET("/stories", h.SearchSongs)
	}

	// Профиль человека
	profile := router.Group("/profile", h.authenticateClient)
	{
		profile.GET("/:id", h.getUserProfile)
		profile.GET("/subscribers/:id", h.getUserSubscribers)
		profile.GET("/subscriptions/:id", h.getUserSubscriptions)
	}
	search := router.Group("/search", h.authenticateClient)
	{
		search.GET("/song", h.SearchSongs)
		search.GET("/album", h.SearchAlbums)
		search.GET("/user", h.SearchUsers)
	}
	action := router.Group("/action", h.authenticateClient)
	{
		action.POST("/review/:id", h.reviewRelease)
		action.POST("/subscribe/:id", h.subscribe)
		action.POST("/unsubscribe/:id", h.unsubscribe)
		action.POST("/username", h.changeUsername)
		action.POST("/nickname", h.changeNickname)
		action.POST("/delete_review/:id", h.deleteReviewById)
		action.POST("/bio", h.CreateClientBio)
		action.POST("/avatar", h.PostAvatar)
		action.POST("/player_id/:id", h.postPlayerId)
		action.POST("/init_appid", h.initPlayerIdTable)
		// deleted
		//action.POST("/post_story/:id", h.postStory)
	}
	reviews := router.Group("/reviews", h.authenticateClient)
	{
		reviews.GET("/to_release/:id", h.ReviewsOfSubscriptions)
	}
	library := router.Group("/library", h.authenticateClient)
	{
		// все обзоры пользователя
		library.GET("all/:id", h.UserAllReviews)
		// все песни
		library.GET("/songs/:id", h.UserSongReviews)
		// все альбомы
		library.GET("/albums/:id", h.UserAlbumReviews)
	}
	img := router.Group("/image")
	{
		//library.GET("/release/:id", h.UserAllReviews)
		//library.GET("/user/:id", h.UserAllReviews)
		img.GET("/release/:id", h.getReleaseCover)
		img.GET("/avatar/:id", h.getAvatar)
		img.POST("/album/:id", h.uploadAlbumCover)
	}
	return router
}

func (h *Handler) initPlayerIdTable(c *gin.Context) {
	err := h.services.User.InitPlayerID()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "InitPlayerID Error")
		return
	}
	c.JSON(http.StatusOK, true)
}
