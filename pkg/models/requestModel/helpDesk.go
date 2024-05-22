package requestmodel

type HelpDeskRequest struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Subject     string `json:"subject" validate:"required"`
	Message     string `json:"message" validate:"required"`
}

type HelpDeskAnswer struct {
	Answer string `json:"answer" validate:"required"`
}
