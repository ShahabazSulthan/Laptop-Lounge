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

func (s *adminUseCase) GetAllSellersDetailAdminDashboard() (*responsemodel.AdminDashBoard, error) {
	var dashboard responsemodel.AdminDashBoard

	//fmt.Println("nnnn", dashboard)
	var err error
	dashboard.TotalSellers, err = s.repo.GetSellersDetailDashBoard("")
	if err != nil {
		log.Println("Error fetching total sellers:", err)
		return nil, err
	}

	dashboard.ActiveSellers, err = s.repo.GetSellersDetailDashBoard("active")
	if err != nil {
		log.Println("Error fetching active sellers:", err)
		return nil, err
	}

	dashboard.BlockSellers, err = s.repo.GetSellersDetailDashBoard("block")
	if err != nil {
		log.Println("Error fetching blocked sellers:", err)
		return nil, err
	}

	dashboard.PendingSellers, err = s.repo.GetSellersDetailDashBoard("pending")
	if err != nil {
		log.Println("Error fetching pending sellers:", err)
		return nil, err
	}

	// dashboard.TotalOrders, dashboard.TotalRevenue, err = s.repo.TotalRevenue()
	// if err != nil {
	// 	log.Println("Error fetching total revenue:", err)
	// 	return nil, err
	// }

	// dashboard.TotalCredit, err = s.repo.GetNetCredit()
	// if err != nil {
	// 	log.Println("Error fetching total credit:", err)
	// 	return nil, err
	// }

	return &dashboard, nil
}
