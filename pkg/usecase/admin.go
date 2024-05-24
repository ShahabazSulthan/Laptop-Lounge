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
	"log"
)

type adminUseCase struct {
	repo             interfaces.IAdminRepository
	tokenSecurityKey config.Token
}

func NewAdminUseCase(adminRepo interfaces.IAdminRepository, key *config.Token) interfaceUseCase.IAdminUseCAse {
	return &adminUseCase{
		repo:             adminRepo,
		tokenSecurityKey: *key,
	}
}

func (s *adminUseCase) AdminLogin(adminData *requestmodel.AdminLoginData) (*responsemodel.AdminLoginres, error) {
	var adminLoginRes responsemodel.AdminLoginres

	hashedPassword, err := s.repo.GetPassword(adminData.Email)
	if err != nil {
		log.Println("Error fetching admin password:", err)
		return nil, errors.New("failed to fetch admin password")
	}

	err = helper.CompairPassword(hashedPassword, adminData.Password)
	if err != nil {
		log.Println("Error comparing passwords:", err)
		return nil, errors.New("password mismatch")
	}

	token, err := service.GenerateRefreshToken(s.tokenSecurityKey.AdminSecurityKey)
	if err != nil {
		log.Println("Error generating refresh token:", err)
		return nil, errors.New("failed to generate refresh token")
	}

	//fmt.Println("hhh", s.tokenSecurityKey.AdminSecurityKey)
	adminLoginRes.Token = token
	return &adminLoginRes, nil
}

func (r *adminUseCase) GetAllSellersDetailAdminDashboard() (*responsemodel.AdminDashBoard, error) {
	var dashBord responsemodel.AdminDashBoard
	var err error

	dashBord.TotalSellers, err = r.repo.GetSellersDetailDashBoard("")
	if err != nil {
		return nil, err
	}

	dashBord.ActiveSellers, err = r.repo.GetSellersDetailDashBoard("active")
	if err != nil {
		return nil, err
	}

	dashBord.BlockSellers, err = r.repo.GetSellersDetailDashBoard("block")
	if err != nil {
		return nil, err
	}

	dashBord.PendingSellers, err = r.repo.GetSellersDetailDashBoard("pending")
	if err != nil {
		return nil, err
	}

	dashBord.TotalOrders, dashBord.TotalRevenue, err = r.repo.TotalRevenue()
	if err != nil {
		return nil, err
	}

	dashBord.TotalCredit, err = r.repo.GetNetCredit()
	if err != nil {
		return nil, err
	}

	dashBord.Coupons, err = r.repo.GetCouponDetails()
	if err != nil {
		return nil, err
	}

	return &dashBord, nil
}
