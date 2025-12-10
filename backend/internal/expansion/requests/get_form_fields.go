package requests

import "github.com/gui-henri/learning-go/internal/expansion/domain"

type GetAvaliationFormResponse struct {
	Data domain.AvaliationFormData
	Err  error
}
