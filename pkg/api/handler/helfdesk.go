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
