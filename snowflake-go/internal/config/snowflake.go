package config

import "errors"

type SnowflakeConfig struct {
	// WorkerID is the unique identifier for the worker/machine generating IDs.
	WorkerID int64 `koanf:"workerID"`
}

func (sc *SnowflakeConfig) Validate() error {
	if sc.WorkerID < 0 || sc.WorkerID > 1023 {
		return errors.New("worker id must be between 0 and 1023 (inclusive)")
	}

	return nil
}
