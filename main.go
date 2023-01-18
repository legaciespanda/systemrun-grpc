/*
Copyright © 2023 Ernest Obot <ernestobot.dev@gmail.com>

*/

package main

import (
	"log"
	"net"
	"runtime"

	fn "github.com/legaciespanda/systemrun-grpc/function"
	pb "github.com/legaciespanda/systemrun-grpc/pb"
	"google.golang.org/grpc"
)

const (
	port = ":2023"
)

func main() {

	//SystemRun only runs on windows machine
	if runtime.GOOS != "windows" {
		panic("SystemRun Error -):  Only windows Operating System is supported for now")
	}

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to startserver %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCommandServer(grpcServer, &fn.Commandserver{})
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to start gRPC server %v", err)
	}
}
