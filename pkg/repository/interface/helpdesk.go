package interfaces

import responsemodel "Laptop_Lounge/pkg/models/responseModel"

type IhelpDeskRepo interface {
	CreateRequest(string, string, string, string) error
	UpdateAnswer(uint, string) error
	GetRepliedRequests() ([]responsemodel.HelpDeskResponse, error)
	GetUnrepliedRequests() ([]responsemodel.HelpDeskResponse, error)
}
