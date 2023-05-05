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
	SubscriberAmount   int
	SubscriptionAmount int
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

type Review struct {
	Id             uuid.UUID
	UserId         uuid.UUID
	IsSongReviewed bool
	ReleaseId      uuid.UUID
	PublishedAt    time.Time
	Score          uint8
	Text           string
}

func (r *Review) ValidateScore() bool {
	return r.Score > 0 && r.Score <= 10
}

func (r *Review) ToEmptyDTO() (review ReviewDTO) {
	review.Id = r.Id
	review.UserId = r.UserId
	review.PublishedAt = r.PublishedAt
	review.Score = r.Score
	review.Text = r.Text
	return
}
func (r *Review) ToSongDTO(song Song) (review ReviewDTO) {
	review.Id = r.Id
	review.UserId = r.UserId
	review.IsSongReviewed = r.IsSongReviewed
	review.Song = song
	review.PublishedAt = r.PublishedAt
	review.Score = r.Score
	review.Text = r.Text
	return
}
func (r *Review) ToAlbumDTO(album Album) (review ReviewDTO) {
	review.Id = r.Id
	review.UserId = r.UserId
	review.IsSongReviewed = r.IsSongReviewed
	review.Album = album
	review.PublishedAt = r.PublishedAt
	review.Score = r.Score
	review.Text = r.Text
	return
}
