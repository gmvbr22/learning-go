package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/gmvbr/learning-go/grpc/service/gen"
)

type Server struct{
	gen.UnimplementedMainServiceServer
}

func (s *Server) Send(ctx context.Context, in *gen.UserRequest) (*gen.UserReply, error) {
	log.Printf("User Name: %s", in.Name)
	return &gen.UserReply{Message: "Hello world"}, nil
}


func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	gen.RegisterMainServiceServer(grpcServer, &Server{})

	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
