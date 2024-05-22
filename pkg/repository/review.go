package repository

import (
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"errors"

	"gorm.io/gorm"
)

type ReviewRepo struct {
	DB *gorm.DB
}

func NewReviewRepository(db *gorm.DB) interfaces.IreviewRepo {
	return &ReviewRepo{DB: db}
}

func (r *ReviewRepo) AddReview(productID, userID string, rating int, comment string) error {
	query := "INSERT INTO reviews (product_id, user_id, rating, comment) VALUES (?, ?, ?, ?)"
	if err := r.DB.Exec(query, productID, userID, rating, comment).Error; err != nil {
		return errors.New("encountered an issue while inserting into reviews")
	}
	return nil
}

func (r *ReviewRepo) GetReviewsByProductID(productID string) ([]responsemodel.ReviewResponse, error) {
	var reviews []responsemodel.ReviewResponse
	if err := r.DB.Table("reviews").Find(&reviews).Error; err != nil {
		return nil, errors.New("failed to fetch reviews for product")
	}
	return reviews, nil
}

func (r *ReviewRepo) DeleteReviewByID(reviewID string) error {
	query := "DELETE FROM reviews WHERE id = ?"
	if err := r.DB.Exec(query, reviewID).Error; err != nil {
		return errors.New("failed to delete review")
	}
	return nil
}

func (r *ReviewRepo) GetAverageRating(productID string) (float64, error) {
	var avgRating float64
	err := r.DB.Table("reviews").Where("product_id = ?", productID).Select("AVG(rating)").Row().Scan(&avgRating)
	if err != nil {
		return 0, errors.New("failed to calculate average rating")
	}
	return avgRating, nil
}
