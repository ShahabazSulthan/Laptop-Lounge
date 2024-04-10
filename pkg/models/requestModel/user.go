package requestmodel

type UserDetails struct {
	Id              string `json:"id"`
	Name            string `json:"name"      validate:"required"`
	Email           string `json:"email"   validate:"email"`
	Phone           string `json:"phone"  validate:"len=10"`
	Password        string `json:"password,omitempty"  validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword,omitempty" validate:"eqfield=Password"`
	ReferalCode     string `json:"referalcode,omitempty"`
}

type OtpVerification struct {
	Otp string `json:"otp"  validate:"len=6"`
}

type SendOtp struct {
	Phone string `json:"phone"  validate:"len=10"`
}

type UserLogin struct {
	Phone    string `json:"phone" validate:"len=10"`
	Password string `json:"password" validate:"required,min=4"`
}

type ForgetPassword struct {
	Otp             string `json:"otp" validate:"len=6"`
	Password        string `json:"password" validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword" validate:"eqfield=Password"`
}
