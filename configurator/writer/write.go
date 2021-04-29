package writer

import (
	"blumaton/bluock-core/config"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var (
	b, _ = os.Getwd()
	configPath = path.Join(b, "config.yml")
)

func WriteConfig(c config.Config) error {
	fmt.Println(b)
	fmt.Println(configPath)
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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
