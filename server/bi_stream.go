package main

import (
	"io"
	"log"

	pb "github.com/shoksin/basic-go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionStreaming(stream pb.GreetService_SayHelloBidirectionStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
