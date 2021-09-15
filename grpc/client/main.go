package main

import (
	"context"
	"log"
	"time"

	"github.com/gmvbr/learning-go/grpc/service/gen"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := gen.NewMainServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	r, err := c.Send(ctx, &gen.UserRequest{Name: "gmvbr"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Recive: %s", r.GetMessage())
}
