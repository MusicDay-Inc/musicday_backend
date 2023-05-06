package service

import (
	"github.com/google/uuid"
	"server/internal/core"
	"server/internal/repository"
)

type ReviewService struct {
	r     repository.Review
	song  repository.Song
	album repository.Album
	user  repository.User
}

func (s *ReviewService) GetSubscriptionReviews(releaseId uuid.UUID, userId uuid.UUID, limit int, offset int) (res []core.ReviewOfUserDTO, err error) {
	reviews, err := s.r.GetSubscriptionReviews(releaseId, userId, limit, offset)
	if err != nil {
		return
	}
	res = make([]core.ReviewOfUserDTO, len(reviews))
	for i, review := range reviews {
		rDomain := review.ToDomain()
		userItem, errUser := s.user.GetById(rDomain.UserId)
		if errUser != nil {
			return make([]core.ReviewOfUserDTO, 0), errUser
		}
		res[i] = rDomain.ToUserDTO(userItem)
	}
	return res, nil
}

func (s *ReviewService) PostReview(reviewReq core.Review) (core.ReviewDTO, error) {
	if !reviewReq.ValidateScore() {
		return core.ReviewDTO{}, core.ErrIncorrectBody
	}
	albumFromRequest, err := s.album.GetById(reviewReq.ReleaseId)
	if err != nil {
		reviewReq.IsSongReviewed = true
	}
	songFromRequest, err := s.song.GetById(reviewReq.ReleaseId)
	if err != nil {
		if reviewReq.IsSongReviewed {
			return core.ReviewDTO{}, core.ErrNotFound
		}
	}
	exists, err := s.r.Exists(reviewReq.UserId, reviewReq.ReleaseId)
	if err != nil {
		return core.ReviewDTO{}, core.ErrInternal
	}
	if exists {
		updateRes, errUpdate := s.r.UpdateReview(reviewReq)
		if errUpdate != nil {
			return core.ReviewDTO{}, core.ErrInternal
		}
		resDomain := updateRes.ToDomain()
		if resDomain.IsSongReviewed {
			return resDomain.ToSongDTO(songFromRequest.ToDomain()), nil
		}
		return resDomain.ToAlbumDTO(albumFromRequest.ToDomain()), nil
	}
	insertRes, err := s.r.InsertReview(reviewReq)
	if err != nil {
		return core.ReviewDTO{}, err
	}
	res := insertRes.ToDomain()
	if reviewReq.IsSongReviewed {
		return res.ToSongDTO(songFromRequest.ToDomain()), nil
	}
	return res.ToAlbumDTO(albumFromRequest.ToDomain()), nil
}

func (s *ReviewService) GetReviewToRelease(releaseId uuid.UUID, userId uuid.UUID) (core.Review, error) {
	review, err := s.r.GetReviewToRelease(releaseId, userId)
	if err != nil {
		return core.Review{}, err
	}
	return review.ToDomain(), nil
}

func NewReviewService(r repository.Review, song repository.Song, album repository.Album, user repository.User) *ReviewService {
	return &ReviewService{
		r:     r,
		song:  song,
		album: album,
		user:  user,
	}
}
