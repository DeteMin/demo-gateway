package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cihub/seelog"
	"gopkg.in/yaml.v2"
)

// const for file name
const (
	configFile = "conf/server.yml"
	empty      = ""
)

// server config dto
type server struct {
	LogType            string `yaml:"logtype"`
	HandleTimeout      int    `yaml:"defaultTimeout"`
	GRPCServiceBind    string `yaml:"gRPCServiceBind"`
	HTTPServiceBind    string `yaml:"httpServiceBind"`
	HttpServiceGRPC    string `yaml:"httpServiceGrpc"`
	HTTPMetricsBind    string `yaml:"httpMetricsBind"`
	EnableHandlingTime bool   `yaml:"enableHandlingTime"`
}

// ServerConfig global config, init at startup, can use anywhere
var ServerConfig server

func initServer(file string) error {
	realFileName := configFile
	if file != empty {
		realFileName = file
	}
	var config server
	bytes, err := ioutil.ReadFile(realFileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return err
	}
	ServerConfig = config
	jsonStr, err := json.Marshal(ServerConfig)
	if err != nil {
		return err
	}
	seelog.Infof("server config %s", string(jsonStr))
	return nil
}

func init() {
	err := initServer(empty)
	if err != nil {
		seelog.Infof("init server config err %v.", err)
	}
}
