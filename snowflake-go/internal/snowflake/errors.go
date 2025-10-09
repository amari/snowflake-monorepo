package snowflake

import "errors"

var (
	ErrInvalidWorkerID  = errors.New("worker ID must be between 0 and 1023 (inclusive)")
	ErrClockBackwards   = errors.New("clock moved backwards, refusing to generate id")
	ErrSequenceOverflow = errors.New("sequence overflow, too many IDs generated in the same millisecond")
)
