package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Worker struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Desc      interface{} `json:"desc"`
}

type Homeland struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Power   []int  `json:"power"`
}

func main() {
	p1 := Worker{}
	bs1 := `{"first_name": "David", "last_name": "Heart", ` +
		`"desc": {"country": "UK", "city": "London", "power": [1,4,6,7]}}`

	if err := json.Unmarshal([]byte(bs1), &p1); err != nil {
		log.Println(err)
	}
	fmt.Println(p1) // {David Heart map[city:London country:UK power:[1 4 6 7]]}

	var msg json.RawMessage
	p2 := Worker{Desc: &msg}

	if err := json.Unmarshal([]byte(bs1), &p2); err != nil {
		log.Println(err)
	}
	fmt.Println(p2) // {David Heart 0xc000116090}

	switch p2.FirstName {
	case "David":
		h := &Homeland{}
		fmt.Println("address of h", h)
		if err := json.Unmarshal(msg, &h); err != nil {
			log.Println(err)
		}
		descCountry, descCity, descPower := h.Country, h.City, h.Power
		fmt.Println(descCountry, descCity, descPower)

		p2.Desc = h
		fmt.Println(p2, p2.Desc) // {David Heart 0xc000124080} &{UK London [1 4 6 7]}
	default:
		log.Printf("unknown first name: %q", p1.FirstName)
	}
}
