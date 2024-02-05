package errorsvc

import "context"

type ErrorRepo interface {
	GetError(ctx context.Context, internalCode string) (*ErrorDetail, error)
}

type Service interface {
	Error(ctx context.Context, internalCode string, description string) *ErrorResponse
}
