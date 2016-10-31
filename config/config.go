package config

import (
	"encoding/json"
	"github.com/DmitryBugrov/log"
	"os"
)

//Config -  struct for config file
type Config struct {
	FileName string
	Port     int
	LogLevel int
}

//Init - set variable config file name
func (c *Config) Init(f string) {
	c.FileName = f
}

//Load - load config file
func (c *Config) Load(l log.Log) error {
	l.Print(log.LogLevelTrace, "Enter to Config.Load")
	file, err := os.Open(c.FileName)
	if err != nil {
		l.Print(log.LogLevelError, "Configuration file cannot be loaded: ", c.FileName)
		return err
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&c)
	if err != nil {
		l.Print(log.LogLevelError, "Unable to decode config into struct", err.Error())
	}

	return nil
}
