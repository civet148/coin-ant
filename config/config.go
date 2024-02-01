package config

import (
	"coin-ant/types"
	"encoding/json"
	"github.com/civet148/log"
	"io/ioutil"
	"os"
	"path"
)

const (
	strConfName   = "config.json"
	RunConfigName = "system-backend"
)

type Config struct {
	Version string `json:"version"`
	DSN     string `json:"dsn" toml:"dsn" db:"dsn" cli:"dsn"`
	Debug   bool   `json:"debug" toml:"debug" db:"debug" cli:"debug"`
}

var strFilePath = path.Join(types.DefaultConfigHome, strConfName)

func (c *Config) Save() (err error) {
	var data []byte
	strConfigPath := types.DefaultConfigHome
	if err = os.MkdirAll(strConfigPath, os.ModePerm); err != nil {
		return log.Errorf("make dir [%s] error [%s]", strConfigPath, err.Error())
	}
	data, err = json.Marshal(c)
	if err != nil {
		return log.Errorf("marshal json error: %s", err.Error())
	}
	err = ioutil.WriteFile(strFilePath, data, os.ModePerm)
	if err != nil {
		return log.Errorf("write to file %s error: %s", strFilePath, err.Error())
	}
	return nil
}

func (c *Config) Load() (err error) {
	var data []byte
	data, err = ioutil.ReadFile(strFilePath)
	if err != nil {
		return log.Errorf("read file %s error: %s", strFilePath, err.Error())
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		return log.Errorf("unmarshal error: %s", err.Error())
	}
	return nil
}
