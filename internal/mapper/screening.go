package mapper

import (
	pb "cinema/api"
	"cinema/internal/entity"
	"cinema/pkg/ent"
	"time"
)

func ToProtoScreening(s entity.Screening) *pb.Screening {
	return &pb.Screening{
		Id:          int32(s.ID),
		MovieId:     int32(s.MovieID),
		StartTime:   s.StartTime.Format(time.RFC3339),
		MinDistance: int32(s.MinDistance),
		Rows:        int32(s.Row),
		Columns:     int32(s.Column),
	}
}

func ToProtoScreenings(s []entity.Screening) []*pb.Screening {
	var screenings []*pb.Screening
	for _, r := range s {
		screenings = append(screenings, ToProtoScreening(r))
	}
	return screenings
}
func FromProtoScreening(ps *pb.Screening) (entity.Screening, error) {
	startTime, err := time.Parse(time.RFC3339, ps.StartTime)
	if err != nil {
		return entity.Screening{}, err
	}

	return entity.Screening{
		ID:          int(ps.Id),
		MovieID:     int(ps.MovieId),
		StartTime:   startTime,
		MinDistance: int(ps.MinDistance),
	}, nil
}

func FromRepoScreening(rs *ent.Screening) entity.Screening {

	if rs == nil {
		return entity.Screening{}
	}
	if rs.Edges.Cinema == nil || rs.Edges.Movie == nil {
		return entity.Screening{
			ID:          int(rs.ID),
			StartTime:   rs.StartTime,
			MinDistance: 0,
			Row:         0,
			Column:      0,
		}
	}

	return entity.Screening{
		ID:          int(rs.ID),
		MovieID:     int(rs.Edges.Movie.ID),
		StartTime:   rs.StartTime,
		MinDistance: int(rs.Edges.Cinema.MinDistance),
		Row:         int(rs.Edges.Cinema.NumRow),
		Column:      int(rs.Edges.Cinema.NumColumn),
	}
}

func FromRepoScreenings(rs []*ent.Screening) []entity.Screening {
	var screenings []entity.Screening
	for _, r := range rs {
		screenings = append(screenings, FromRepoScreening(r))
	}
	return screenings
}
