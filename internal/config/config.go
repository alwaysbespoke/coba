package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	configs map[string]interface{}
}

// New returns a new Config instance
func New() *Config {
	return &Config{
		configs: make(map[string]interface{}),
	}
}

// Run processes the configs
func (c *Config) Run() {
	for key, config := range c.configs {
		if err := envconfig.Process(key, config); err != nil {
			panic(fmt.Errorf("failed to process config: %w", err))
		}
	}
}

// Set sets a new config
func (c *Config) Set(key string, config interface{}) {
	c.configs[key] = config
}

// Get returns a config by the key
func (c *Config) Get(key string) interface{} {
	return c.configs[key]
}
