package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config holds values from the local yaml
// TODO::Config store interface for modularity
type Config struct {
	env  string            `yaml:"environment"`
	name string            `yaml:"name"`
	kv   map[string]string `yaml:"kv,omitempty"`
}

// BuildConfigFromFile returns the config from a path
func (c *Config) BuildConfigFromFile(path string) *Config {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading config %v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// GetValueMap returns the loaded kv map from config
func (c *Config) GetValueMap() map[string]string {
	return c.kv
}
