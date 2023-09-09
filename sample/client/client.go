package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sample/proto/sample1"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8000", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := sample1.NewGreetingServiceClient(cc)
	request := &sample1.GreetingServiceRequest{Name: "Mohit"}

	resp, err := client.Greeting(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("receive response => %s", resp.Message)
}
