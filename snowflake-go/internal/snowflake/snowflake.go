package snowflake

import (
	"fmt"
	"time"
)

type Snowflake int64

const relativeEpoch int64 = 1356069600000 // December 21, 2012 in milliseconds relative to the Unix epoch

// NewSnowflake creates a new Snowflake ID from the given components.
func NewSnowflake(timestamp int64, workerID int64, sequence int64) Snowflake {
	return Snowflake((timestamp << 22) | (workerID << 12) | sequence)
}

// Time returns the time when the Snowflake ID was generated.
func (s Snowflake) Time() time.Time {
	return time.UnixMilli(s.timestamp() + relativeEpoch)
}

// timestamp extracts the timestamp component from the Snowflake ID.
func (s Snowflake) timestamp() int64 {
	return int64((s >> 22) & 0x1FFFFFFFFFF)
}

// workerID extracts the worker ID component from the Snowflake ID.
func (s Snowflake) workerID() int64 {
	return int64((s >> 12) & 0x3FF)
}

// sequence extracts the sequence component from the Snowflake ID.
func (s Snowflake) sequence() int64 {
	return int64(s & 0xFFF)
}

// String returns the string representation of the Snowflake ID.
func (s Snowflake) String() string {
	// Zero-pad to 19 digits b/c log_10(2^63) = 18.96
	return fmt.Sprintf("%019d", s)
}
