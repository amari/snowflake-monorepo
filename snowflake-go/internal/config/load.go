package config

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/cliflagv3"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/urfave/cli/v3"
)

func LoadConfigFile(k *koanf.Koanf, path string) error {
	return k.Load(file.Provider(path), yaml.Parser())
}

func LoadConfigDir(k *koanf.Koanf, path string) error {
	var errs []error

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for _, entry := range entries {
		if err := LoadConfigFile(k, filepath.Join(path, entry.Name())); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func LoadConfigEnvVars(k *koanf.Koanf, prefix string, delim string) error {
	return k.Load(env.ProviderWithValue(prefix, delim, func(k string, v string) (string, any) {
		key := strings.ReplaceAll(
			strings.ToLower(strings.TrimPrefix(k, prefix)),
			"_", delim,
		)
		if strings.Contains(v, ",") {
			return key, strings.Split(v, ",")
		}

		return key, v
	}), nil)
}

var wellKnownCliFlagNameMap map[string]string = map[string]string{
	"config-dir": "configDir",
	"root-dir":   "rootDir",
}

func LoadConfigCliFlags(k *koanf.Koanf, f *cli.Command, delim string) error {
	if err := k.Load(cliflagv3.Provider(f, "."), nil); err != nil {
		return err
	}

	prefix := f.Name + "."

	// Strip "{f.Name}." prefix from all CLI keys
	for key, val := range k.All() {
		if after, ok := strings.CutPrefix(key, prefix); ok {
			newKey := after
			k.Set(newKey, val)
			k.Delete(key)
		}
	}

	// Remap kebab-case to camelCase
	for path, key := range wellKnownCliFlagNameMap {
		if k.Exists(path) {
			k.Set(key, k.Get(path))
		}
	}

	return nil
}
