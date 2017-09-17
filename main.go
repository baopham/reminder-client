package main

import (
	pb "github.com/baopham/goproto/reminder"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:8000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewReminderServiceClient(conn)

	createResponse, err := c.Create(context.Background(), &pb.CreateRequest{UserId: "123", Name: "foo"})

	if err != nil {
		log.Fatalf("could not create reminder: %v", err)
	}

	log.Printf(createResponse.Id)

	response, err := c.Get(context.Background(), &pb.GetRequest{Id: createResponse.Id})
	if err != nil {
		log.Fatalf("could not get reminder: %v", err)
	}
	log.Printf("Reminder: %s", response)
}
