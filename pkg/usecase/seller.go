package usecase

import (
	"Laptop_Lounge/pkg/config"
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"Laptop_Lounge/pkg/service"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"errors"
)

type sellerUseCase struct {
	repo  interfaces.ISellerRepo
	token config.Token
}

func NewSellerUseCase(sellerRepo interfaces.ISellerRepo, token *config.Token) interfaceUseCase.ISellerUseCase {
	return &sellerUseCase{repo: sellerRepo,
		token: *token}
}

func (r *sellerUseCase) SellerSignup(sellerSignupData *requestmodel.SellerSignup) (*responsemodel.SellerSignupRes, error) {
	var SellerSignupRes responsemodel.SellerSignupRes

	count, err := r.repo.IsSellerExist(sellerSignupData.Email)
	if err != nil {
		return &SellerSignupRes, err
	} else {
		if count >= 1 {
			return &SellerSignupRes, errors.New("seller exist with same email id, ")
		}
	}

	hashPassword := helper.HashPassword(sellerSignupData.Password)
	sellerSignupData.Password = hashPassword

	err = r.repo.CreateSeller(sellerSignupData)
	if err != nil {
		return &SellerSignupRes, err
	}

	SellerSignupRes.Result = "Registeration saved ! Your request is now in processing. You will receive a confirmation once you have been admitted and granted access to start selling."
	return &SellerSignupRes, nil
}

func (r *sellerUseCase) SellerLogin(loginData *requestmodel.SellerLogin) (*responsemodel.SellerLoginRes, error) {
	var loginResponse responsemodel.SellerLoginRes

	hashedPassword, sellerID, status, err := r.repo.GetHashPassAndStatus(loginData.Email)
	if err != nil {
		return &loginResponse, err
	}

	if status == "block" {
		return &loginResponse, errors.New("vender blocked by admin")
	}

	if status == "pending" {
		return &loginResponse, errors.New("your request under process pls whait ")
	}

	err = helper.CompairPassword(hashedPassword, loginData.Password)
	if err != nil {
		return &loginResponse, err
	}

	accessToken, err := service.GenerateAcessToken(r.token.SellerSecurityKey, sellerID)
	if err != nil {
		return &loginResponse, err
	}

	refreshToken, err := service.GenerateRefreshToken(r.token.SellerSecurityKey)
	if err != nil {
		return &loginResponse, err
	}

	loginResponse.AccessToken = accessToken
	loginResponse.RefreshToken = refreshToken

	return &loginResponse, nil
}

func (r *sellerUseCase) GetAllSellers(page string, limit string) (*[]responsemodel.SellerDetails, *int, error) {
	ch := make(chan int)

	go r.repo.SellerCount(ch)
	count := <-ch

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, &count, err
	}

	SellerDetails, err := r.repo.AllSellers(offSet, limits)
	if err != nil {
		return nil, nil, err
	}

	return SellerDetails, &count, nil
}

func (r *sellerUseCase) BlockSeller(id string) error {
	err := r.repo.BlockSeller(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *sellerUseCase) ActiveSeller(id string) error {
	err := r.repo.UnblockSeller(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *sellerUseCase) GetAllPendingSellers(page string, limit string) (*[]responsemodel.SellerDetails, error) {

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	SellerDetails, err := r.repo.GetPendingSellers(offSet, limits)
	if err != nil {
		return nil, err
	}

	return SellerDetails, nil
}

func (r *sellerUseCase) FetchSingleSeller(id string) (*responsemodel.SellerDetails, error) {
	sellerData, err := r.repo.GetSingleSeller(id)
	if err != nil {
		return nil, err
	}
	return sellerData, nil
}
