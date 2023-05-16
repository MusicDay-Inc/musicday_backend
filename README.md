# Серверная часть MusicDay

## Курсового проекта "Мобильное приложение-аудиотека для обмена музыкаль- ными предпочтениями «MusicDay»"
## Выполнил: Киселев Иван БПИ217

### Установка и запуск программы
Данный репозиторий представляет из себя серверную часть приложения,
предполгается его запускать на VPS (виртуальном сервере),
в нашем же случае уже запущен, и к нему можно обратиться по адресу
http://134.0.119.220:8000/

В случае, если требуется установить данный програмный продукт.
Зайдя в данную директорию нужно прописать
```docker run --name=musicday_db -e POSTGRES_PASSWORD=password  -p 5432:5432 -d postgres```,
чтобы запустить в docker контейнер postgres, затем нужно запустить саму программу командой ```./build.exe```

Для запуска программы требуется наличие docker и postgreSQL (для помещения postgreSQL в docker надо прописть```docker pull postgres```).
### Описание репозитория
Данный репозиторий представляет из себя серверную часть приложения MusicDay


## Строение репозитория:

### Точка входа в программу и инициализация конфигураций
**[main.go](https://github.com/MusicDay-Inc/musicday_backend/blob/main/cmd/app/main.go)**
> В данном файле запускается вся прогамма:
> 
> Инициализируется logrus (библеотека логирования)
> 
> Подключаеся БД (инициализируется СУБД sqlx вместе с драйвером)
> 
> Инициализируются все три слоя чистой арихтектуры ссылаясь 
> друг на друга
```gotemplate
repos := repository.New(db)
services := service.NewService(repos)
handlers := transport.NewHandler(services)
```
> Запускается сервер на 8000 порту

### Коллекция моделей сущностей с которыми аппелирует программа, а также паттерны DTO DAO для них
**[сore/](https://github.com/MusicDay-Inc/musicday_backend/blob/main/core/)**
>Коллекция моделей сущностей с которыми аппелирует программа, 
> а также паттерны DTO DAO для них
```gotemplate
type User struct {
Id                 uuid.UUID
Gmail              string
Username           string
Nickname           string
IsRegistered       bool
HasProfilePic      bool
SubscriberAmount   int32
SubscriptionAmount int32
}

func (u *User) ToDTO() (user UserDTO) {
user.Id = u.Id
user.Gmail = u.Gmail
user.Nickname = u.Nickname
user.Username = u.Username
user.IsRegistered = u.IsRegistered
user.HasProfilePic = u.HasProfilePic
user.SubscriberAmount = u.SubscriberAmount
user.SubscriptionAmount = u.SubscriptionAmount
return
}
```

### Транспортный слой програмым
**[transport/](https://github.com/MusicDay-Inc/musicday_backend/blob/main/transport/)**
> Код (логика) связанная с принямием http запросов по определенным эндпоинтами (хэндлеры), 
> а также обращение в слой бизнес логики (/service). Например:
```go
func (h *Handler) start(c *gin.Context) {
var t idToken
if !bindRequestBody(c, &t) {
return
}

gmail, err := service.GetGmail(t.IdToken)
if err != nil {
newErrorResponse(c, http.StatusBadRequest, core.CodeTokenInvalid, 
	err.Error())
logrus.Info("invalid Google token")
return
}

jwt, err := h.services.GetJWT(gmail)
if err != nil {
newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, 
	core.ErrInternal.Error())
logrus.Errorf("while generating JWT" + err.Error())
return
}

c.JSON(http.StatusOK, jwt.ToResponse())
}
```
> А также там gin роутер, который распределяет запросы. Его часть:
```go
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/start", h.start)
		auth.POST("/sign_up", h.signUp)
	}
```

### Сервисный слой (слой бизнес-логики)
**[service/](https://github.com/MusicDay-Inc/musicday_backend/blob/main/service/)**
> Код связанный с основной логикой обработки различных данных, 
> ее проверка, провекра относительно других сущностей, аккомуляция данных и другое
>
> Данный слой обращается к слою репозитория (/repository/)


### Слой репозитория (слой бращения к хранилищу данных)
**[repository/](https://github.com/MusicDay-Inc/musicday_backend/blob/main/repository/)**
> Код связанный с получаение изменением и записью данных в БД (postgreSQL)
> Данный слой использует библиотеку sqlx и обращения к БД SQL запросами:
> 
```go
func (r ReviewRepository) GetReviewsOfUserSubscriptions(clientId uuid.UUID, limit int, offset int) (reviews []core.ReviewDAO, err error) {
q := `
SELECT id, user_id, is_song_reviewed, release_id, published_at, score, review_text
FROM reviews
INNER JOIN subscriptions on (reviews.user_id = subscriptions.subscription_id AND
subscriptions.subscriber_id = $1)
ORDER BY published_at DESC
LIMIT $2 OFFSET $3
`
logrus.Trace(formatQuery(q))
err = r.db.Select(&reviews, q, clientId, limit, offset)
if err != nil {
return reviews, err
}
return
}
```

## Методы различных слоев
### Транспортный
``` go
package transport

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

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
```
### Сервисный
```go
package service

type Token interface {
	GetJWT(gmail string) (core.JWT, error)
	ParseToken(token string) (uuid.UUID, bool, error)
	GenerateJWT(userId uuid.UUID, registered bool) (core.JWT, error)
}

type User interface {
	RegisterUser(userId uuid.UUID, user core.User) (core.User, error)
	Subscribe(clientId uuid.UUID, userId uuid.UUID) (core.UserDTO, error)
	ChangeUsername(clientId uuid.UUID, username string) (core.User, error)
	ChangeNickname(clientId uuid.UUID, nickname string) (core.User, error)
	SearchUsers(query string, clientId uuid.UUID, limit int, offset int) ([]core.UserDTO, error)
	Exists(id uuid.UUID) bool
	GetById(id uuid.UUID) (core.UserDTO, error)
	SubscriptionExists(clientId uuid.UUID, userId uuid.UUID) bool
	Unsubscribe(clientId uuid.UUID, userId uuid.UUID) (core.UserDTO, error)
	GetSubscribers(userId uuid.UUID, limit int, offset int) ([]core.UserDTO, error)
	GetSubscriptions(userId uuid.UUID, limit int, offset int) ([]core.UserDTO, error)
	GetBio(userId uuid.UUID) (string, error)
	CreateBio(clientId uuid.UUID, bio string) (string, error)
	UploadAvatar(clientId uuid.UUID) (core.UserDTO, error)
}

type Song interface {
	GetById(songId uuid.UUID) (core.Song, error)
	SearchSongsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) ([]core.SongWithReviewDTO, error)
}

type Album interface {
	GetById(songId uuid.UUID) (core.Album, error)
	GetSongsFromAlbum(id uuid.UUID) ([]core.Song, error)
	SearchAlbumsWithReview(query string, userId uuid.UUID, limit int, offset int) ([]core.AlbumWithReviewDTO, error)
	GetCoverId(srcId uuid.UUID) (uuid.UUID, error)
}

type Review interface {
	GetReviewToRelease(releaseId uuid.UUID, userId uuid.UUID) (core.Review, error)
	PostReview(review core.Review) (core.ReviewDTO, error)
	GetSubscriptionReviews(releaseId uuid.UUID, clientId uuid.UUID, limit int, offset int) ([]core.ReviewOfUserDTO, error)
	DeleteReviewFromUser(userId uuid.UUID, reviewId uuid.UUID) (core.ReviewDTO, error)
	GetAllUserReviews(userId uuid.UUID, limit int, offset int) ([]core.ReviewDTO, error)
	GetSongReviewsOfUser(userId uuid.UUID, limit int, offset int) ([]core.ReviewDTO, error)
	GetAlbumReviewsOfUser(userId uuid.UUID, limit int, offset int) ([]core.ReviewDTO, error)
	GetReviewsOfUserSubscriptions(clientId uuid.UUID, limit int, offset int) ([]core.ReviewOfUserDTO, error)
	CountSongReviewsOf(userId uuid.UUID) (int32, error)
	CountAlbumReviewsOf(userId uuid.UUID) (int32, error)
}
```

### Репозитория
```go
package repository

type User interface {
// Create returns id of new user, and changes his id
Create(gmail string) (uuid.UUID, error)
Exists(gmail string) bool
GetById(userId uuid.UUID) (user core.User, err error)
GetByUsername(username string) (core.UserDAO, error)
GetByGmail(gmail string) (user core.User, err error)
Register(u core.User) (user core.UserDAO, err error)
ChangeUsername(id uuid.UUID, username string) (user core.UserDAO, err error)
ChangeNickname(id uuid.UUID, nickname string) (user core.UserDAO, err error)
InstallPicture(id uuid.UUID) (user core.UserDAO, err error)
Subscribe(clientId uuid.UUID, userId uuid.UUID) (core.User, error)
SearchUsers(query string, clientId uuid.UUID, limit int, offset int) ([]core.UserDAO, error)
ExistsWithId(id uuid.UUID) bool
IsSubscriptionExists(clientId uuid.UUID, userId uuid.UUID) bool
Unsubscribe(clientId uuid.UUID, userId uuid.UUID) (core.User, error)
GetSubscribers(userId uuid.UUID, limit int, offset int) ([]core.UserDAO, error)
GetSubscriptionsOf(userId uuid.UUID, limit int, offset int) ([]core.UserDAO, error)
GetBio(userId uuid.UUID) (string, error)
CreateBio(userId uuid.UUID, bio string) (string, error)
}

type Song interface {
GetById(songId uuid.UUID) (core.SongDAO, error)
SearchSongsWithReview(searchReq string, userId uuid.UUID, limit int, offset int) ([]core.SongWithReviewDAO, error)
}

type Album interface {
GetById(album uuid.UUID) (core.AlbumDAO, error)
GetSongsFromAlbum(id uuid.UUID) ([]core.SongDAO, error)
SearchAlbumsWithReview(query string, userId uuid.UUID, limit int, offset int) ([]core.AlbumWithReviewDAO, error)
GetContainingSong(songId uuid.UUID) (uuid.UUID, error)
}

type Review interface {
GetById(id uuid.UUID) (core.ReviewDAO, error)
GetReviewToRelease(releaseId uuid.UUID, userId uuid.UUID) (core.ReviewDAO, error)
InsertReview(review core.Review) (core.ReviewDAO, error)
ExistsToRelease(userId uuid.UUID, releaseId uuid.UUID) (bool, error)
ExistsFromUser(userId uuid.UUID, releaseId uuid.UUID) (bool, error)
UpdateReview(review core.Review) (core.ReviewDAO, error)
GetSubscriptionReviews(releaseId uuid.UUID, clientId uuid.UUID) ([]core.ReviewDAO, error)
Delete(id uuid.UUID) error
GetSongReviewsFromUser(userId uuid.UUID, limit int, offset int) ([]core.ReviewDAO, error)
GetAlbumReviewsFromUser(userId uuid.UUID, limit int, offset int) ([]core.ReviewDAO, error)
GetReviewsFromUser(userId uuid.UUID, limit int, offset int) ([]core.ReviewDAO, error)
GetReviewsOfUserSubscriptions(clientId uuid.UUID, limit int, offset int) ([]core.ReviewDAO, error)
CountReviewsOfUser(userId uuid.UUID, isToSongs bool) (int32, error)
}
```
