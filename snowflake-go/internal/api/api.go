package api

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/metric"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	snowflakev1 "github.com/amari/snowflake-monorepo/snowflake-go/internal/proto/snowflake/v1"
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/snowflake"
	"github.com/rs/zerolog"
)

type SnowflakeServiceServer struct {
	snowflakev1.UnimplementedSnowflakeServiceServer

	svc snowflake.SnowflakeService

	metrics *SnowflakeServiceServerMetrics
}

func NewSnowflakeServiceServer(snowflakeService *snowflake.SnowflakeService, metrics *SnowflakeServiceServerMetrics) *SnowflakeServiceServer {
	return &SnowflakeServiceServer{
		svc:     *snowflakeService,
		metrics: metrics,
	}
}

func (s *SnowflakeServiceServer) NextSnowflake(ctx context.Context, req *snowflakev1.NextSnowflakeRequest) (*snowflakev1.NextSnowflakeResponse, error) {
	id, err := s.svc.NextID(ctx, req.Wait)
	if err != nil {
		if errors.Is(err, snowflake.ErrSequenceOverflow) {
			if s.metrics != nil {
				s.metrics.SequenceOverflowErrorsCounter.Add(ctx, 1)
			}

			// Return a more specific error message for sequence overflow
			// ResourceExhausted is typically used when a resource limit is hit and may require remedial action.
			// However, for transient conditions like sequence overflow (which may resolve on retry), Unavailable is more appropriate.
			return nil, status.Errorf(codes.Unavailable, "sequence overflow: too many IDs generated in the same millisecond")
		}
		if errors.Is(err, snowflake.ErrClockBackwards) {
			if s.metrics != nil {
				s.metrics.ClockBackwardsErrorsCounter.Add(ctx, 1)
			}

			zerolog.Ctx(ctx).Err(err).Msg("clock moved backwards")

			// Return a more specific error message for clock going backwards
			return nil, status.Errorf(codes.FailedPrecondition, "clock moved backwards, refusing to generate id")
		}

		return nil, err
	}

	if s.metrics != nil {
		s.metrics.SnowflakeCounter.Add(ctx, 1)
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
		if errors.Is(err, snowflake.ErrSequenceOverflow) {
			if s.metrics != nil {
				s.metrics.SequenceOverflowErrorsCounter.Add(ctx, 1)
			}

			// Return a more specific error message for sequence overflow
			// ResourceExhausted is typically used when a resource limit is hit and may require remedial action.
			// However, for transient conditions like sequence overflow (which may resolve on retry), Unavailable is more appropriate.
			return nil, status.Errorf(codes.Unavailable, "sequence overflow: too many IDs generated in the same millisecond")
		}
		if errors.Is(err, snowflake.ErrClockBackwards) {
			if s.metrics != nil {
				s.metrics.ClockBackwardsErrorsCounter.Add(ctx, 1)
			}

			zerolog.Ctx(ctx).Err(err).Msg("clock moved backwards")

			// Return a more specific error message for clock going backwards
			return nil, status.Errorf(codes.FailedPrecondition, "clock moved backwards, refusing to generate id")
		}

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

	if s.metrics != nil {
		s.metrics.SnowflakeCounter.Add(ctx, int64(len(ids)))
	}

	return resp, nil
}

type SnowflakeServiceServerMetrics struct {
	SnowflakeCounter              metric.Int64Counter
	ClockBackwardsErrorsCounter   metric.Int64Counter
	SequenceOverflowErrorsCounter metric.Int64Counter
}
