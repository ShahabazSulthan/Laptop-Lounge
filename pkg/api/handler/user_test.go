package handler

import (
	"Laptop_Lounge/pkg/mock/mockUseCase"
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserSignUp(t *testing.T) {
	// Define test cases as a map with descriptive names
	testCase := map[string]struct {
		input         requestmodel.UserDetails
		buildstub     func(useCaseMock *mockUseCase.MockIuserUseCase, signupData requestmodel.UserDetails)
		checkResponse func(t *testing.T, responserecorder *httptest.ResponseRecorder)
	}{
		"success": {
			// Define the input for the success case
			input: requestmodel.UserDetails{
				Name:            "Shahabaz",
				Email:           "shahabaz@gmail.com",
				Phone:           "9496703880",
				Password:        "dskjb986",
				ConfirmPassword: "dskjb986",
			},
			// Stub the behavior of the mock use case for the success case
			buildstub: func(useCaseMock *mockUseCase.MockIuserUseCase, signupData requestmodel.UserDetails) {
				// Validate the input data
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("Validation Failed")
				}
				// Expect the UserSignup function to be called once and return a successful response
				useCaseMock.EXPECT().UserSignup(&signupData).Times(1).Return(&responsemodel.SignupData{
					Name:        "Shahabaz",
					Email:       "shahabaz@gmail.com",
					Phone:       "9496703880",
					ID:          "21",
					ReferalCode: "aabb2",
				}, nil)
			},
			// Check the response for the success case
			checkResponse: func(t *testing.T, responserecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responserecorder.Code)
			},
		},
		"bad request": {
			// Define the input for the bad request case
			input: requestmodel.UserDetails{
				Name:            "Shahabaz",
				Email:           "shahabaz@gmail.com",
				Phone:           "9496703880",
				Password:        "dskjb986",
				ConfirmPassword: "dskjb986",
			},
			// Stub the behavior of the mock use case for the bad request case
			buildstub: func(useCaseMock *mockUseCase.MockIuserUseCase, signupData requestmodel.UserDetails) {
				// Validate the input data
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("Validation Failed")
				}
				// Expect the UserSignup function to be called once and return an error
				useCaseMock.EXPECT().UserSignup(&signupData).Times(1).Return(nil, errors.New("user request not satisfy credential"))
			},
			// Check the response for the bad request case
			checkResponse: func(t *testing.T, responserecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responserecorder.Code)
			},
		},
	}

	// Loop over each test case
	for testname, test := range testCase {
		// Capture the range variable to avoid concurrency issues
		test := test
		t.Run(testname, func(t *testing.T) {
			t.Parallel() // Run tests in parallel

			// Create a new mock controller and mock use case
			ctrl := gomock.NewController(t)
			mockUseCase := mockUseCase.NewMockIuserUseCase(ctrl)

			// Build the stub for the current test case
			test.buildstub(mockUseCase, test.input)

			// Create a new user handler with the mock use case
			userHandler := NewUserHandler(mockUseCase)

			// Set up the Gin server and define the signup route
			server := gin.Default()
			server.POST("/signup", userHandler.UserSignup)

			// Marshal the input data to JSON
			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			// Create a new HTTP POST request with the JSON body
			mockRequest, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)

			// Create a new response recorder to capture the response
			responseRecord := httptest.NewRecorder()
			server.ServeHTTP(responseRecord, mockRequest)

			// Check the response using the checkResponse function for the current test case
			test.checkResponse(t, responseRecord)
		})
	}
}
