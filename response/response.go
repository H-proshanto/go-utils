package respSvc

import (
	"context"

	"github.com/H-proshanto/go-utils/helper"
)

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Response(ctx context.Context, description string, data interface{}) *ResponseData {
	return &ResponseData{
		Timestamp:   helper.GetCurrentTimestamp(),
		Description: description,
		Data:        data,
	}
}
