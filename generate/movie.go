package generate

import (
	"github.com/eresh/movielibrary/pb"
	"github.com/golang/protobuf/ptypes"
)

// NewMovie generate new Movie
func NewMovie() *pb.Movie {
	movie := &pb.Movie{
		Id:              randomID(),
		Title:           randomString(),
		ProductionHouse: randomString(),
		ReleaseYear:     uint32(randomYear()),
		Budget:          randomBudget(),
		OriginalInfo:    NewOriginalMovie(),
		CastCrew:        NewCastCrew(),
		UpdatedAt:       ptypes.TimestampNow(),
	}

	return movie
}

// NewOriginalMovie generator
func NewOriginalMovie() *pb.OriginalInfo {
	originalInfo := &pb.OriginalInfo{
		OriginalTitle:    randomString(),
		OriginalBudget:   randomBudget(),
		OriginalLanguage: randomString(),
	}
	return originalInfo
}

// NewCastCrew generator
func NewCastCrew() *pb.CastCrew {
	castCrew := &pb.CastCrew{
		Actors:           randomStringArray(),
		Writers:          randomStringArray(),
		Producers:        randomStringArray(),
		ProductionHouse:  randomStringArray(),
		MusicComposition: randomStringArray(),
	}
	return castCrew
}
