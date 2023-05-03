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

func (h *Handler) authenticateUser(context *gin.Context) {

}

func (h *Handler) identifyUser(context *gin.Context) {

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
	lib := router.Group("/lib", h.identifyUser)
	{
		// Отправляю ответ вместе с оценкой
		lib.GET("/song/:id", h.getSongById)
		lib.GET("/album/:id", h.getAlbumById)
		// Отправляю без оценки
		lib.GET("/song/search", h.SearchSongs)
		lib.GET("/album/search", h.SearchAlbums)
		lib.GET("/user/search", h.SearchUsers)
	}
	//api := router.Group("/api", h.userIdentity)
	api := router.Group("/api", h.authenticateUser)
	{
		api.POST("/song/:id", h.reviewSong)
		api.POST("/album/:id", h.reviewAlbum)
		api.POST("/story/:id", h.likeStory)
		api.POST("/user/:id", h.subscribe)

		//lists := api.Group("/lists")
		//{
		//	lists.POST("/", h.createList)
		//	lists.GET("/", h.getAllLists)
		//	lists.GET("/:id", h.getListById)
		//	lists.PUT("/:id", h.updateList)
		//	lists.DELETE("/:id", h.deleteList)
		//
		//	items := lists.Group(":id/items")
		//	{
		//		items.POST("/", h.createItem)
		//		items.GET("/", h.getAllItems)
		//	}
		//}
		//items := api.Group("items")
		//{
		//	items.GET("/:id", h.getItemById)
		//	items.PUT("/:id", h.updateItem)
		//	items.DELETE("/:id", h.deleteItem)
		//}
	}
	return router
}

func (h *Handler) getSongById(context *gin.Context) {
}

func (h *Handler) getAlbumById(context *gin.Context) {

}

func (h *Handler) SearchSongs(context *gin.Context) {

}

func (h *Handler) SearchAlbums(context *gin.Context) {

}

func (h *Handler) SearchUsers(context *gin.Context) {

}

func (h *Handler) reviewSong(context *gin.Context) {

}

func (h *Handler) likeStory(context *gin.Context) {

}

func (h *Handler) subscribe(context *gin.Context) {

}

func (h *Handler) reviewAlbum(context *gin.Context) {

}
