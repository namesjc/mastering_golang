package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Engineer struct {
	ID         string             `yaml:"id"`
	Sector     int                `yaml:"sector"`
	Tasks      []string           `yaml:"tasks"`
	DailyHours []int              `yaml:"dailyHours"`
	Languages  map[string]float32 `yaml:"languages"`
}

func main() {
	e := Engineer{}
	source, err := ioutil.ReadFile("file.yaml")
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(source, &e)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Println(e)
}
