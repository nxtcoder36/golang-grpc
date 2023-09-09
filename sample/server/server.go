package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sample/proto/sample1"
)

type server struct {
	sample1.GreetingServiceServer
}

func (s *server) Greeting(ctx context.Context, req *sample1.GreetingServiceRequest) (*sample1.GreetingServiceReply, error) {
	message := fmt.Sprintf("Hello, %s!", req.Name)
	reply := &sample1.GreetingServiceReply{
		Message: message,
	}
	return reply, nil
}

func main() {
	listner, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	ser := grpc.NewServer()
	sample1.RegisterGreetingServiceServer(ser, &server{})
	if err := ser.Serve(listner); err != nil {
		log.Fatal("failed to serve", err)
	}
}
