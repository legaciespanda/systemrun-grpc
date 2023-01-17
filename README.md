## SystemRun gRPC

### Introduction

SystemRun gRPC is a GO unary gRPC that allows a client exexcute commands, administrative tasks and control the remote gRPC sever machine

#### - OBJECTIVE
Assuming the client runs on another machine and the server runs on a different machine.

The client can remotely perform some administrative task on the server machine provided, the client
is connected to the gRPC server.

#### - FEATURES
 - The client can launch/run executable file and also system programs on the server machine
 - Client can shutdown, restart, hybernate or logoff the remote server machine
 - Client can terminate or kill a process on the server machine
 - Client can execute remote commands on the server machine
 - Client can open a url on default browser of the server machine

#### - OS SUPPORTED
 - Windows 8,10,11

#### - HOW TO USE
 - Make sure server is running
 - Run client in the ```client``` folder


#### - CLIENT
```
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


    //execute remote command on ther server
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
```
### ShutdownSytem Supported Commands
- shutdown
- restart
- hibernate
- loggof


