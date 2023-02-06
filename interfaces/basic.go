package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (p PhoneConnector) Name() string {
	return p.name
}

func (p PhoneConnector) Connect() {
	fmt.Println("Connected:", p.name)
}

type TVConnector struct {
	name string
}

func (tv TVConnector) Connect() {
	fmt.Println("Connected:", tv.name)
}

func main() {
	pc := PhoneConnector{
		name: "PhoneConnector",
	}
	var a Connector
	a = Connector(pc)
	a.Connect()

	// interface convert failed
	tv := TVConnector{
		name: "TVConnector",
	}
	var b USB
	b = USB(tv)
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device.")
	}

}
