package temporal

import (
	"context"
	"errors"

	"github.com/amari/snowflake-monorepo/snowflake-go/internal/snowflake"
	snowflaketemporalv1 "github.com/amari/snowflake-monorepo/snowflake-go/pkg/proto/snowflake/temporal/v1"
	snowflakev1 "github.com/amari/snowflake-monorepo/snowflake-go/pkg/proto/snowflake/v1"
	"go.temporal.io/sdk/temporal"
)

// ActivityObject defines the interface for snowflake-related activities.
type ActivityObject interface {
	// NextSnowflakeActivity generates the next snowflake.
	NextSnowflakeActivity(ctx context.Context, input *snowflaketemporalv1.NextSnowflakeActivityInput) (*snowflaketemporalv1.NextSnowflakeActivityOutput, error)

	// BatchNextSnowflakeActivity generates a batch of snowflakes.
	BatchNextSnowflakeActivity(ctx context.Context, input *snowflaketemporalv1.BatchNextSnowflakeActivityInput) (*snowflaketemporalv1.BatchNextSnowflakeActivityOutput, error)
}

// GRPCActivityObject is an implementation of ActivityObject that uses a SnowflakeServiceClient.
type GRPCActivityObject struct {
	SnowflakeClient snowflakev1.SnowflakeServiceClient
}

func (a *GRPCActivityObject) NextSnowflakeActivity(ctx context.Context, input *snowflaketemporalv1.NextSnowflakeActivityInput) (*snowflaketemporalv1.NextSnowflakeActivityOutput, error) {
	resp, err := a.SnowflakeClient.NextSnowflake(ctx, &snowflakev1.NextSnowflakeRequest{
		Wait: input.Wait,
	})
	if err != nil {
		return nil, snowflakeErrorToTemporalError(err)
	}

	return &snowflaketemporalv1.NextSnowflakeActivityOutput{
		Snowflake: resp.Snowflake,
	}, nil
}

func (a *GRPCActivityObject) BatchNextSnowflakeActivity(ctx context.Context, input *snowflaketemporalv1.BatchNextSnowflakeActivityInput) (*snowflaketemporalv1.BatchNextSnowflakeActivityOutput, error) {
	resp, err := a.SnowflakeClient.BatchNextSnowflake(ctx, &snowflakev1.BatchNextSnowflakeRequest{
		BatchSize: input.Count,
		Wait:      input.Wait,
	})
	if err != nil {
		return nil, snowflakeErrorToTemporalError(err)
	}

	return &snowflaketemporalv1.BatchNextSnowflakeActivityOutput{
		Snowflakes: resp.Snowflakes,
	}, nil
}

// ServiceActivityObject is an implementation of ActivityObject that uses a SnowflakeService.
type ServiceActivityObject struct {
	SnowflakeService *snowflake.SnowflakeService
}

func (a *ServiceActivityObject) NextSnowflakeActivity(ctx context.Context, input *snowflaketemporalv1.NextSnowflakeActivityInput) (*snowflaketemporalv1.NextSnowflakeActivityOutput, error) {
	id, err := a.SnowflakeService.NextID(ctx, input.Wait)
	if err != nil {
		return nil, snowflakeErrorToTemporalError(err)
	}

	return &snowflaketemporalv1.NextSnowflakeActivityOutput{
		Snowflake: &snowflakev1.Snowflake{
			Int64Value:  int64(id),
			StringValue: id.String(),
		},
	}, nil
}

func (a *ServiceActivityObject) BatchNextSnowflakeActivity(ctx context.Context, input *snowflaketemporalv1.BatchNextSnowflakeActivityInput) (*snowflaketemporalv1.BatchNextSnowflakeActivityOutput, error) {
	ids, err := a.SnowflakeService.BatchNextID(ctx, int(input.Count), input.Wait)
	if err != nil {
		return nil, snowflakeErrorToTemporalError(err)
	}

	snowflakes := make([]*snowflakev1.Snowflake, len(ids))
	for i, id := range ids {
		snowflakes[i] = &snowflakev1.Snowflake{
			Int64Value:  int64(id),
			StringValue: id.String(),
		}
	}

	return &snowflaketemporalv1.BatchNextSnowflakeActivityOutput{
		Snowflakes: snowflakes,
	}, nil
}

func snowflakeErrorToTemporalError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, snowflake.ErrSequenceOverflow) {
		return temporal.NewApplicationError("sequence overflow: too many IDs generated in the same millisecond", "snowflake.ErrSequenceOverflow")
	}

	if errors.Is(err, snowflake.ErrClockBackwards) {
		return temporal.NewApplicationError("clock moved backwards, refusing to generate id", "snowflake.ErrClockBackwards")
	}

	return temporal.NewNonRetryableApplicationError("an internal error occurred", "snowflake", err)
}
