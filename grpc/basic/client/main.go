package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:5000"
	defaultName = "Jihar"
)

func main() {
	// Set up a connection to the server.
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := address
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, _ := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	log.Printf("Greeting: %s", r.Message)
}
