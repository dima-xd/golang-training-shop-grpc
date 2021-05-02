package main

import (
	"context"
	"log"
	"net/http"

	pb "golang-training-shop-grpc/proto/go_proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()
	err = pb.RegisterProductServiceHandler(context.Background(), grpcMux, conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe("localhost:8081", grpcMux))
}
