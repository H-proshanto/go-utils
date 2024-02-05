package respsvc

import "context"

type ResponseService interface {
	Response(ctx context.Context, description string, data interface{}) *ResponseData
}
