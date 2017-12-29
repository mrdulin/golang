package models

import "fmt"

type AppErrorCode int

const (
	BusinessErrorCode  AppErrorCode = 1000
	ParameterErrorCode AppErrorCode = 2000
)

type AppError struct {
	Err  string
	Code AppErrorCode
}

func (e *AppError) Error() string {
	return fmt.Sprintf("error code: %d, error: %s", e.Code, e.Err)
}
