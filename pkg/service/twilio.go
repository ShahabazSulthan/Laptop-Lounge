package service

import (
	"Laptop_Lounge/pkg/config"
	"errors"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

type twilioOtp struct {
	requirements config.OTP
}

var twilioOTP twilioOtp

func OtpServices(details config.OTP) error {

	// Validate config.OTP details
	if details.AccountSid == "" || details.AuthToken == "" || details.ServiceSid == "" {
		return errors.New("invalid Twilio OTP configuration details")
	}

	twilioOTP.requirements = details

	return nil
}

var tw *twilio.RestClient

func TwilioSetup() error {

	if twilioOTP.requirements.AccountSid == "" || twilioOTP.requirements.AuthToken == "" {
		return errors.New("twilio credentials not set")
	}

	// Initialize Twilio REST client
	tw = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioOTP.requirements.AccountSid,
		Password: twilioOTP.requirements.AuthToken,
	})
	return nil
}

func SendOtp(phone string) (string, error) {
	if phone == "" {
		return "", errors.New("phone number cannot be empty")
	}

	// Validate phone number format or any other necessary checks

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo("+91" + phone)
	params.SetChannel("sms")

	res, err := tw.VerifyV2.CreateVerification(twilioOTP.requirements.ServiceSid, params)
	if err != nil {
		// Handle specific Twilio API errors
		return "", fmt.Errorf("failed to send OTP: %v", err)
	}

	return *res.Sid, nil
}

func VerifyOtp(phone string, otp string) error {
	if phone == "" || otp == "" {
		return errors.New("phone number and OTP are required")
	}

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(otp)

	res, err := tw.VerifyV2.CreateVerificationCheck(twilioOTP.requirements.ServiceSid, params)
	if err != nil {
		// Handle specific Twilio API errors
		return fmt.Errorf("failed to verify OTP: %v", err)
	}

	if *res.Status == "approved" {
		return nil
	}

	// Handle OTP verification failure with a specific error message
	return errors.New("OTP verification failed")
}
