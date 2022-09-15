package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/jashfeer/gRPC_test/chat"
)

func main() {
	lis, err := net.Listen("tcp",":9090")
	if err !=nil{
		log.Fatalf("Failed to listen on port 9090 : %v",err)

	}
	s:=chat.Server{}


	grpcServer :=grpc.NewServer()
	chat.RegisterChartServiceServer(grpcServer, &s)
	
	if err:= grpcServer.Serve(lis);err!=nil{
		log.Fatalf("Failed to serve grpc server over port 9000:%v",err)
	}
	

}