package v3

type CreateOrder struct {
	OutOrderNo          string         `json:"out_order_no"`
	AppID               string         `json:"appid"`
	ServiceID           string         `json:"service_id"`
	ServiceIntroduction string         `json:"service_introduction"`
	PostPayments        []PostPayment  `json:"post_payments,omitempty"`
	PostDiscounts       []PostDiscount `json:"post_discounts,omitempty"`
	TimeRange           *TimeRange     `json:"time_range"`
	Location            *Location      `json:"location"`
	RiskFund            *RiskFund      `json:"risk_fund"`
	Attach              string         `json:"attach,omitempty"`
	NotifyUrl           string         `json:"notify_url"`
	OpenID              string         `json:"openid"`
	NeedUserConfirm     bool           `json:"need_user_confirm"`
}

type PostPayment struct {
	Name        string `json:"name,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Count       uint32 `json:"count,omitempty"`
}

type PostDiscount struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Count       uint32 `json:"count,omitempty"`
}

type TimeRange struct {
	StartTime       string `json:"start_time"`
	StartTimeRemark string `json:"start_time_remark,omitempty"`
	EndTime         string `json:"end_time,omitempty"`
	EndTimeRemark   string `json:"end_time_remark,omitempty"`
}

type Location struct {
	StartLocation string `json:"start_location,omitempty"`
	EndLocation   string `json:"end_location,omitempty"`
}

type RiskFund struct {
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Description string `json:"description,omitempty"`
}

type CancelOrder struct {
	OutOrderNo string `json:"-"`
	AppID      string `json:"appid"`
	ServiceID  string `json:"service_id"`
	Reason     string `json:"reason"`
}

type QueryOrder struct {
	OutOrderNo string `json:"-"`
	QueryID    string
	ServiceID  string
	AppID      string
}
