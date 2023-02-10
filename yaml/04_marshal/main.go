package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Company struct {
	ID      string             `yaml:"id"`
	Active  bool               `yaml:"active"`
	People  []string           `yaml:"people"`
	Budget  int64              `yaml:"budget"`
	Sectors map[string]float32 `yaml:"sectors"`
}

type CompanyGroup struct {
	Companies []Company `yaml:"companies"`
}

func main() {
	c1 := Company{
		"514", true, []string{"David", "Fred"}, 90000,
		map[string]float32{"Golang": 3, "Python": 4},
	}
	c2 := Company{
		"640", false, []string{"Alex", "Paul"}, 90000,
		map[string]float32{"C": 5, "Crystal": 2},
	}

	list := CompanyGroup{[]Company{c1, c2}}
	// list.Companies = append(list.Companies, *c1, *c2)

	y, err := yaml.Marshal(&list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(y))

	err = ioutil.WriteFile("file.yaml", y, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
