package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HelpDeskHandler struct {
	useCase interfaceUseCase.IHelpDeskUseCase
}

func NewHelpDeskHandler(useCase interfaceUseCase.IHelpDeskUseCase) *HelpDeskHandler {
	return &HelpDeskHandler{useCase: useCase}
}

// @Summary Create Help Desk Request
// @Description Create a new help desk request.
// @Tags Help Desk
// @Accept json
// @Produce json
// @Param request body requestmodel.HelpDeskRequest true "Help desk request details"
// @Success 201 {object} response.Response "Request created successfully"
// @Failure 400 {object} response.Response "Bad Request. Invalid input."
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /helpdesk/request [post]
func (h *HelpDeskHandler) CreateRequest(c *gin.Context) {
	var req requestmodel.HelpDeskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.CreateRequest(req.Name, req.PhoneNumber, req.Subject, req.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Request created successfully"})
}

// @Summary Update Help Desk Answer
// @Description Update the answer for a help desk request.
// @Tags Help Desk
// @Accept json
// @Produce json
// @Param requestID path int true "ID of the request to update"
// @Param answer body requestmodel.HelpDeskAnswer true "Updated answer details"
// @Success 200 {object} response.Response "Answer updated successfully"
// @Failure 400 {object} response.Response "Bad Request. Invalid input."
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /helpdesk/request/:requestID [patch]
func (h *HelpDeskHandler) UpdateAnswer(c *gin.Context) {
	var ans requestmodel.HelpDeskAnswer
	if err := c.ShouldBindJSON(&ans); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestID := c.Param("requestID")
	id, err := strconv.Atoi(requestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request ID"})
		return
	}

	if err := h.useCase.UpdateAnswer(uint(id), ans.Answer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer updated successfully"})
}

// @Summary Get Replied Help Desk Requests
// @Description Retrieve all help desk requests that have been replied to.
// @Tags Help Desk
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "Successfully retrieved replied requests"
// @Failure 400 {object} response.Response "Bad Request. Unable to retrieve replied requests."
// @Router /helpdesk/replied [get]
func (h *HelpDeskHandler) GetRepliedRequests(c *gin.Context) {

	replay, err := h.useCase.GetRepliedRequests()
	if err != nil {
		// If there's an error, construct a response with a JSON error message and status 400 (Bad Request)
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return // Exit the function to avoid further processing
	}

	// If successful, construct the final response with the product data and status 200 (OK)
	finalResult := response.Responses(http.StatusOK, "", replay, nil)
	c.JSON(http.StatusOK, finalResult)
}

// @Summary Get Unreplied Help Desk Requests
// @Description Retrieve all help desk requests that are yet to be replied to.
// @Tags Help Desk
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "Successfully retrieved unreplied requests"
// @Failure 400 {object} response.Response "Bad Request"
// @Router /helpdesk/unreplied [get]
func (h *HelpDeskHandler) GetUnrepliedRequests(c *gin.Context) {
	unreplay, err := h.useCase.GetUnrepliedRequests()
	if err != nil {
		// If there's an error, construct a response with a JSON error message and status 400 (Bad Request)
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return // Exit the function to avoid further processing
	}

	// If successful, construct the final response with the product data and status 200 (OK)
	finalResult := response.Responses(http.StatusOK, "", unreplay, nil)
	c.JSON(http.StatusOK, finalResult)
}
