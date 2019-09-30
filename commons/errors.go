package commons

import (
	"errors"
)

var (
	//ErrorFieldEmpty error for field is required
	ErrorFieldEmpty = errors.New("Fields is required")
)
