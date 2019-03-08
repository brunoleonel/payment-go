package resources

//Payment represents a payment
type Payment struct {
	AccountID int64   `json:"account_id"`
	Amount    float32 `json:"amount"`
}

//FailedPayment represents a failed payment
type FailedPayment struct {
	Payment Payment `json:"payment"`
	Reason  string  `json:"reason"`
}

//PaymentResponse represents the response of payment operation
type PaymentResponse struct {
	Paid   []Payment       `json:"paid"`
	Failed []FailedPayment `json:"failed"`
}
