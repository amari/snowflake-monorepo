package config

import "errors"

type SnowflakeConfig struct {
	// MachineID is the unique identifier for the worker/machine generating IDs.
	MachineID int64 `koanf:"machineID"`
}

func (sc *SnowflakeConfig) Validate() error {
	if sc.MachineID < 0 || sc.MachineID > 1023 {
		return errors.New("worker id must be between 0 and 1023 (inclusive)")
	}

	return nil
}
