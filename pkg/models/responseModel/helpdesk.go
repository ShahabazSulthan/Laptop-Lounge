package responsemodel

import "time"

type HelpDeskResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Subject     string    `json:"subject"`
	Message     string    `json:"message"`
	Answer      string    `json:"answer,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
