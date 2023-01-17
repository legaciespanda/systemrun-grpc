/*
Copyright © 2023 Ernest Obot <ernestobot.dev@gmail.com>

*/

package main

import (
	"context"
	"log"

	pb "github.com/legaciespanda/systemrun-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":2023"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to server %v", err)
	}

	defer conn.Close()

	client := pb.NewCommandClient(conn)

	//terminate or kill a process on the server machine
	terminateProcess, err := client.TerminateProcess(context.Background(), &pb.TerminateProcessRequest{ProcessName: "notepad.exe"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", terminateProcess)

	//runs executable file on the server machine
	runExecutable, err := client.RunExecutableFile(context.Background(), &pb.RunExecutableRequest{FileName: "notepad"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", runExecutable)

	//open a url on default browser of the server machine
	launchUrl, err := client.LaunchUrl(context.Background(), &pb.LaunchUrlRequest{Url: "https://megtrix.com"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", launchUrl)

	////execute remote command on ther server
	executeRemoteCommand, err := client.ExecuteRemoteCommand(context.Background(), &pb.RemoteCommandRequest{Command: "cmd.exe"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", executeRemoteCommand)

	//shutsdown a the remote server machine
	executeShutdown, err := client.ShutdownSytem(context.Background(), &pb.ShutdownSytemRequest{Command: "restart"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", executeShutdown)

}
