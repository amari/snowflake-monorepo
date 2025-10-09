package snowflake

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestNewSnowflakeService(t *testing.T) {
	options := SnowflakeServiceOptions{
		WorkerID: 1,
	}

	service, err := NewSnowflakeService(options)
	if err != nil {
		t.Fatalf("Failed to create SnowflakeService: %v", err)
	}

	if service.workerID != int64(options.WorkerID) {
		t.Errorf("Expected workerID %d, got %d", options.WorkerID, service.workerID)
	}
}

func TestSnowflakeServiceNextID(t *testing.T) {
	options := SnowflakeServiceOptions{
		WorkerID: 1,
	}
	service, _ := NewSnowflakeService(options)

	ctx := context.Background()
	id, err := service.NextID(ctx, false)
	if err != nil {
		t.Fatalf("Failed to generate next ID: %v", err)
	}

	if id.workerID() != service.workerID {
		t.Errorf("Expected workerID %d, got %d", service.workerID, id.workerID())
	}
}

func TestSnowflakeServiceBatchNextID(t *testing.T) {
	options := SnowflakeServiceOptions{
		WorkerID: 1,
	}
	service, _ := NewSnowflakeService(options)

	ctx := context.Background()
	batchSize := 5
	ids, err := service.BatchNextID(ctx, batchSize, false)
	if err != nil {
		t.Fatalf("Failed to generate batch IDs: %v", err)
	}

	if len(ids) != batchSize {
		t.Errorf("Expected batch size %d, got %d", batchSize, len(ids))
	}

	for i, id := range ids {
		if id.workerID() != service.workerID {
			t.Errorf("ID %d: Expected workerID %d, got %d", i, service.workerID, id.workerID())
		}
	}
}

func TestSnowflakeServiceClockBackwards(t *testing.T) {
	options := SnowflakeServiceOptions{
		WorkerID: 1,
	}
	service, _ := NewSnowflakeService(options)

	service.lastTimestamp = time.Now().UnixMilli() - relativeEpoch + 1

	ctx := context.Background()
	_, err := service.NextID(ctx, false)
	if !errors.Is(err, ErrClockBackwards) {
		t.Errorf("Expected ErrClockBackwards, got %v", err)
	}
}
