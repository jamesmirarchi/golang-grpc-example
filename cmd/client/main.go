package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-go/cmd/client/pb/event"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial("localhost:50001", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewEventServiceClient(cc)
	req := &pb.InsertRequest{
		Id:         "4",
		Pilot:      "Icer",
		Ac:         "F35",
		LastUpdate: time.Now().AddDate(0, -2, 8).String(),
		Squadron:   "abc",
		FuelStatus: float32(0.93),
	}

	resp, err := client.Insert(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive response => %s ", resp.Response["Received"])
}
