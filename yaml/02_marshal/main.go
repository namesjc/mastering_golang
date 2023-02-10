package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func main() {
	p := Person{"Jun", 30}
	y, err := yaml.Marshal(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(y))

	err = ioutil.WriteFile("file.yaml", y, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
