package service

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"server/internal/core"
	"server/internal/repository"
)

type ReviewService struct {
	rev   repository.Review
	song  repository.Song
	album repository.Album
	user  repository.User
}

func (s *ReviewService) CountSongReviewsOf(userId uuid.UUID) (int32, error) {
	amount, err := s.rev.CountReviewsOfUser(userId, true)
	return amount, err
}

func (s *ReviewService) CountAlbumReviewsOf(userId uuid.UUID) (int32, error) {
	amount, err := s.rev.CountReviewsOfUser(userId, false)
	return amount, err
}

func (s *ReviewService) GetReviewsOfUserSubscriptions(clientId uuid.UUID, limit int, offset int) (res []core.ReviewOfUserDTO, err error) {
	reviews, err := s.rev.GetReviewsOfUserSubscriptions(clientId, limit, offset)
	if err != nil {
		return
	}
	res = make([]core.ReviewOfUserDTO, len(reviews))
	for i, review := range reviews {
		rDomain := review.ToDomain()
		if rDomain.IsSongReviewed {
			song, errSong := s.song.GetById(rDomain.ReleaseId)
			sD := song.ToDomain()
			if errSong != nil {
				logrus.Errorf("Unexpected error %v", errSong)
				return make([]core.ReviewOfUserDTO, 0), errSong
			}
			userItem, errUser := s.user.GetById(rDomain.UserId)
			if errUser != nil {
				return make([]core.ReviewOfUserDTO, 0), errUser
			}
			rDTO := rDomain.ToUserDTO(userItem)
			rDTO.Song = sD.ToDTO()
			res[i] = rDTO
		} else {
			album, errAlbum := s.album.GetById(rDomain.ReleaseId)
			aD := album.ToDomain()
			if errAlbum != nil {
				logrus.Errorf("Unexpected error %v", errAlbum)
				return make([]core.ReviewOfUserDTO, 0), errAlbum
			}
			userItem, errUser := s.user.GetById(rDomain.UserId)
			if errUser != nil {
				return make([]core.ReviewOfUserDTO, 0), errUser
			}
			rDTO := rDomain.ToUserDTO(userItem)
			rDTO.Album = aD.ToDTO()
			res[i] = rDTO
		}

	}
	return res, nil
}

func (s *ReviewService) GetAlbumReviewsOfUser(userId uuid.UUID, limit int, offset int) (res []core.ReviewDTO, err error) {
	reviews, err := s.rev.GetAlbumReviewsFromUser(userId, limit, offset)
	if err != nil {
		return
	}
	res = make([]core.ReviewDTO, len(reviews))
	for i, review := range reviews {
		rDomain := review.ToDomain()
		album, errSong := s.album.GetById(rDomain.ReleaseId)
		if errSong != nil {
			logrus.Errorf("Unexpected error %v", errSong)
		}
		res[i] = rDomain.ToAlbumDTO(album.ToDomain())
	}
	return res, nil
}

func (s *ReviewService) GetAllUserReviews(userId uuid.UUID, limit int, offset int) (res []core.ReviewDTO, err error) {
	reviews, err := s.rev.GetReviewsFromUser(userId, limit, offset)
	if err != nil {
		return
	}
	res = make([]core.ReviewDTO, len(reviews))
	for i, review := range reviews {
		rDomain := review.ToDomain()
		if rDomain.IsSongReviewed {
			song, errSong := s.song.GetById(rDomain.ReleaseId)
			if errSong != nil {
				logrus.Errorf("Unexpected error %v", errSong)
			}
			res[i] = rDomain.ToSongDTO(song.ToDomain())
		} else {
			album, errSong := s.album.GetById(rDomain.ReleaseId)
			if errSong != nil {
				logrus.Errorf("Unexpected error %v", errSong)
			}
			res[i] = rDomain.ToAlbumDTO(album.ToDomain())
		}
	}
	return res, nil
}

func (s *ReviewService) GetSongReviewsOfUser(userId uuid.UUID, limit int, offset int) (res []core.ReviewDTO, err error) {
	reviews, err := s.rev.GetSongReviewsFromUser(userId, limit, offset)
	if err != nil {
		return
	}
	res = make([]core.ReviewDTO, len(reviews))
	for i, review := range reviews {
		rDomain := review.ToDomain()
		song, errSong := s.song.GetById(rDomain.ReleaseId)
		if errSong != nil {
			logrus.Errorf("Unexpected error %v", errSong)
		}
		res[i] = rDomain.ToSongDTO(song.ToDomain())
	}
	return res, nil
}

func (s *ReviewService) DeleteReviewFromUser(userId uuid.UUID, reviewId uuid.UUID) (core.ReviewDTO, error) {
	exists, err := s.rev.ExistsFromUser(userId, reviewId)
	if err != nil {
		return core.ReviewDTO{}, core.ErrInternal
	}
	if !exists {
		return core.ReviewDTO{}, core.ErrNotFound
	}
	r, err := s.rev.GetById(reviewId)
	if err != nil {
		return core.ReviewDTO{}, core.ErrNotFound
	}
	err = s.rev.Delete(reviewId)
	if err != nil {
		return core.ReviewDTO{}, core.ErrInternal
	}
	if r.IsSongReviewed {
		song, errS := s.song.GetById(r.ReleaseId)
		sD := song.ToDomain()
		if errS != nil {
			return core.ReviewDTO{}, core.ErrInternal
		}
		return core.ReviewDTO{Song: sD.ToDTO()}, nil
	} else {
		a, errA := s.album.GetById(r.ReleaseId)
		aD := a.ToDomain()
		if errA != nil {
			return core.ReviewDTO{}, core.ErrInternal
		}
		return core.ReviewDTO{Album: aD.ToDTO()}, nil
	}
}

func (s *ReviewService) GetSubscriptionReviews(releaseId uuid.UUID, clientId uuid.UUID, limit int, offset int) (res []core.ReviewOfUserDTO, err error) {
	reviews, err := s.rev.GetSubscriptionReviews(releaseId, clientId)
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
	exists, err := s.rev.ExistsToRelease(reviewReq.UserId, reviewReq.ReleaseId)
	if err != nil {
		return core.ReviewDTO{}, core.ErrInternal
	}
	if exists {
		updateRes, errUpdate := s.rev.UpdateReview(reviewReq)
		if errUpdate != nil {
			return core.ReviewDTO{}, core.ErrInternal
		}
		resDomain := updateRes.ToDomain()
		if resDomain.IsSongReviewed {
			return resDomain.ToSongDTO(songFromRequest.ToDomain()), nil
		}
		return resDomain.ToAlbumDTO(albumFromRequest.ToDomain()), nil
	}
	insertRes, err := s.rev.InsertReview(reviewReq)
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
	review, err := s.rev.GetReviewToRelease(releaseId, userId)
	if err != nil {
		return core.Review{}, err
	}
	return review.ToDomain(), nil
}

func NewReviewService(r repository.Review, song repository.Song, album repository.Album, user repository.User) *ReviewService {
	return &ReviewService{
		rev:   r,
		song:  song,
		album: album,
		user:  user,
	}
}
