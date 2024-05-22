package interfaces

import responsemodel "Laptop_Lounge/pkg/models/responseModel"

type IreviewRepo interface {
	AddReview(string, string, int, string) error
	GetReviewsByProductID(string) ([]responsemodel.ReviewResponse, error)
	DeleteReviewByID(string) error
	GetAverageRating(string) (float64, error)
}
