package main

import (
	"context"
	"log"

	"github.com/jashfeer/gRPC_test/chat"
	"google.golang.org/grpc"
)

func main(){
	var conn *grpc.ClientConn
	conn,err :=grpc.Dial(":9090",grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("could not connected: %s",err)
	}
	defer conn.Close()

	 c:=chat.NewChartServiceClient(conn)

	 message:=chat.Message{
		Body: "HEllo from the clinte!!",
	 }

	 response,err:=c.SayHello(context.Background(),&message)
	 if err !=nil{
		log.Fatalf("Hello when calling SayHello: %s",err)
	 }
	  
	 log.Printf("Response from Server: %s",response.Body)

}


