build:
	protoc -I/usr/local/include -I. \
  	  -I$(GOPATH)/src \
  	  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  	  --go_out=plugins=grpc:. \
  	  --proto_path=$(GOPATH)/src/github.com/Punam-Gaikwad/RestVsGRPC/gRPC-service-Caller \
	  proto/record.proto
	  GOOS=linux GOARCH=amd64 go build

gateway:
	protoc -I/usr/local/include -I. \
  	  -I$(GOPATH)/src \
  	  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  	  --grpc-gateway_out=logtostderr=true:. \
  	  --proto_path=$(GOPATH)/src/github.com/Punam-Gaikwad/RestVsGRPC/gRPC-service-Caller \
	  proto/record.proto