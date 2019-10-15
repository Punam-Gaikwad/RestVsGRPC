package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/Punam-Gaikwad/RestVsGRPC/gRPC-service-Reciever/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port       = ":9091"
	moviesFile = "movies.json"
)

type Service struct{}

func (s *Service) GetMovies(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	var movies []*pb.Movie
	data, err := ioutil.ReadFile(moviesFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &movies)

	for _, v := range movies {
		log.Println(v.Title, v.Year, v.ImdbID)
	}

	return &pb.Response{Movies: movies}, nil
}

func main() {

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterMovieServiceServer(s, &Service{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
