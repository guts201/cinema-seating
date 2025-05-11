package mapper

import (
	pb "cinema/api"
	"cinema/internal/entity"
)

func ToProtoSeatGroup(group entity.SeatGroup) *pb.SeatGroup {
	protoSeats := make([]*pb.Seat, len(group.Seats))
	for i, s := range group.Seats {
		protoSeats[i] = ToProtoSeat(s)
	}
	return &pb.SeatGroup{
		Seats: protoSeats,
	}
}
