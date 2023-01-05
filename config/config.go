package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// TODO:Ref go-micro

// Endpoint
type Endpoint struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Config
type Config struct {
	Backend struct {
		Endpoints []Endpoint `yaml:"endpoints"`
	} `yaml:"backend"`
}

func (c *Config) Init(filepath string) {
	// Read the YAML file
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	// Unmarshal the data into a Config struct
	err = yaml.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
}
func Init(c *Config, filepath string) {
	c.Init(filepath)
}
