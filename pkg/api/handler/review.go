package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	useCase interfaceUseCase.IReviewUseCase
}

func NewReviewHandler(useCase interfaceUseCase.IReviewUseCase) *ReviewHandler {
	return &ReviewHandler{useCase: useCase}
}

func (r *ReviewHandler) AddReview(c *gin.Context) {
	var reviewRequest *requestmodel.ReviewRequest

	if err := c.BindJSON(&reviewRequest); err != nil {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	data, err := helper.Validation(reviewRequest)
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalResult := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	productID := c.Param("productID")

	err = r.useCase.AddReview(productID, userID, reviewRequest.Rating, reviewRequest.Comment)
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
	} else {
		finalResult := response.Responses(http.StatusCreated, "", reviewRequest, nil) // Use HTTP 201 for resource creation
		c.JSON(http.StatusCreated, finalResult)
	}
}

func (r *ReviewHandler) GetReviewsByProductID(c *gin.Context) {
	productID := c.Param("productID")

	reviews, err := r.useCase.GetReviewsByProductID(productID)
	if err != nil {
		finalResult := response.Responses(http.StatusInternalServerError, "", nil, err.Error())
		c.JSON(http.StatusInternalServerError, finalResult)
		return
	}

	finalResult := response.Responses(http.StatusOK, "", reviews, nil)
	c.JSON(http.StatusOK, finalResult)
}

func (r *ReviewHandler) DeleteReviewByID(c *gin.Context) {
	reviewID := c.Param("reviewID")

	err := r.useCase.DeleteReviewByID(reviewID)
	if err != nil {
		finalResult := response.Responses(http.StatusInternalServerError, "", nil, err.Error())
		c.JSON(http.StatusInternalServerError, finalResult)
		return
	}

	finalResult := response.Responses(http.StatusOK, "Review deleted successfully", nil, "")
	c.JSON(http.StatusOK, finalResult)
}

func (r *ReviewHandler) GetAverageRating(c *gin.Context) {
	productID := c.Param("productID")

	avgRating, err := r.useCase.GetAverageRating(productID)
	if err != nil {
		finalResult := response.Responses(http.StatusInternalServerError, "", nil, err.Error())
		c.JSON(http.StatusInternalServerError, finalResult)
		return
	}

	responseData := map[string]float64{
		"average_rating": avgRating,
	}

	finalResult := response.Responses(http.StatusOK, "", responseData, nil)
	c.JSON(http.StatusOK, finalResult)
}

