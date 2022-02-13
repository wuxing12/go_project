package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"textgrpc/send"
)

func main() {
	rpcServer := grpc.NewServer()
	send.RegisterSendServiceServer(rpcServer, new(send.SendService))

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcServer.Serve(lis)
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
