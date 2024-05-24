package domain

import "time"

type HelpDesk struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	PhoneNumber string `gorm:"type:varchar(20);not null"`
	Subject     string `gorm:"type:varchar(255);not null"`
	Message     string `gorm:"type:text;not null"`
	Answer      string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
