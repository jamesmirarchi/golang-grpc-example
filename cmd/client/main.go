package main

import (
	"context"
	"fmt"
	"io"
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

	stream, err := client.FmbStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var fmb []*pb.Flight
	ctx := stream.Context()
	done := make(chan bool)

	// first goroutine sends random increasing numbers to stream
	// and closes it after 10 iterations
	go func() {
		for i := 1; i <= 10; i++ {
			id := fmt.Sprintf("%d", i)
			pilot := fmt.Sprintf("pilot-%d", i)
			req := pb.FmbStreamRequest{
				Id:         id,
				Pilot:      pilot,
				Ac:         "F35",
				LastUpdate: time.Now().AddDate(0, -1, i).String(),
				Squadron:   "abc",
				FuelStatus: float32(0.98),
				Action:     "add",
			}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("%v sent", &req)
			time.Sleep(time.Millisecond * 500)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()
	// second goroutine receives data from stream
	// and saves result in flights variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			fmb = resp.Flights
			log.Printf("new flight board %s received", fmb)
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	log.Printf("finished stream with fmb=%s", fmb)
	<-done
}
