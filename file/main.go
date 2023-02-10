// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type ExampleValue struct {
	Name   string
	Values string
}

type YAMLFile struct {
	ExampleValues []ExampleValue `yaml:"examples"`
}

func main() {

	// var findings []string

	// lineBreakRegExp := regexp.MustCompile(`\r?\n`)

	// fileContents, err := os.ReadFile("/home/cje5szh/learning/mastering_golang/file/file.tpl")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fileLines := lineBreakRegExp.Split(string(fileContents), -1)

	// includeFound := false

	// for _, line := range fileLines {
	// 	// fmt.Println(line)
	// 	if strings.Contains(line, "/}}") {
	// 		includeFound = false
	// 	}

	// 	if includeFound {
	// 		findings = append(findings, line)
	// 	}

	// 	if strings.Contains(line, "includeStatement") {
	// 		includeFound = true
	// 	}

	// }

	// fmt.Println(strings.Join(findings, "\n"))

	allExampleValueFiles := ReadFileExtractPatternFindingsBetweenMarkers("/home/cje5szh/learning/mastering_golang/file/file.tpl", "exampleValues", "/}}")

	fmt.Println("value of allExampleValueFiles:\n", allExampleValueFiles)

	// type ExampleValue struct {
	// 	Name   string
	// 	Values string
	// }

	// type YAMLFile struct {
	// 	ExampleValues []ExampleValue `yaml:"examples"`
	// }

	var decodedExampleValues YAMLFile
	var exampleValues = make(map[string]string)
	var exampleTitles []string

	err := yaml.Unmarshal([]byte(allExampleValueFiles), &decodedExampleValues)
	if err != nil {
		log.Fatalf("unmarshal failed with '%s'\n", err)
	}

	// dec := yaml.NewDecoder(strings.NewReader(allExampleValueFiles))
	// err := dec.Decode(&decodedExampleValues)
	// if err != nil {
	// 	log.Fatalf("dec.Decode() failed with '%s'\n", err)
	// }

	fmt.Println("value of decodedExampleValues:\n", decodedExampleValues)
	// get all examples to slice
	for _, examples := range decodedExampleValues.ExampleValues {
		exampleTitles = append(exampleTitles, examples.Name)
		exampleValues[examples.Name] = examples.Values
	}
	fmt.Println("value of exampleTitles:\n", exampleTitles)
	fmt.Println("value of exampleValues:\n", exampleValues)
}

func ReadFileExtractPatternFindingsBetweenMarkers(filePath, markerStart, markerEnd string) string {

	var findings []string
	lineBreakRegExp := regexp.MustCompile(`\r?\n`)

	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	fileLines := lineBreakRegExp.Split(string(fileContents), -1)

	includeFound := false

	for _, line := range fileLines {

		if strings.Contains(line, markerEnd) {
			includeFound = false
		}

		if includeFound {
			findings = append(findings, line)
		}

		if strings.Contains(line, markerStart) {
			includeFound = true
		}

	}

	return strings.Join(findings, "\n")
}
