package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/MalteBlackN/29decTrial/proto"

	"google.golang.org/grpc"
)

var hashTable = make(map[int32]int32)

var lock = make(chan bool, 1)

type Server struct {
	pb.UnimplementedHashTableServer
}

func (s *Server) Put(ctx context.Context, in *pb.PutRequest) (*pb.PutResponse, error) {
	// lock
	<-lock
	fmt.Println("Put")

	hashTable[in.Key] = in.Value
	// unlock
	lock <- true
	return &pb.PutResponse{Success: true}, nil
}

func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	// lock
	<-lock
	fmt.Println("Get")
	if hashTable[in.Key] == 0 {
		// unlock
		lock <- true
		return &pb.GetResponse{Value: hashTable[in.Key], Success: false}, nil
	} else {
		// unlock
		lock <- true
		return &pb.GetResponse{Value: hashTable[in.Key], Success: true}, nil
	}
}

func main() {

	//unlock
	lock <- true

	//Setting portnumber
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 9080
	ownPortStr := strconv.Itoa(int(ownPort))
	log.Println("Starting server on port " + ownPortStr)

	//Listening on own port and creating and setting up server
	list, err := net.Listen("tcp", ":"+ownPortStr)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", ownPortStr, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHashTableServer(grpcServer, &Server{})
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
