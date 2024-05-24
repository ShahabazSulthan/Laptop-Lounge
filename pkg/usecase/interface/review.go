package interfaceUseCase

import responsemodel "Laptop_Lounge/pkg/models/responseModel"

type IReviewUseCase interface {
	AddReview(string, string, int, string) error
	GetReviewsByProductID(string) ([]responsemodel.ReviewResponse, error)
	DeleteReviewByID(string) error
	GetAverageRating(string) (float64, error)
}
