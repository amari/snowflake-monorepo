package snowflake

import (
	"context"
	"errors"
	"time"
)

type SnowflakeServiceOptions struct {
	// MachineID is a unique identifier for the worker/machine (0-1023).
	MachineID uint16
}

type SnowflakeService struct {
	machineID     int64
	sequence      int64
	lastTimestamp int64

	mutexCh chan struct{}
}

func NewSnowflakeService(options SnowflakeServiceOptions) (*SnowflakeService, error) {
	if options.MachineID > 1023 {
		return nil, ErrInvalidWorkerID
	}

	return &SnowflakeService{
		machineID: int64(options.MachineID),
		mutexCh:   make(chan struct{}, 1), // buffered, acts like a mutex
	}, nil
}

func (s *SnowflakeService) nextID() (Snowflake, error) {
	// This method assumes the caller has already acquired the mutex lock.

	timestamp := time.Now().UnixMilli() - relativeEpoch

	if timestamp < s.lastTimestamp {
		// Clock moved backwards, refuse to generate ID
		return 0, ErrClockBackwards
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & 0xFFF // 12 bits for sequence
		if s.sequence == 0 {
			// Sequence overflow, return error
			return 0, ErrSequenceOverflow
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	return NewSnowflake(timestamp, s.machineID, s.sequence), nil
}

func (s *SnowflakeService) NextID(ctx context.Context, wait bool) (Snowflake, error) {
	select {
	case s.mutexCh <- struct{}{}: // try lock
		defer func() {
			select {
			case <-s.mutexCh: // unlock without blocking
			default:
			}
		}()
	case <-ctx.Done():
		return 0, ctx.Err()
	}

	// Retry loop for acquiring lock and generating ID
	for {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			id, err := s.nextID()

			if err != nil {
				if errors.Is(err, ErrSequenceOverflow) && wait {
					// If sequence overflowed and wait is true, retry after a short delay
					select {
					case <-time.After(500 * time.Microsecond):
						continue
					case <-ctx.Done():
						return 0, ctx.Err()
					}
				}

				return 0, err
			}

			return id, nil
		}
	}
}

func (s *SnowflakeService) batchNextID(n int) ([]Snowflake, error) {
	// This method assumes the caller has already acquired the mutex lock.

	timestamp := time.Now().UnixMilli() - relativeEpoch
	if timestamp < s.lastTimestamp {
		// Clock moved backwards, refuse to generate ID
		return nil, ErrClockBackwards
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & 0xFFF // 12 bits for sequence
		if s.sequence == 0 {
			// Sequence overflow, return error
			return nil, ErrSequenceOverflow
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	maxN := int(4096 - s.sequence)
	if n > maxN {
		n = maxN // cap n to prevent sequence overflow
	}

	ids := make([]Snowflake, n)

	for i := range n {
		ids[i] = NewSnowflake(timestamp, s.machineID, s.sequence+int64(i))
	}
	s.sequence += int64(n - 1) // update sequence to last used

	return ids, nil
}

func (s *SnowflakeService) BatchNextID(ctx context.Context, n int, wait bool) ([]Snowflake, error) {
	if n <= 0 {
		return nil, nil
	}

	if n > 32 {
		n = 32 // cap at 32 to prevent excessive load
	}

	select {
	case s.mutexCh <- struct{}{}: // try lock
		defer func() {
			select {
			case <-s.mutexCh: // unlock without blocking
			default:
			}
		}()
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	ret := make([]Snowflake, 0, n)

	// Retry loop for acquiring lock and generating ID
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			ids, err := s.batchNextID(n)

			if err != nil {
				if errors.Is(err, ErrSequenceOverflow) && wait {
					// If sequence overflowed and wait is true, retry after a short delay
					select {
					case <-time.After(500 * time.Microsecond):
						continue
					case <-ctx.Done():
						return nil, ctx.Err()
					}
				}

				return nil, err
			}

			n -= len(ids)
			ret = append(ret, ids...)

			return ret, nil
		}
	}
}
