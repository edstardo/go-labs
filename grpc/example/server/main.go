package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net"
	"time"

	pb "grpc-example/search"

	"google.golang.org/grpc"
)

type SearchServer struct {
	pb.UnimplementedSearchServer
}

func (s *SearchServer) GetPerson(ctx context.Context, q *pb.Query) (*pb.Person, error) {
	log.Printf("query received: %v\n", q)

	cleanup := func() {
		// simulation of cleanup function
		log.Println("context done, running cleanup...")
		time.Sleep(time.Second)
		log.Println("cleanup done")
	}

	for {
		select {
		case <-ctx.Done():
			// simulation of actual cleanup function
			// necessary when client context is done
			cleanup()

			return nil, errors.New("ctx done")
		case <-time.After(time.Duration(500 + rand.Int63n(500)*time.Hour.Milliseconds())):
			// simmulation of a time comsuming task
			temp := &pb.Person{
				Name: q.GetName(),
				Age:  20 + rand.Int31n(20),
			}

			return temp, nil
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":54321")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	searchServer := &SearchServer{}

	pb.RegisterSearchServer(grpcServer, searchServer)

	log.Printf("server listening at %v\n", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
