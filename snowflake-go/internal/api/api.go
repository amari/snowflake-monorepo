package api

import (
	"context"

	snowflakev1 "github.com/amari/snowflake-monorepo/snowflake-go/internal/proto/snowflake/v1"
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/snowflake"
)

type SnowflakeServiceServer struct {
	snowflakev1.UnimplementedSnowflakeServiceServer

	svc snowflake.SnowflakeService
}

func NewSnowflakeServiceServer(snowflakeService *snowflake.SnowflakeService) *SnowflakeServiceServer {
	return &SnowflakeServiceServer{
		svc: *snowflakeService,
	}
}

func (s *SnowflakeServiceServer) NextSnowflake(ctx context.Context, req *snowflakev1.NextSnowflakeRequest) (*snowflakev1.NextSnowflakeResponse, error) {
	id, err := s.svc.NextID(ctx, req.Wait)
	if err != nil {
		return nil, err
	}
	return &snowflakev1.NextSnowflakeResponse{
		Snowflake: &snowflakev1.Snowflake{
			Int64Value:  int64(id),
			StringValue: id.String(),
		},
	}, nil
}

func (s *SnowflakeServiceServer) BatchNextSnowflake(ctx context.Context, req *snowflakev1.BatchNextSnowflakeRequest) (*snowflakev1.BatchNextSnowflakeResponse, error) {
	ids, err := s.svc.BatchNextID(ctx, int(req.BatchSize), req.Wait)
	if err != nil {
		return nil, err
	}

	snowflakes := make([]*snowflakev1.Snowflake, len(ids))

	resp := &snowflakev1.BatchNextSnowflakeResponse{
		Snowflakes: snowflakes,
	}

	for i, id := range ids {
		resp.Snowflakes[i] = &snowflakev1.Snowflake{
			Int64Value:  int64(id),
			StringValue: id.String(),
		}
	}

	return resp, nil
}
