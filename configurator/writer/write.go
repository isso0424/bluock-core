package configurator

import (
	"blumaton/bluock-core/config"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

var (
	_, b, _, _ = runtime.Caller(0)
	configPath = path.Join(filepath.Dir(b), "config.yml")
)

func WriteConfig(c config.Config) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
