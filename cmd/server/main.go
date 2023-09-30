package main

import (
	"context"
	pb "grpc-go/cmd/server/pb/event"
	"io"
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

func (s *server) FmbStream(srv pb.EventService_FmbStreamServer) error {
	log.Println("starting stream...")
	var fmb []*pb.Flight
	ctx := srv.Context()
	for {
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		// // continue if number reveived from stream
		// // less than max
		// if req.Num <= max {
		// 	continue
		// }
		if req.Action == "add" {
			fmb = append(fmb, &pb.Flight{Id: req.Id, Pilot: req.Pilot, Squadron: req.Squadron, FuelStatus: req.FuelStatus, Ac: req.Ac, LastUpdate: time.Now().String()})
		}

		// update max and send it to stream
		resp := pb.FmbStreamResponse{Flights: fmb}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send new flights=%s", fmb)
	}
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
