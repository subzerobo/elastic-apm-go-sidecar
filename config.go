package main

import (
	"encoding/json"
	"github.com/urfave/cli"
	"os"
)

// Config for server
type Config struct {
	Listen     string `json:"listen"`
	Gelf       string `json:"gelf"`
	BufferSize int    `json:"bufferSize"`
	Quiet      bool   `json:"quiet"`
}

func (config *Config) FillFromContext(c *cli.Context) {
	config.Listen = c.String("listen")
	config.Gelf  = c.String("gelf")
	config.BufferSize = c.Int("size")
	config.Quiet = c.Bool("quiet")
	
	if c.String("config") != "" {
		// Now only support json config file
		err := config.parseJSONConfig(c.String("c"))
		checkFatalError(err)
	}
}

func (config *Config) parseJSONConfig(path string) error {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(config)
}
