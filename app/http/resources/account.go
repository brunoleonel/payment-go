package resources

//AvailableCreditLimit represents the credit limit
type AvailableCreditLimit struct {
	Amount float32 `json:"amount"`
}

//AvailableWithdrawalLimit represents the withdrawal limit
type AvailableWithdrawalLimit struct {
	Amount float32 `json:"amount"`
}

//Account represents a new account to be created
type Account struct {
	ID                       int64                    `json:"id"`
	AvailableCreditLimit     AvailableCreditLimit     `json:"available_credit_limit"`
	AvailableWithdrawalLimit AvailableWithdrawalLimit `json:"available_withdrawal_limit"`
}
