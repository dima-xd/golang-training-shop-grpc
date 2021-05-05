package main

import (
	"context"
	"log"
	"net/http"
	"os"

	pb "github.com/dimaxdqwerty/golang-training-shop-grpc/proto/go_proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	serverAddr = os.Getenv("product-server")
	listen     = os.Getenv("LISTEN")
)

func init() {
	if serverAddr == "" {
		serverAddr = "localhost:8080"
	}
	if listen == "" {
		serverAddr = "localhost:8081"
	}
}

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()
	err = pb.RegisterProductServiceHandler(context.Background(), grpcMux, conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(listen, grpcMux))
}
