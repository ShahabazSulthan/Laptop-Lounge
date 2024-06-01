package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	useCase interfaceUseCase.IReviewUseCase
}

func NewReviewHandler(useCase interfaceUseCase.IReviewUseCase) *ReviewHandler {
	return &ReviewHandler{useCase: useCase}
}

// @Summary Add Review
// @Description Add a review for a product.
// @Tags Reviews
// @Accept json
// @Produce json
// @Security     BearerTokenAuth
// @Security     RefreshtokenAuth
// @Param productID path string true "ID of the product"
// @Param reviewRequest body requestmodel.ReviewRequest true "Review details"
// @Success 201 {object} response.Response "Review added successfully"
// @Failure 400 {object} response.Response "Bad Request. Invalid input."
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /review/{productID} [post]
// @example reviewRequest Example: { "rating": 3, "comment": "its just Okay" }
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

// @Summary Get Reviews by Product ID
// @Description Retrieve reviews for a product by its ID.
// @Tags Reviews
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security RefreshtokenAuth
// @Param productID path string true "ID of the product"
// @Success 200 {object} response.Response "Successfully retrieved reviews"
// @Failure 500 {object} response.Response "Internal Server Error"
// @Router /review/ [get]
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

// @Summary Delete Review by ID
// @Description Delete a review by its ID.
// @Tags Reviews
// @Accept json
// @Produce json
// @Security     BearerTokenAuth
// @Security     RefreshtokenAuth
// @Param reviewID path string true "ID of the review"
// @Success 200 {object} response.Response "Review deleted successfully"
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /review/{reviewID} [delete]
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

// @Summary Get Average Rating
// @Description Get the average rating for a product.
// @Tags Reviews
// @Accept json
// @Produce json
// @Security     BearerTokenAuth
// @Security     RefreshtokenAuth
// @Param productID path string true "ID of the product"
// @Success 200 {object} response.Response "Successfully retrieved average rating"
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /review/{productID} [get]
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

// Other code from your ReviewHandler...

// NewReviewHandler and other methods remain unchanged

// @Summary Get Log File
// @Description Retrieve and display the content of the log file.
// @Tags Logs
// @Accept json
// @Produce plain
// @Success 200 {string} string "Log file content"
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /log [get]
func (r *ReviewHandler) GetLogFile(c *gin.Context) {
	logData, err := ioutil.ReadFile("app.log")
	if err != nil {
		finalResult := response.Responses(http.StatusInternalServerError, "", nil, "Failed to read log file")
		c.JSON(http.StatusInternalServerError, finalResult)
		return
	}

	// Set plain text content type
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, "%s", logData)
}


func (r *ReviewHandler) GetExcelReport(c *gin.Context) {
	reportPath := "C:\\Users\\shaha\\OneDrive\\Desktop\\GO-Workplace\\First Project\\Laptop_Lounge\\Report\\salesReport.pdf" // Path to the Excel report file

	c.File(reportPath)
}


func (r *ReviewHandler) GetInvoice(c *gin.Context) {
	invoicePath := "C:\\Users\\shaha\\OneDrive\\Desktop\\GO-Workplace\\First Project\\Laptop_Lounge\\Report\\invoice.pdf" // Path to the invoice file

	c.File(invoicePath)
}
