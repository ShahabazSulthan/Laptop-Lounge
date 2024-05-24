package repository

import (
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"fmt"

	"gorm.io/gorm"
)

type HelpDeskRepo struct {
	DB *gorm.DB
}

func NewHelpDeskRepository(db *gorm.DB) interfaces.IhelpDeskRepo {
	return &HelpDeskRepo{DB: db}
}

func (r *HelpDeskRepo) CreateRequest(name, phoneNumber, subject, message string) error {
	sql := "INSERT INTO help_desks (name, phone_number, subject, message, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())"
	if err := r.DB.Exec(sql, name, phoneNumber, subject, message).Error; err != nil {
		return fmt.Errorf("encountered an issue while inserting into help desk: %w", err)
	}
	return nil
}

func (r *HelpDeskRepo) UpdateAnswer(requestID uint, answer string) error {
	sql := "UPDATE help_desks SET answer = ?, updated_at = NOW() WHERE id = ?"
	if err := r.DB.Exec(sql, answer, requestID).Error; err != nil {
		return fmt.Errorf("encountered an issue while updating the help desk answer: %w", err)
	}
	return nil
}

func (r *HelpDeskRepo) GetRepliedRequests() ([]responsemodel.HelpDeskResponse, error) {
	var repliedRequests []responsemodel.HelpDeskResponse
	query := "SELECT id, name, phone_number, subject, message, answer, created_at, updated_at FROM help_desks WHERE answer IS NOT NULL"
	if err := r.DB.Raw(query).Scan(&repliedRequests).Error; err != nil {
		return nil, fmt.Errorf("encountered an issue while retrieving replied help desk requests: %w", err)
	}
	return repliedRequests, nil
}

func (r *HelpDeskRepo) GetUnrepliedRequests() ([]responsemodel.HelpDeskResponse, error) {
	var unrepliedRequests []responsemodel.HelpDeskResponse
	query := "SELECT id, name, phone_number, subject, message, created_at, updated_at FROM help_desks WHERE answer IS NULL"
	if err := r.DB.Raw(query).Scan(&unrepliedRequests).Error; err != nil {
		return nil, fmt.Errorf("encountered an issue while retrieving unreplied help desk requests: %w", err)
	}
	return unrepliedRequests, nil
}
