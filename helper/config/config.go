package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	DbUsername string `yaml:"DB_USERNAME"`
	DbPassword string `yaml:"DB_PASSWORD"`
	DbHost     string `yaml:"DB_HOST"`
	DbName     string `yaml:"DB_NAME"`
	Port       string `yaml:"PORT"`
}

var Config Configs

func InitConfig(path string) {
	// Check environment (todo)

	envConf := "master"
	fileName := path + "/" + envConf + ".yaml"

	log.Printf("Intiating Config from %s", fileName)
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("read file error,err=%v", err)
		panic(err)
	}
	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		log.Printf("unmarshal config error,err=%v", err)
		panic(err)
	}
}
