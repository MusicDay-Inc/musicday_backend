package core

import (
	"github.com/google/uuid"
	"time"
)

type SongPayload struct {
	Id       uuid.UUID     `json:"id,omitempty"`
	Name     string        `json:"name,omitempty"`
	Author   string        `json:"author,omitempty"`
	Date     time.Time     `json:"date,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
}
type SongWithReviewPayload struct {
	SongPayload   `json:"song,omitempty"`
	ReviewPayload `json:"review,omitempty"`
}

type AlbumWithReviewPayload struct {
	AlbumPayload  `json:"album,omitempty"`
	ReviewPayload `json:"review,omitempty"`
}

type AlbumPayload struct {
	Id         uuid.UUID     `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	Author     string        `json:"author,omitempty"`
	Date       time.Time     `json:"date,omitempty"`
	SongAmount int           `json:"song_amount,omitempty"`
	Duration   time.Duration `json:"duration,omitempty"`
	Songs      []SongPayload `json:"songs,omitempty"`
}

type SubscriptionReviews struct {
	Reviews   []ReviewOfUserPayload `json:"user_reviews"`
	MeanScore float32               `json:"mean_score"`
}
type SubscriptionIdListReviews struct {
	Reviews   []uuid.UUID `json:"user_reviews_id,omitempty"`
	MeanScore float32     `json:"mean_score,omitempty"`
}

//type SubscriptionReviews struct {
//	Reviews   []ReviewOfUserPayload `json:"user_reviews,omitempty"`
//	MeanScore float32           `json:"mean_score,omitempty"`
//}
