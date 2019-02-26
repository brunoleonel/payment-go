package controllers

//PaymentController handles requests related to payments
type PaymentController struct {
}

//Post creates one or more payments
func (c *PaymentController) Post() []string {
	arr := []string{"Payment"}
	return arr
}
