package main

import (
	"context"
	"flag"
	"log"

	"github.com/eresh/movielibrary/generate"
	"github.com/eresh/movielibrary/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	movieClient := pb.NewMovieLibraryServiceClient(conn)
	movie := generate.NewMovie()

	request := &pb.AddMovieRequest{
		Movie: movie,
	}
	log.Printf("Movie Request : %s", request)
	response, err := movieClient.SaveMovie(context.Background(), request)

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Fatal("Already Exists")
		} else {
			log.Fatal("Can not Save Movie")
		}
		return
	}
	log.Printf("Created Movie With id : %s", response.Id)
}
