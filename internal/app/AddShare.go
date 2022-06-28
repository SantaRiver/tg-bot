package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "homework-2/pkg/api"
)

func (s *Server) AddShare(ctx context.Context, shareReq *pb.Share) (*empty.Empty, error) {
	err := s.repo.AddShare(ctx, shareReq)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
