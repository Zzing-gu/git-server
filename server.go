package main

import (
	
	pb "proto"
	"context"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"git"
	
)

type server struct{}


func main(){
	
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGitServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}
	
}

func (s *server) CreateAndInitDirectory(ctx  context.Context, request *pb.Request_Path) (*pb.Response_Result, error) {
	git.CreateAndInitDirectory(request.Path)
	return &pb.Response_Result{Result:"createandinit"}, nil
}

func (s *server) AddOrUpdateFile(ctx  context.Context, request *pb.Request_File) (*pb.Response_Result, error) {
	git.AddOrUpdateFile(request.Path, request.Filedata, request.Filename, request.Filemode)
	return &pb.Response_Result{Result:"addorupdate"}, nil
}

// func (s *server) GitInit(ctx  context.Context, request *pb.Request) (*pb.Response, error) {
// 	git.HosukTest()
// 	git.InitRepository("./",false)
// 	return &pb.Response{Output:1}, nil
// }

// func (s *server) GitInit(ctx  context.Context, request *pb.Request) (*pb.Response, error) {
// 	git.HosukTest()
// 	git.InitRepository("./",false)
// 	return &pb.Response{Output:1}, nil
// }