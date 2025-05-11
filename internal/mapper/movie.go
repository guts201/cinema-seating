package mapper

import (
	pb "cinema/api"
	"cinema/internal/entity"
	"cinema/pkg/ent"
)

func ToProtoMovie(m entity.Movie) *pb.Movie {
	return &pb.Movie{
		Id:              int32(m.ID),
		Title:           m.Title,
		DurationMinutes: int32(m.DurationMinutes),
	}
}

func FromProtoMovie(pm *pb.Movie) entity.Movie {
	return entity.Movie{
		ID:              int(pm.Id),
		Title:           pm.Title,
		DurationMinutes: int(pm.DurationMinutes),
	}
}

func FromRepoMovie(m *ent.Movie) *entity.Movie {
	return &entity.Movie{
		ID:              int(m.ID),
		Title:           m.Title,
		DurationMinutes: int(m.Duration),
	}
}
