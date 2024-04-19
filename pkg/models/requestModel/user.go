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

type Address struct {
	ID          string `json:"addressID"`
	Userid      string `json:"userid" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName"`
	Street      string `json:"street" validate:"required,alpha"`
	City        string `json:"city" validate:"required,alpha"`
	State       string `json:"state" validate:"required,alpha"`
	Pincode     string `json:"pincode" validate:"min=6"`
	LandMark    string `json:"landmark" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10,number"`
}

type EditAddress struct {
	ID          string `json:"addressID" validate:"required"`
	Userid      string `json:"userid" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Street      string `json:"street" validate:"required,alpha"`
	City        string `json:"city" validate:"required,alpha"`
	State       string `json:"state" validate:"required,alpha"`
	Pincode     string `json:"pincode" validate:"required,min=6"`
	LandMark    string `json:"landmark" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10,number"`
}

type UserEditProfile struct {
	Id              string `json:"userid,omitempty" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required,len=10"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
	ReferalCode     string `json:"-"`
}
