package errorsvc

import "context"

type ErrorRepo interface {
	GetError(ctx context.Context, internalCode string) (*ErrorDetail, error)
	CreateErrorForCode(ctx context.Context, errDetail *ErrorDetail) (*ErrorDetail, error)
}

type ErrorService interface {
	Error(ctx context.Context, internalCode string, description string) *ErrorResponse
	CreateErrorForCode(ctx context.Context, errDetail *ErrorDetail) (*ErrorDetail, error)
}
