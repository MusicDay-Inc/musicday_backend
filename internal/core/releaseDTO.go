package core

import (
	"github.com/google/uuid"
	"time"
)

type SongDTO struct {
	Id       uuid.UUID     `json:"id,omitempty"`
	Name     string        `json:"name,omitempty"`
	Author   string        `json:"author,omitempty"`
	Date     time.Time     `json:"date,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
}
type SongWithReviewDTO struct {
	SongDTO   `json:"song,omitempty"`
	ReviewDTO `json:"review,omitempty"`
}

type AlbumWithReviewDTO struct {
	AlbumDTO  `json:"album,omitempty"`
	ReviewDTO `json:"review,omitempty"`
}

type AlbumDTO struct {
	Id         uuid.UUID     `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	Author     string        `json:"author,omitempty"`
	Date       time.Time     `json:"date,omitempty"`
	SongAmount int           `json:"song_amount,omitempty"`
	Duration   time.Duration `json:"duration,omitempty"`
	Songs      []SongDTO     `json:"songs,omitempty"`
}

type SubscriptionReviews struct {
	Reviews   []ReviewOfUserDTO `json:"user_reviews"`
	MeanScore float32           `json:"mean_score"`
}
type SubscriptionIdListReviews struct {
	Reviews   []uuid.UUID `json:"user_reviews_id,omitempty"`
	MeanScore float32     `json:"mean_score,omitempty"`
}

//type SubscriptionReviews struct {
//	Reviews   []ReviewOfUserDTO `json:"user_reviews,omitempty"`
//	MeanScore float32           `json:"mean_score,omitempty"`
//}

//type ReleaseDTO struct {
//	Id         uuid.UUID     `json:"id,omitempty"`
//	IsAlbum    bool          `json:"is_album,omitempty"`
//	Name       string        `json:"name,omitempty"`
//	Author     string        `json:"author,omitempty"`
//	Date       time.Time     `json:"date,omitempty"`
//	Duration   time.Duration `json:"duration,omitempty"`
//	SongAmount int           `json:"song_amount,omitempty"`
//	Songs      []SongDTO     `json:"songs,omitempty"`
//}
