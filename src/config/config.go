package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config holds values from the local yaml
// TODO::Config store interface for modularity
type Config struct {
	Env  string            `yaml:"environment"`
	Name string            `yaml:"name"`
	KV   map[string]string `yaml:"kv,omitempty"`
}

// BuildConfigFromFile returns the config from a path
func (c *Config) BuildConfigFromFile(path string) *Config {
	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading config %v ", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal failure: %v", err)
	}

	return &config
}

// GetValueMap returns the loaded kv map from config
func (c *Config) GetValueMap() map[string]string {
	return c.KV
}
