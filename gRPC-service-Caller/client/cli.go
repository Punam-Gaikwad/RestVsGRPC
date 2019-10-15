package main

import (
	"log"

	"context"

	pb "github.com/Punam-Gaikwad/RestVsGRPC/gRPC-service-Caller/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9092"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewRecordsServiceClient(conn)

	getAll, er := client.GetRecords(context.Background(), &pb.Request{})
	if er != nil {
		log.Fatalf("Could not list records: %v", err)
	}

	for _, v := range getAll.Records {
		log.Println(v.Title, v.Year, v.ImdbID)
	}
}
