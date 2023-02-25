package config

import (
    "io/ioutil"
    yaml "gopkg.in/yaml.v3"
    "fmt"
)

type YamlConfig struct {
    BindAddr string `yaml:"bind_addr"`
    BindPort int64 `yaml:"bind_port"`
    LogDir string `yaml:"log_dir"`
}

var GlobalConfig YamlConfig

func InitConfig(config_file string) {
    data, err := ioutil.ReadFile(config_file)
    if err != nil {
        fmt.Printf("Read File Error: [%v]", err)
    }
    err = yaml.Unmarshal([]byte(data), &GlobalConfig)
    if err != nil {
        fmt.Printf("Yaml Unmarshal Error: [%v]", err)
    }
}
