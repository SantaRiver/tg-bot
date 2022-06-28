package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework-2/internal/app"
	"homework-2/internal/db"
	"homework-2/internal/mw"
	"homework-2/internal/repository"
	pb "homework-2/pkg/api"
	"log"
	"net"
	"net/http"
	"time"
)

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterSharesServiceHandlerFromEndpoint(ctx, mux, ":12201", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}

func runGrpc() {
	ctx := context.Background()

	adp, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	rep := repository.New(adp)
	newServer := app.New(rep)
	lis, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("grpc server listening at 8082")

	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(midware.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSharesServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	go runRest()
	runGrpc()
}
