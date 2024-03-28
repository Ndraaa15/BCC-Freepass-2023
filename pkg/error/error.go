package error

type CustomError struct {
	Code     int
	Location string
	Message  string
	Err      error
}

func (c *CustomError) Error() string {
	return c.Err.Error()
}

func NewCustomError(code int, location string, message string, err error) *CustomError {
	return &CustomError{
		Code:     code,
		Location: location,
		Err:      err,
	}
}
