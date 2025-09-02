package errors

import "errors"

var (
	NotFound        = errors.New("Resource not found")
	InvalidInput    = errors.New("Invalid Input")
	AlreadyFinished = errors.New("Task already finished")
)
