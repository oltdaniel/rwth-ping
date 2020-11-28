package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Telegram struct {
		Token        string              `yaml:"token"`
		TargetGroups map[string][]string `yaml:"targetGroups"`
	} `yaml:"telegram"`
	Tasks map[string]struct {
		Enabled  bool `yaml:"enabled"`
		Interval uint `yaml:"interval"`
	} `yaml:"tasks"`
}

func (c *Config) GetInterval(slug string, d time.Duration) time.Duration {
	if t, ok := c.Tasks[slug]; ok {
		if t.Interval > 0 {
			return time.Duration(t.Interval) * time.Second
		}
	}
	return d
}

func (c *Config) IsEnabled(slug string) bool {
	if t, ok := c.Tasks[slug]; ok {
		return t.Enabled
	}
	return true
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
	fmt.Println(config)
	return config
}()
