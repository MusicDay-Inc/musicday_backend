package core

import (
	"github.com/google/uuid"
	"time"
)

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

func (u *User) ToPayload() (user UserPayload) {
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

type Review struct {
	Id             uuid.UUID
	UserId         uuid.UUID
	IsSongReviewed bool
	ReleaseId      uuid.UUID
	PublishedAt    time.Time
	Score          int32
	Text           string
}

func (r *Review) ValidateScore() bool {
	return r.Score > 0 && r.Score <= 5
}

func (r *Review) ToEmptyPayload() (review ReviewPayload) {
	review.Id = r.Id
	review.UserId = r.UserId
	review.PublishedAt = r.PublishedAt
	review.Score = r.Score
	review.Text = r.Text
	return
}
func (r *Review) ToUserPayload(user User) (review ReviewOfUserPayload) {
	review.Id = r.Id
	review.User = user.ToPayload()
	review.PublishedAt = r.PublishedAt
	review.Score = r.Score
	review.Text = r.Text
	return
}
func (r *Review) ToSongPayload(song Song) (review ReviewPayload) {
	review.Id = r.Id
	review.UserId = r.UserId
	review.IsSongReviewed = r.IsSongReviewed
	review.Song = song.ToPayload()
	review.PublishedAt = r.PublishedAt
	review.Score = r.Score
	review.Text = r.Text
	return
}
func (r *Review) ToAlbumPayload(album Album) (review ReviewPayload) {
	review.Id = r.Id
	review.UserId = r.UserId
	review.IsSongReviewed = r.IsSongReviewed
	review.Album = album.ToPayload()
	review.PublishedAt = r.PublishedAt
	review.Score = r.Score
	review.Text = r.Text
	return
}
