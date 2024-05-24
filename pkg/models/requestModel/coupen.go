package requestmodel

type Coupon struct {
	Name            string `json:"name" validate:"required"`
	Type            string `json:"type" validate:"required,alpha"`
	Discount        uint   `json:"discount" validate:"min=1,max=100"`
	MinimumRequired uint   `json:"minimum_required" validate:"min=0"`
	MaximumAllowed  uint   `json:"maximum_allowed" validate:"gtcsfield=MinimumRequired"`
	ExpireDate      string `json:"expire_date" validate:"required,date=2006-01-02"`
}
