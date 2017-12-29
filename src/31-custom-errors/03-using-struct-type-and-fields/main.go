package main

import (
	"errors"
	"fmt"
)

type AppErrorCode int

const (
	BusinessExceptionCode  AppErrorCode = 1000
	ParameterExceptionCode AppErrorCode = 2000
)

type AppError struct {
	err  string
	code AppErrorCode
}

func (e *AppError) Error() string {
	return fmt.Sprintf("error code: %d, err: %s", e.code, e.err)
}

func businessOperate(id int) ([]int, error) {
	if id < 0 {
		return nil, &AppError{"id is negative", ParameterExceptionCode}
	}
	if id > 100 {
		return nil, &AppError{"no valid campaign found", BusinessExceptionCode}
	}
	if id > 200 {
		return nil, errors.New("system error")
	}
	campaignIds := []int{1, 2}
	return campaignIds, nil
}

func main() {
	campaignIds, err := businessOperate(101)
	if err != nil {
		if err, ok := err.(*AppError); ok {
			fmt.Printf("%v", err)
			return
		}
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%#v", campaignIds)
}
