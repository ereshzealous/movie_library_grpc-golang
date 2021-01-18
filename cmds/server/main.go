package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/eresh/movielibrary/pb"
	"github.com/eresh/movielibrary/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("server started on port %d", *port)
	movieServer := service.NewMovieLibraryService(service.NewInMemoryMovieStore())
	grpcServer := grpc.NewServer()
	pb.RegisterMovieLibraryServiceServer(grpcServer, movieServer)
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listner, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Can not start server")
	}
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal("Can not start server")
	}

}
