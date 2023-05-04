package service

import (
	"github.com/google/uuid"
	"server/internal/core"
	"server/internal/repository"
)

type ReviewService struct {
	r repository.Review
}

func (s *ReviewService) GetReviewToRelease(releaseId uuid.UUID, userId uuid.UUID) (core.Review, error) {
	review, err := s.r.GetReviewToRelease(releaseId, userId)
	if err != nil {
		return core.Review{}, err
	}
	return review.ToDomain(), nil
}

func NewReviewService(r repository.Review) *ReviewService {
	return &ReviewService{r: r}
}
