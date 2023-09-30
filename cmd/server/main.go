package main

import (
	"context"
	pb "grpc-go/cmd/server/pb/event"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedEventServiceServer
}

func (s *server) Retrieve(ctx context.Context, in *pb.RetrieveRequest) (*pb.RetrieveResponse, error) {
	log.Printf("Received Request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.RetrieveResponse{
		Flights: getFlights(),
	}, nil
}

func (s *server) Insert(ctx context.Context, in *pb.InsertRequest) (*pb.InsertResponse, error) {
	log.Printf("Received Insert: %v", in)
	return &pb.InsertResponse{
		Response: map[string]string{"Received": "success"},
	}, nil
}

func getFlights() []*pb.Flight {
	sampleFlights := []*pb.Flight{
		{
			Id:         "1",
			Pilot:      "Viper",
			Ac:         "F18",
			LastUpdate: time.Now().AddDate(0, -4, 0).String(),
			Squadron:   "abc",
			FuelStatus: float32(0.95),
		},
		{
			Id:         "2",
			Pilot:      "Goose",
			Ac:         "F18",
			LastUpdate: time.Now().AddDate(0, -1, 4).String(),
			Squadron:   "abc",
			FuelStatus: float32(0.93),
		},
		{
			Id:         "3",
			Pilot:      "Maverick",
			Ac:         "F18",
			LastUpdate: time.Now().AddDate(0, -4, 2).String(),
			Squadron:   "abc",
			FuelStatus: float32(0.90),
		},
	}
	return sampleFlights
}

func main() {
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalln("error starting listener: ", err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterEventServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
