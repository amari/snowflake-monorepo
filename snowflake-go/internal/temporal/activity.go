package temporal

import (
	snowflaketemporalv1 "github.com/amari/snowflake-monorepo/snowflake-go/internal/proto/snowflake/temporal/v1"
	snowflakev1 "github.com/amari/snowflake-monorepo/snowflake-go/internal/proto/snowflake/v1"
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/snowflake"
)

// ActivityObject defines the interface for snowflake-related activities.
type ActivityObject interface {
	// NextSnowflakeActivity generates the next snowflake.
	NextSnowflakeActivity(input *snowflaketemporalv1.NextSnowflakeActivityInput) (*snowflaketemporalv1.NextSnowflakeActivityOutput, error)

	// BatchNextSnowflakeActivity generates a batch of snowflakes.
	BatchNextSnowflakeActivity(input *snowflaketemporalv1.BatchNextSnowflakeActivityInput) (*snowflaketemporalv1.BatchNextSnowflakeActivityOutput, error)
}

// GRPCActivityObject is an implementation of ActivityObject that uses a SnowflakeServiceClient.
type GRPCActivityObject struct {
	client snowflakev1.SnowflakeServiceClient
}

// ServiceActivityObject is an implementation of ActivityObject that uses a SnowflakeService.
type ServiceActivityObject struct {
	svc *snowflake.SnowflakeService
}
