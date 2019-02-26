package controllers

//TransactionController Handles requests related to transactions
type TransactionController struct {
}

//Post Creates a transaction
func (c *TransactionController) Post() []string {
	arr := []string{"Transaction"}
	return arr
}
