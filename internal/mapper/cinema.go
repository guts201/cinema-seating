package mapper

import (
	pb "cinema/api"
	"cinema/internal/entity"
	"cinema/pkg/ent"
)

func ToProtoCinema(c entity.Cinema) *pb.Cinema {
	return &pb.Cinema{
		Id:          int64(c.ID),
		Name:        c.Name,
		MinDistance: uint32(c.MinDistance),
		Rows:        uint32(c.Rows),
		Columns:     uint32(c.Columns),
	}
}

func ToProtoCinemas(cs []entity.Cinema) []*pb.Cinema {
	var cinemas []*pb.Cinema
	for _, c := range cs {
		cinemas = append(cinemas, ToProtoCinema(c))
	}
	return cinemas
}

func FromProtoCinema(pc *pb.Cinema) entity.Cinema {
	return entity.Cinema{
		ID:          int(pc.Id),
		Name:        pc.Name,
		MinDistance: int(pc.MinDistance),
		Rows:        int(pc.Rows),
		Columns:     int(pc.Columns),
	}
}

func ToRepoCinema(c entity.Cinema) *ent.Cinema {
	return &ent.Cinema{
		ID:          int64(c.ID),
		Name:        c.Name,
		MinDistance: uint32(c.MinDistance),
		NumRow:      uint32(c.Rows),
		NumColumn:   uint32(c.Columns),
	}
}

func FromRepoCinema(rc *ent.Cinema) *entity.Cinema {
	return &entity.Cinema{
		ID:          int(rc.ID),
		Name:        rc.Name,
		MinDistance: int(rc.MinDistance),
		Rows:        int(rc.NumRow),
		Columns:     int(rc.NumColumn),
	}
}

func FromRepoCinemas(rcs []*ent.Cinema) []entity.Cinema {
	var cinemas []entity.Cinema
	for _, rc := range rcs {
		cinemas = append(cinemas, *FromRepoCinema(rc))
	}
	return cinemas
}
