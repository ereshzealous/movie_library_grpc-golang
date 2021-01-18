package service_test

import (
	"context"
	"log"
	"testing"

	"github.com/eresh/movielibrary/generate"
	"github.com/eresh/movielibrary/pb"
	"github.com/eresh/movielibrary/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	movie := generate.NewMovie()
	movie.Id = ""

	movieInvalidID := generate.NewMovie()
	movieInvalidID.Id = "invalid-uuid"

	movieDuplicateID := generate.NewMovie()
	storeDuplicateID := service.NewInMemoryMovieStore()
	err := storeDuplicateID.Save(movieDuplicateID)
	require.Nil(t, err)
	testCases := []struct {
		name  string
		movie *pb.Movie
		store service.MovieStore
		code  codes.Code
	}{
		{
			name:  "success_with_id",
			movie: generate.NewMovie(),
			store: service.NewInMemoryMovieStore(),
			code:  codes.OK,
		},
		{
			name:  "success_no_id",
			movie: generate.NewMovie(),
			store: service.NewInMemoryMovieStore(),
			code:  codes.OK,
		},
		{
			name:  "failure_invalid_id",
			movie: movieInvalidID,
			store: service.NewInMemoryMovieStore(),
			code:  codes.InvalidArgument,
		},
		{
			name:  "failure_duplicate_id",
			movie: movieDuplicateID,
			store: storeDuplicateID,
			code:  codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.AddMovieRequest{
				Movie: tc.movie,
			}

			server := service.NewMovieLibraryService(tc.store)
			res, err := server.SaveMovie(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.movie.Id) > 0 {
					require.Equal(t, tc.movie.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
	log.Println("Completed tests....")
}
