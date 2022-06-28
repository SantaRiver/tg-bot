package app

import (
	pb "homework-2/pkg/api"
)

type Server struct {
	repo Repository
	pb.UnimplementedSharesServiceServer
}

func New(repo Repository) *Server {
	return &Server{repo: repo}
}
