package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

const input = `
{
	"type": "processor",
	"desc": "AMD 2950x",
	"cores": 16
}
`

type MachinePart struct {
	Type string `json:"type"`
}

type Processor struct {
	Description string `json:"desc"`
	Cores       int    `json:"cores"`
}

func main() {
	var mp MachinePart
	buf := []byte(input)
	if err := json.Unmarshal(buf, &mp); err != nil {
		log.Println(err)
	}
	fmt.Println(mp) // {processor}
	switch mp.Type {
	case "processor":
		var s struct {
			MachinePart
			Processor
		}
		if err := json.Unmarshal(buf, &s); err != nil {
			log.Println(err)
		}
		json_type, desc, cores := s.Type, s.Description, s.Cores
		fmt.Println(json_type + " > " + desc + " > " + strconv.Itoa(cores))
	default:
		log.Printf("unknown message type: %q", mp.Type)
	}
}
