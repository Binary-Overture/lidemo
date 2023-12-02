package config

import (
	"github.com/spf13/viper"
	"os"
)

var Config *Conf

type Conf struct {
	System *System           `yaml:"system"`
	Mysql  map[string]*Mysql `yaml:"mysql"`
	Redis  *Redis            `yaml:"redis"`
}

type System struct {
	AppEnv   string `yaml:"appEnv"`
	Version  string `yaml:"version"`
	Domain   string `yaml:"domain"`
	HttpPort string `yaml:"httpPort"`
	Host     string `yaml:"host"`
}

type Mysql struct {
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Redis struct {
	Host     string `yaml:"redisHost"`
	Port     string `yaml:"redisPort"`
	Password string `yaml:"redisPassword"`
	DbName   int    `yaml:"redisDbName"`
	NetWork  string `yaml:"redisNetWork"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/local")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
