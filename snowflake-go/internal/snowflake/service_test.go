package snowflake

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestNewSnowflakeService(t *testing.T) {
	options := SnowflakeServiceOptions{
		MachineID: 1,
	}

	service, err := NewSnowflakeService(options)
	if err != nil {
		t.Fatalf("Failed to create SnowflakeService: %v", err)
	}

	if service.machineID != int64(options.MachineID) {
		t.Errorf("Expected workerID %d, got %d", options.MachineID, service.machineID)
	}
}

func TestSnowflakeServiceNextID(t *testing.T) {
	options := SnowflakeServiceOptions{
		MachineID: 1,
	}
	service, _ := NewSnowflakeService(options)

	ctx := context.Background()
	id, err := service.NextID(ctx, false)
	if err != nil {
		t.Fatalf("Failed to generate next ID: %v", err)
	}

	if id.workerID() != service.machineID {
		t.Errorf("Expected workerID %d, got %d", service.machineID, id.workerID())
	}
}

func TestSnowflakeServiceBatchNextID(t *testing.T) {
	options := SnowflakeServiceOptions{
		MachineID: 1,
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
		if id.workerID() != service.machineID {
			t.Errorf("ID %d: Expected workerID %d, got %d", i, service.machineID, id.workerID())
		}
	}
}

func TestSnowflakeServiceClockBackwards(t *testing.T) {
	options := SnowflakeServiceOptions{
		MachineID: 1,
	}
	service, _ := NewSnowflakeService(options)

	service.lastTimestamp = time.Now().UnixMilli() - relativeEpoch + 1

	ctx := context.Background()
	_, err := service.NextID(ctx, false)
	if !errors.Is(err, ErrClockBackwards) {
		t.Errorf("Expected ErrClockBackwards, got %v", err)
	}
}
