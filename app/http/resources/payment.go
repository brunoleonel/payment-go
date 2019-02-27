package resources

//Payment represents a payment
type Payment struct {
	AccountID int64   `json:"account_id"`
	Amount    float32 `json:"amount"`
}
