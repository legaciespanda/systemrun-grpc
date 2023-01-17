/*
Copyright © 2023 Ernest Obot <ernestobot.dev@gmail.com>

*/

package function

import (
	"context"
	"fmt"
	"log"

	hp "github.com/legaciespanda/systemrun-grpc/helper"
	pb "github.com/legaciespanda/systemrun-grpc/pb"
)

// We define a server struct that implements the server interface.
type Commandserver struct {
	//pb.UnimplementedCommandServer
	pb.CommandServer
}

//runs remote executable file on the server
func (s *Commandserver) RunExecutableFile(ctx context.Context, in *pb.RunExecutableRequest) (*pb.RunExecutableReply, error) {
	log.Print("Receive command from client. Launching " + in.FileName + " on the server")
	result, err := hp.StartEXE(in.FileName)

	if err != nil {
		return &pb.RunExecutableReply{Message: "Execution Error"}, fmt.Errorf("%v", err)
	}

	return &pb.RunExecutableReply{Message: result}, nil
}

//kil/terminates remote executable file on the server
func (s *Commandserver) TerminateProcess(ctx context.Context, in *pb.TerminateProcessRequest) (*pb.TerminateProcessReply, error) {
	log.Print("Receive command from client. Terminating " + in.ProcessName + " on the server")
	result, err := hp.Kill(in.ProcessName)

	if err != nil {
		return &pb.TerminateProcessReply{Message: "Termination Error"}, fmt.Errorf("%v", err)
	}

	return &pb.TerminateProcessReply{Message: result}, nil
}

//execute remote command on ther server
func (s *Commandserver) ExecuteRemoteCommand(ctx context.Context, in *pb.RemoteCommandRequest) (*pb.RemoteCommandReply, error) {
	log.Print("Receive command from client. Executing " + in.Command + " on the server")
	result, err := hp.RemoteCommand(in.Command)

	if err != nil {
		return &pb.RemoteCommandReply{Message: "Execution Error"}, fmt.Errorf("%v", err)
	}

	return &pb.RemoteCommandReply{Message: result}, nil
}

//execute remote command on ther server
func (s *Commandserver) LaunchUrl(ctx context.Context, in *pb.LaunchUrlRequest) (*pb.LaunchUrlReply, error) {
	log.Print("Receive command from client. Launching " + in.Url + " on the server default browser")
	result, err := hp.OpenURL(in.Url)

	if err != nil {
		return &pb.LaunchUrlReply{Message: "Launching Error"}, fmt.Errorf("%v", err)
	}

	return &pb.LaunchUrlReply{Message: result}, nil
}

//shutsdown a the remote server machine
func (s *Commandserver) ShutdownSytem(ctx context.Context, in *pb.ShutdownSytemRequest) (*pb.ShutdownSytemReply, error) {
	log.Print("Receive command from client. Executing " + in.Command + " the server")
	result, err := hp.PowerOptions(in.Command)

	if err != nil {
		return &pb.ShutdownSytemReply{Message: "Shutdown Error"}, fmt.Errorf("%v", err)
	}

	return &pb.ShutdownSytemReply{Message: result}, nil
}
