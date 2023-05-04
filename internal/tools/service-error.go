package tools

import "fmt"

type ServiceError struct {
	Code uint32
	Err  error
}

func (err *ServiceError) Error() string {
	return fmt.Sprintf("%v", err.Err)
}
