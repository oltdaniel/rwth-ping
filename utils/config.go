package utils

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Telegram struct {
		Token        string              `yaml:"token"`
		TargetGroups map[string][]string `yaml:"targetGroups"`
	} `yaml:"telegram"`
}

var CONFIG = func() Config {
	fp, err := os.Open("./config.yml")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	var config Config
	yaml.Unmarshal(data, &config)
	return config
}()
