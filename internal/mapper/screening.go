package mapper

import (
    "cinema/internal/entity"
    pb "cinema/api"
    "time"
)

func ToProtoScreening(s entity.Screening) *pb.Screening {
    return &pb.Screening{
        Id: int32(s.ID),
        MovieId: int32(s.MovieID),
        StartTime: s.StartTime.Format(time.RFC3339),
        Rows: int32(s.Rows),
        Columns: int32(s.Columns),
        MinDistance: int32(s.MinDistance),
    }
}

func FromProtoScreening(ps *pb.Screening) (entity.Screening, error) {
    startTime, err := time.Parse(time.RFC3339, ps.StartTime)
    if err != nil {
        return entity.Screening{}, err
    }

    return entity.Screening{
        ID: int(ps.Id),
        MovieID: int(ps.MovieId),
        StartTime: startTime,
        Rows: int(ps.Rows),
        Columns: int(ps.Columns),
        MinDistance: int(ps.MinDistance),
    }, nil
}
