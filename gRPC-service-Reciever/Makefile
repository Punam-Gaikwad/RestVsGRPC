build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/Punam-Gaikwad/RestVsGRPC/gRPC-service-Reciever \
	  proto/movie.proto
	GOOS=linux GOARCH=amd64 go build