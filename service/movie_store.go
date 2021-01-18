package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/eresh/movielibrary/pb"
	"github.com/jinzhu/copier"
)

// AlreadyExistsError is returned when a record with the same ID already exists in the store
var AlreadyExistsError = errors.New("Movie already exists")

//MovieStore MovieStore
type MovieStore interface {
	Save(movie *pb.Movie) error
}

// InMemoryMovieStore stores Movies in memory
type InMemoryMovieStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Movie
}

// NewInMemoryMovieStore returns a new InMemoryMovieStore
func NewInMemoryMovieStore() *InMemoryMovieStore {
	return &InMemoryMovieStore{
		data: make(map[string]*pb.Movie),
	}
}

// Save saves the laptop to the store
func (store *InMemoryMovieStore) Save(laptop *pb.Movie) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return AlreadyExistsError
	}

	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

func deepCopy(laptop *pb.Movie) (*pb.Movie, error) {
	other := &pb.Movie{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy Movie data: %w", err)
	}
	return other, nil
}
