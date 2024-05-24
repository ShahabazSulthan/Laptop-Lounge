package usecase

import (
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
)

type ReviewUseCase struct {
	repo interfaces.IreviewRepo
}

func NewReviewtUseCase(repository interfaces.IreviewRepo) interfaceUseCase.IReviewUseCase {
	return &ReviewUseCase{repo: repository}
}

func (r *ReviewUseCase) AddReview(productID, userID string, rating int, comment string) error {
	err := r.repo.AddReview(productID, userID, rating, comment)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReviewUseCase) GetReviewsByProductID(productID string) ([]responsemodel.ReviewResponse, error) {
	reviews, err := r.repo.GetReviewsByProductID(productID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewUseCase) DeleteReviewByID(reviewID string) error {
	err := r.repo.DeleteReviewByID(reviewID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReviewUseCase) GetAverageRating(productID string) (float64, error) {
	return r.repo.GetAverageRating(productID)
}
