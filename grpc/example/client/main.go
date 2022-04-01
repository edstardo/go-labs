package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "grpc-example/search"
)

const (
	address = "localhost:54321"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewSearchClient(conn)

	for {
		var name string
		fmt.Print("\nSearch name: ")
		fmt.Scanln(&name)

		if name == "exit" {
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		res, err := c.GetPerson(ctx, &pb.Query{Name: name})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Person -  Name: %s, Age: %d", res.GetName(), res.GetAge())
	}
}
