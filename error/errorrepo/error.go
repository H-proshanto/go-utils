package errorrepo

import (
	"context"
	"fmt"

	"github.com/H-proshanto/go-utils/error/errorsvc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ErrorRepo interface {
	errorsvc.ErrorRepo
}

type errorRepo struct {
	svc            *mongo.Database
	collectionName string
}

func NewErrorRepo(collectionName string, svc *mongo.Database) ErrorRepo {
	return &errorRepo{
		collectionName: collectionName,
		svc:            svc,
	}
}

func (r *errorRepo) GetError(ctx context.Context, internalCode string) (*errorsvc.ErrorDetail, error) {
	collection := r.svc.Collection(r.collectionName)
	input := bson.D{{Key: "InternalCode", Value: internalCode}}

	res := collection.FindOne(ctx, input)
	if res.Err() != nil {
		return nil, fmt.Errorf("Failed to get item: %v", res.Err())

	}

	var error *errorsvc.ErrorDetail
	err := res.Decode(&error)
	if err != nil {
		return nil, fmt.Errorf("Error decoding user: %v", err)
	}

	return error, nil
}
