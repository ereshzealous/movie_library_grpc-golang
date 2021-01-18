package service

import (
	"context"
	"errors"
	"log"

	"github.com/eresh/movielibrary/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MovieLibraryService is service struct
type MovieLibraryService struct {
	movieStore MovieStore
}

// NewMovieLibraryService is new instance.
func NewMovieLibraryService(movieStore MovieStore) *MovieLibraryService {
	return &MovieLibraryService{movieStore}
}

// SaveMovie is save method.
func (server *MovieLibraryService) SaveMovie(
	ctx context.Context,
	req *pb.AddMovieRequest) (*pb.AddMovieResponse, error) {
	movie := req.GetMovie()
	log.Printf("Received Save Movie Request with id => %s", movie.Id)
	if len(movie.Id) > 0 {
		_, err := uuid.Parse(movie.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Movie ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new Movie ID: %v", err)
		}
		movie.Id = id.String()
	}

	if err := contextError(ctx); err != nil {
		return nil, err
	}
	// save the Movie to in-memory datastore
	err := server.movieStore.Save(movie)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, AlreadyExistsError) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save Movie Details to the store: %v", err)
	}

	log.Printf("saved Movie with id: %s", movie.Id)

	res := &pb.AddMovieResponse{
		Id: movie.Id,
	}
	return res, nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
