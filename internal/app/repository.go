package app

import (
	"context"
	"errors"
	pb "homework-2/pkg/api"
)

var AlreadyExistErr = errors.New("already exist")
var NotFoundErr = errors.New("not found")

type Repository interface {
	AddShare(ctx context.Context, shareReq *pb.Share) error
	GetShare(ctx context.Context, shareReq *pb.Share) (shareRes *pb.Share, err error)
}
