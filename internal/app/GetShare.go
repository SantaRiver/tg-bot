package app

import (
	"context"
	pb "homework-2/pkg/api"
	"log"
)

func (s *Server) GetShare(ctx context.Context, shareReq *pb.Share) (share *pb.Share, err error) {
	share, err = s.repo.GetShare(ctx, shareReq)
	if err != nil {
		return &pb.Share{}, err
	}
	log.Printf("Get Share: %v", share)
	return
}
