package domain

type status string

const (
	Active  status = "Active"
	Block   status = "Blocked"
	Delete  status = "Deleted"
	Pending status = "Pending"
)

type Users struct {
	ID          uint `gorm:"unique;not null;primary key"`
	Name        string
	Email       string `validate:"required,email"`
	Phone       string
	Password    string
	ReferalCode string 
	Status      status `status:"default:pending"`
}

type Address struct {
	ID          uint `gorm:"unique;not null;primaryKey"`
	Userid      uint
	User        Users `gorm:"foreignkey:Userid;association_foreignkey:ID"`
	FirstName   string
	LastName    string
	Street      string
	City        string
	State       string
	Pincode     string
	LandMark    string
	PhoneNumber string
	Status      status `gorm:"default:active"`
}
