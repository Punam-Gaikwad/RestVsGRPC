package main

import (
	"context"
	"encoding/json"
	"log"
	"net"

	pb "github.com/Punam-Gaikwad/RestVsGRPC/gRPC-service-Caller/proto"
	moviepb "github.com/Punam-Gaikwad/RestVsGRPC/gRPC-service-Reciever/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:9091"
	port    = ":9092"
)

//Service -
type Service struct {
	movieClient moviepb.MovieServiceClient
}

//GetRecords -
func (s *Service) GetRecords(ctx context.Context, req *pb.Request) (*pb.GetResponse, error) {
	movies, err := s.movieClient.GetMovies(context.Background(), &moviepb.GetRequest{})
	if err != nil {
		return nil, err
	}
	bytes, _ := json.Marshal(movies.Movies)

	var records []*pb.Record
	json.Unmarshal(bytes, &records)

	for _, v := range records {
		log.Println(v.Title, v.Year, v.ImdbID)
	}

	return &pb.GetResponse{Records: records}, nil
}

func main() {

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Set up a connection to the movie client.
	conn, e := grpc.Dial(address, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Did not connect: %v", e)
	}
	defer conn.Close()

	movieClient := moviepb.NewMovieServiceClient(conn)

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterRecordsServiceServer(s, &Service{movieClient})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
