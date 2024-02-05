package respsvc

import (
	"context"

	"github.com/H-proshanto/go-utils/helper"
)

type responseService struct {
}

func NewService() ResponseService {
	return &responseService{}
}

func (s *responseService) Response(ctx context.Context, description string, data interface{}) *ResponseData {
	return &ResponseData{
		Timestamp:   helper.GetCurrentTimestamp(),
		Description: description,
		Data:        data,
	}
}
