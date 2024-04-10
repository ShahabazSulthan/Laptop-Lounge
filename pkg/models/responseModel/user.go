package responsemodel

type SignupData struct {
	ID            string `json:"userID,omitempty"`
	Name          string `json:"name,omitempty"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phone,omitempty"`
	OTP           string `json:"otp,omitempty"`
	Token         string `json:"token,omitempty"`
	IsUserExist   string `json:"isUserExist,omitempty"`
	ReferalCode   string `json:"referalCode,omitempty"`
	//WalletBalance uint   `json:"walletBalance,omitempty"`
}

type OtpValidation struct {
	Phone        string `json:"phone,omitempty"`
	Otp          string `json:"otp,omitempty"`
	Result       string `json:"result,omitempty"`
	Token        string `json:"token,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type UserLogin struct {
	Phone        string `json:"phone,omitempty"`
	Password     string `json:"password,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	Error        string `json:"error,omitempty"`
}

type TokenVerificationMiddleware struct {
	Error string `json:"error,omitempty"`
}

type UserDetails struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Status string `json:"status,omitempty"`
}

type Errors struct {
	Err string `json:"err,omitempty"`
}
