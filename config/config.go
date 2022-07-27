package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ProxyServerAddr string      `json:"proxy_server_addr"`
	IoTsDevices     []IotConfig `json:"iots_devices"`
}

type IotConfig struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	BaseNumber int    `json:"base_number"`
	Step       int    `json:"step"`

	Formulas Formulas `json:"formulas"`
}

type Formulas struct {
	Formula1 string `json:"formula_1,omitempty"`
	Formula2 string `json:"formula_2,omitempty"`
	Formula3 string `json:"formula_3,omitempty"`
	Formula4 string `json:"formula_4,omitempty"`
	Formula5 string `json:"formula_5,omitempty"`
}

var configPath = "conf.config"

func LoadConfig() (loadedConf Config, err error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("loadconfig err:%v", err)
	}

	err = json.Unmarshal(data, &loadedConf)
	return
}
