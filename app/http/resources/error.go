package resources

//Error represents an error
type Error struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

//Error is a custom error
func (e *Error) Error() string {
	return e.Message
}
