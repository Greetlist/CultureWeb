package config

import (
    "io/ioutil"
    yaml "gopkg.in/yaml.v3"
    "fmt"
    "os"
)

type TLS struct {
    ServerCrt string `yaml:"server_crt"`
    ServerKey string `yaml:"server_key"`
}

type Mysql struct {
    User string `yaml:"user"`
    Password string `yaml:"password"`
    Host string `yaml:"host"`
    Port int `yaml:"port"`
    DBName string `yaml:"db_name"`
    Params string `yaml:"params"`
    MaxIdleConns int `yaml:"max_idle_conns"`
    MaxOpenConns int `yaml:"max_open_conns"`
}

type YamlConfig struct {
    BindAddr string `yaml:"bind_addr"`
    BindPort int64 `yaml:"bind_port"`
    LogDir string `yaml:"log_dir"`
    TLSConfig TLS `yaml:"tls_config"`
    MysqlConfig Mysql `yaml:"mysql_config"`
}

var GlobalConfig YamlConfig

func InitConfig(config_file string) {
    data, err := ioutil.ReadFile(config_file)
    if err != nil {
        fmt.Printf("Read File Error: [%v]", err)
        os.Exit(-1)
    }
    err = yaml.Unmarshal([]byte(data), &GlobalConfig)
    if err != nil {
        fmt.Printf("Yaml Unmarshal Error: [%v]", err)
        os.Exit(-1)
    }
}
