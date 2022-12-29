package main

import (
	"context"

	"log"

	"os"
	"strconv"
	"strings"

	"bufio"

	pb "github.com/MalteBlackN/29decTrial/proto"

	"google.golang.org/grpc"
)

var totalPorts int64
var reader = bufio.NewReader(os.Stdin)

var clients []pb.HashTableClient

func main() {
	//Loading id and total amount of ports to connect to
	totalPorts, _ = strconv.ParseInt(os.Args[1], 10, 32)

	//Creating connection to all servers
	for i := 0; i < int(totalPorts); i++ {
		// Create a virtual RPC Client Connection on port 9080 + i
		var conn *grpc.ClientConn
		var port int = 9080 + i
		portStr := strconv.Itoa(port)

		conn, err := grpc.Dial(":"+portStr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		// Defer means: When this function returns, call this method (meaning, one main is done, close connection)
		defer conn.Close()

		//  Create new Client from generated gRPC code from proto
		c := pb.NewHashTableClient(conn)
		clients = append(clients, c)
	}

	log.Print("wlcome to the hash table service")

	//Starting method for continuously recieving input from user
	for {
		takeInput()
	}
}

func takeInput() {
	for {
		log.Println("write put to put a value in the hash table, write get to get a value from the hash table.")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		log.Print("(You wrote: " + input + ")")

		if input == "put" {
			log.Println("write the key you want to put in the hash table, followed by enter:")

			tempKey, _ := reader.ReadString('\n')
			tempKey = strings.TrimSpace(tempKey)
			key, err := strconv.Atoi(tempKey)
			if err != nil {
				log.Println("Faulty input, please try again")
				continue
			}

			log.Println("write the value you want to put in the hash table, followed by enter:")
			tempValue, _ := reader.ReadString('\n')
			tempValue = strings.TrimSpace(tempValue)
			value, err := strconv.Atoi(tempValue)
			if err != nil {
				log.Println("Faulty input, please try again")
				continue
			}

			//Calling method to put value in hash table
			result, _ := putValueOnAll(&pb.PutRequest{Key: int32(key), Value: int32(value)})

			//Displaying information to user based on whether the value was put in the hash table or not
			if !result.Success {
				log.Printf("the value was not put onto key)\n")
				continue
			} else {
				log.Printf("the value was put onto key)\n")
				break
			}
		}
		if input == "get" {

			log.Println("write the key you want to get from the hash table, followed by enter:")

			tempKey, _ := reader.ReadString('\n')
			tempKey = strings.TrimSpace(tempKey)
			key, err := strconv.Atoi(tempKey)
			if err != nil {
				log.Println("Faulty input, please try again")
				continue
			}

			//Calling method to get value from hash table
			result, _ := getValueFromAll(&pb.GetRequest{Key: int32(key)})

			if !result.Success {
				log.Printf("the value could not be retrieved)\n")
				continue
			} else {
				log.Printf(("the value is: %d)\n"), result.Value)
				break
			}

		}
	}
}

func putValueOnAll(req *pb.PutRequest) (*pb.PutResponse, error) {
	var ack *pb.PutResponse
	for _, c := range clients {
		tempAck, err := c.Put(context.Background(), req)
		if err == nil {
			ack = tempAck
		}
	}
	return ack, nil
}

func getValueFromAll(req *pb.GetRequest) (*pb.GetResponse, error) {
	var result *pb.GetResponse
	for _, c := range clients {
		tempResult, err := c.Get(context.Background(), req)
		if err == nil {
			result = tempResult
		}
	}
	return result, nil
}
