package respsvc

import "context"

type Service interface {
	Response(ctx context.Context, description string, data interface{}) *ResponseData
}
