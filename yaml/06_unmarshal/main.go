package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type DirAndFiles struct {
	DataSpace   string `yaml:"dataSpace"`
	ConfigSpace string `yaml:"configSpace"`
}

type InstanceConfig struct {
	Port          string   `yaml:"port"`
	Host          string   `yaml:"host"`
	BindIP        string   `yaml:"bindIp"`
	CustomOptions []string `yaml:"customOptions"`
}

type Instance struct {
	Installation bool           `yaml:"install"`
	InstanceType string         `yaml:"instanceType"`
	DirAndFiles  DirAndFiles    `yaml:"dirAndFiles"`
	ConfigFile   InstanceConfig `yaml:"configFile"`
}

func ReadInstanceYaml(serversDataFile string, obj *Instance) {
	source1, err := ioutil.ReadFile(serversDataFile)
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}
	err = yaml.Unmarshal(source1, obj)
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}
}

func main() {
	e := Instance{}
	ReadInstanceYaml("file.yaml", &e)
	fmt.Println(e)
}
