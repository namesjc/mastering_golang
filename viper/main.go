package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	vi := viper.New()
	vi.SetConfigFile("test.yaml")
	vi.ReadInConfig()
	fmt.Println(vi.GetInt("port"))
	fmt.Println(vi.GetString("hostname"))
}
