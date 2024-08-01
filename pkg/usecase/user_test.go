package usecase

import (
	"Laptop_Lounge/pkg/config"
	"Laptop_Lounge/pkg/mock/mockRepository"
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAddress(t *testing.T) {
	ctrl := gomock.NewController(t)

	userRepo := mockRepository.NewMockIUserRepo(ctrl)

	paymentRepo := mockRepository.NewMockIPaymentRepository(ctrl)

	userUseCase := NewUserUseCase(userRepo, paymentRepo, &config.Token{})

	testData := map[string]struct {
		userID    string
		addressID string
		stub      func(mockRepository.MockIPaymentRepository, mockRepository.MockIUserRepo, string, string)
		wantErr   error
	}{
		"success": {
			userID:    "1",
			addressID: "5",
			stub: func(mir1 mockRepository.MockIPaymentRepository, mir2 mockRepository.MockIUserRepo, s1, s2 string) {
				userRepo.EXPECT().DeleteAddress(s1, s2).Times(1).Return(nil)
			},
			wantErr: nil,
		},
		"failed": {
			userID:    "1",
			addressID: "5",
			stub: func(mir1 mockRepository.MockIPaymentRepository, mir2 mockRepository.MockIUserRepo, s1, s2 string) {
				userRepo.EXPECT().DeleteAddress(s1, s2).Times(1).Return(errors.New("no address exist"))
			},
			wantErr: errors.New("no address exist"),
		},
	}

	for _, tt := range testData {
		tt.stub(*paymentRepo, *userRepo, tt.userID, tt.addressID)
		err := userUseCase.DeleteAddress(tt.userID, tt.addressID)
		assert.Equal(t, err, tt.wantErr)
	}

}

func TestAddAddress(t *testing.T) {
	ctrl := gomock.NewController(t)

	userRepo := mockRepository.NewMockIUserRepo(ctrl)
	paymentRepo := mockRepository.NewMockIPaymentRepository(ctrl)
	userUseCase := NewUserUseCase(userRepo, paymentRepo, &config.Token{})

	testData := map[string]struct {
		args    *requestmodel.Address
		stub    func(mockRepository.MockIPaymentRepository, mockRepository.MockIUserRepo, *requestmodel.Address)
		want    *requestmodel.Address
		wantErr error
	}{
		"success": {
			args: &requestmodel.Address{
				Userid:      "21",
				FirstName:   "Shahabaz",
				LastName:    "Sulthan",
				Street:      "Pallikera",
				City:        "Kasaragod",
				State:       "Kerala",
				Pincode:     "671316",
				LandMark:    "Near Talkies",
				PhoneNumber: "9496703880",
			},
			stub: func(mir1 mockRepository.MockIPaymentRepository, mir2 mockRepository.MockIUserRepo, a *requestmodel.Address) {
				userRepo.EXPECT().CreateAddress(a).Times(1).Return(&requestmodel.Address{
					Userid:      "21",
					FirstName:   "Shahabaz",
					LastName:    "Sulthan",
					Street:      "Pallikera",
					City:        "Kasaragod",
					State:       "Kerala",
					Pincode:     "671316",
					LandMark:    "Near Talkies",
					PhoneNumber: "9496703880"}, nil)
			},
			want: &requestmodel.Address{
				Userid:      "21",
				FirstName:   "Shahabaz",
				LastName:    "Sulthan",
				Street:      "Pallikera",
				City:        "Kasaragod",
				State:       "Kerala",
				Pincode:     "671316",
				LandMark:    "Near Talkies",
				PhoneNumber: "9496703880",
			},
			wantErr: nil,
		},
		"fail": {
			args: &requestmodel.Address{
				Userid:      "21",
				FirstName:   "Shahabaz",
				LastName:    "Sulthan",
				Street:      "Pallikera",
				City:        "Kasaragod",
				State:       "Kerala",
				Pincode:     "671316",
				LandMark:    "Near Talkies",
				PhoneNumber: "9496703880",
			},
			stub: func(mir1 mockRepository.MockIPaymentRepository, mir2 mockRepository.MockIUserRepo, a *requestmodel.Address) {
				userRepo.EXPECT().CreateAddress(a).Times(1).Return(nil, errors.New("missmatch in address data"))
			},
			want:    nil,
			wantErr: errors.New("missmatch in address data"),
		},
	}

	for _, tt := range testData {
		tt.stub(*paymentRepo, *userRepo, tt.args)
		result, err := userUseCase.AddAddress(tt.args)
		assert.Equal(t, tt.want, result)
		assert.Equal(t, tt.wantErr, err)
	}
}
