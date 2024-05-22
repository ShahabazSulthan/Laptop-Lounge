package usecase

import (
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
)

type HelpDeskUseCase struct {
	repo interfaces.IhelpDeskRepo
}

func NewHelpDeskUseCase(repository interfaces.IhelpDeskRepo) interfaceUseCase.IHelpDeskUseCase {
	return &HelpDeskUseCase{repo: repository}
}

func (h *HelpDeskUseCase) CreateRequest(name, phoneNumber, subject, message string) error {
	err := h.repo.CreateRequest(name, phoneNumber, subject, message)
	if err != nil {
		return err
	}
	return nil
}

func (h *HelpDeskUseCase) UpdateAnswer(requestID uint, answer string) error {
	err := h.repo.UpdateAnswer(requestID, answer)
	if err != nil {
		return err
	}
	return nil
}

func (h *HelpDeskUseCase) GetRepliedRequests() ([]responsemodel.HelpDeskResponse, error) {
	replay, err := h.repo.GetRepliedRequests()
	if err != nil {
		return nil, err
	}
	return replay, nil
}

func (h *HelpDeskUseCase) GetUnrepliedRequests() ([]responsemodel.HelpDeskResponse, error) {
	unreplay, err := h.repo.GetUnrepliedRequests()
	if err != nil {
		return nil, err
	}
	return unreplay, nil
}
