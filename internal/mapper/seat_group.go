package mapper

import (
	pb "cinema/api"
	"cinema/internal/entity"
)

func ToProtoSeatGroup(seats []entity.Seat) *pb.SeatGroup {
	protoSeats := make([]*pb.Seat, len(seats))
	for i, s := range seats {
		protoSeats[i] = ToProtoSeat(s)
	}
	return &pb.SeatGroup{
		Seats: protoSeats,
	}
}
