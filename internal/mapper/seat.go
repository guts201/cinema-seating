package mapper

import (
	pb "cinema/api"
	entity "cinema/internal/entity"
)

func ToProtoSeat(seat entity.Seat) *pb.Seat {
	return &pb.Seat{
		Row:    int32(seat.Row),
		Column: int32(seat.Column),
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
