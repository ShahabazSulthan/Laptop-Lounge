package responsemodel

type AdminLoginres struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Result   string `json:"result,omitempty"`
	Token    string `json:"token,omitempty"`
}

type AdminDashBoard struct {
	TotalSellers   uint     `json:"totalsellers"`
	BlockSellers   uint     `json:"blocksellers"`
	PendingSellers uint     `json:"pendingsellers"`
	ActiveSellers  uint     `json:"activesellers"`
	TotalRevenue   uint     `json:"totalrevenue"`
	TotalOrders    uint     `json:"totalorders"`
	TotalCredit    uint     `json:"totalcredit"`
	Coupons        []Coupon `json:"coupons"`
}
