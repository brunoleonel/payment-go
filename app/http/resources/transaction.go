package resources

//Transaction represents a transaction
type Transaction struct {
	AccountID       int64   `json:"account_id"`
	OperationTypeID int64   `json:"operation_type_id"`
	Amount          float32 `json:"amount"`
}
