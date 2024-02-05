package errSvc

import (
	"context"

	"github.com/H-proshanto/go-utils/helper"
	"github.com/H-proshanto/go-utils/logger"
)

type service struct {
	errRepo ErrorRepo
}

func NewService(

	errorRepo ErrorRepo,

) Service {
	return &service{

		errRepo: errorRepo,
	}
}

func (s *service) Error(ctx context.Context, internalCode string, description string) *ErrorResponse {
	var errDetail *ErrorDetail

	// get from cache
	// errString, err := s.cache.Get(internalCode)
	// if err != nil {
	// 	logger.Error(ctx, "cannot get from redis", err)
	// }
	// if len(errString) > 0 {
	// 	err = json.Unmarshal([]byte(errString), &errDetail)
	// 	if err != nil {
	// 		logger.Error(ctx, "cannot unmarshal error detail", err)
	// 	}
	// }

	// // found in cache
	// if errDetail != nil && len(errDetail.InternalCode) == 0 {
	// 	return &ErrorResponse{
	// 		Timestamp:   helper.GetCurrentTimestamp(),
	// 		Description: description,
	// 		Error:       errDetail,
	// 	}
	// }

	// not found in cache
	// get from db
	errDetail, err := s.errRepo.GetError(ctx, internalCode)
	if err != nil {
		logger.Error(ctx, "cannot get from db", err)
		return &ErrorResponse{
			Timestamp:   helper.GetCurrentTimestamp(),
			Description: description,
			Error: &ErrorDetail{
				InternalCode: internalCode,
				MessageEn:    "Not Set",
				MessageBn:    "Not Set",
			},
		}
	}

	errResponse := &ErrorResponse{
		Timestamp:   helper.GetCurrentTimestamp(),
		Description: description,
		Error:       errDetail,
	}

	return errResponse
}