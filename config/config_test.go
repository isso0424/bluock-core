package config_test

import (
	"blumaton/bluock-core/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	c, err := config.LoadConfig("example.yml")
	if err != nil {
		t.Fatal(err)

		return
	}

	assert.Equal(t, 100, c.RSSIThreshold)
	assert.Equal(t, "hoge", c.MACAddresses[0])
	assert.Equal(t, "fuga", c.MACAddresses[1])
}
