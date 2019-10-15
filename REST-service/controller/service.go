package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Punam-Gaikwad/RestVsGRPC/REST-service/views"
	pb "github.com/Punam-Gaikwad/RestVsGRPC/gRPC-service-Reciever/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9091"
)

func movieRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Did not connect: %v", err)
			}
			defer conn.Close()
			client := pb.NewMovieServiceClient(conn)

			movies, _ := client.GetMovies(context.Background(), &pb.GetRequest{})
			data := views.Response{
				Code: http.StatusOK,
				Body: movies,
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
