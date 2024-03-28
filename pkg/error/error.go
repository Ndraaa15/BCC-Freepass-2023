package error

type CustomError struct {
	Code     int
	Location string
	Err      error
}

func (c *CustomError) Error() string {
	return c.Err.Error()
}

func NewCustomError(code int, location string, err error) *CustomError {
	return &CustomError{
		Code:     code,
		Location: location,
		Err:      err,
	}
}
