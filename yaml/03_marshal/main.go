package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Developer struct {
	ID         string             `yaml:"id"`
	Sector     int                `yaml:"sector"`
	Salary     float32            `yaml:"salary"`
	Tasks      []string           `yaml:"tasks"`
	DailyHours []int              `yaml:"dailyHours"`
	Languages  map[string]float32 `yaml:"languages"`
}

func main() {
	d := Developer{
		"51436256",
		3,
		9000.00,
		[]string{"feature/task312", "feture/task314"},
		[]int{7, 9, 6, 8, 7},
		map[string]float32{"Golang": 9.1, "Python": 7.9},
	}
	y, err := yaml.Marshal(&d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(y))

	err = ioutil.WriteFile("file.yaml", y, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
