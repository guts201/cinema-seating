package mapper

import (
	pb "cinema/api"
	entity "cinema/internal/entity"
	"cinema/pkg/ent"
)

func ToProtoSeat(seat entity.Seat) *pb.Seat {
	return &pb.Seat{
		Row:    int32(seat.Row),
		Column: int32(seat.Column),
	}
}
func ToProtoSeatWithId(seat entity.SeatWithId) *pb.Seat {
	return &pb.Seat{
		Row:    int32(seat.Row),
		Column: int32(seat.Column),
		Id:     int64(seat.ID),
	}
}
func FromProtoSeat(seat *pb.Seat) entity.Seat {
	return entity.Seat{
		Row:    int(seat.Row),
		Column: int(seat.Column),
	}
}

func FromProtoSeatGroup(g *pb.SeatGroup) entity.SeatGroup {
	seats := make([]entity.Seat, len(g.Seats))
	for i, s := range g.Seats {
		seats[i] = FromProtoSeat(s)
	}
	return entity.SeatGroup{Seats: seats}
}
func FromRepoSeat(g *ent.SeatReservation) entity.Seat {
	return entity.Seat{
		Row:    int(g.RowNum),
		Column: int(g.ColumnNum),
	}
}

func FromRepoSeatWithId(g *ent.SeatReservation) entity.SeatWithId {
	return entity.SeatWithId{
		Row:    int(g.RowNum),
		Column: int(g.ColumnNum),
		ID:     int(g.ID),
	}
}

func ToProtoSeats(seats []entity.SeatWithId) []*pb.Seat {
	protoSeats := make([]*pb.Seat, len(seats))
	for i, s := range seats {
		protoSeats[i] = ToProtoSeatWithId(s)
	}
	return protoSeats
}
