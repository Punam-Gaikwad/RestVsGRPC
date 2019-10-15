package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/Punam-Gaikwad/RestVsGRPC/gRPC-service-Caller/proto"
)

const (
	grpcPort = "9092"
)

var (
	getEndpoint = flag.String("get", "localhost:"+grpcPort, "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterRecordsServiceHandlerFromEndpoint(ctx, mux, *getEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8000", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
