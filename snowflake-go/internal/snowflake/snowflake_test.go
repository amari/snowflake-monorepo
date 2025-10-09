package snowflake

import (
	"testing"
	"time"
)

func TestNewSnowflake(t *testing.T) {
	timestamp := int64(1633036800000) // Example timestamp
	workerID := int64(1)
	sequence := int64(42)

	snowflakeID := NewSnowflake(timestamp, workerID, sequence)

	if snowflakeID.timestamp() != timestamp {
		t.Errorf("Expected timestamp %d, got %d", timestamp, snowflakeID.timestamp())
	}

	if snowflakeID.workerID() != workerID {
		t.Errorf("Expected workerID %d, got %d", workerID, snowflakeID.workerID())
	}

	if snowflakeID.sequence() != sequence {
		t.Errorf("Expected sequence %d, got %d", sequence, snowflakeID.sequence())
	}
}

func TestSnowflakeTime(t *testing.T) {
	timestamp := int64(1633036800000) // Example timestamp
	workerID := int64(1)
	sequence := int64(42)

	snowflakeID := NewSnowflake(timestamp, workerID, sequence)
	expectedTime := time.UnixMilli(timestamp + relativeEpoch)

	if !snowflakeID.Time().Equal(expectedTime) {
		t.Errorf("Expected time %v, got %v", expectedTime, snowflakeID.Time())
	}
}

func TestSnowflakeString(t *testing.T) {
	timestamp := int64(1633036800000) // Example timestamp
	workerID := int64(1)
	sequence := int64(42)

	snowflakeID := NewSnowflake(timestamp, workerID, sequence)
	expectedString := "0000000000000000000" // Replace with actual expected string

	if snowflakeID.String() == expectedString {
		t.Errorf("Expected string %s, got %s", expectedString, snowflakeID.String())
	}
}
