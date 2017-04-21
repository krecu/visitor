package config

import (
	"os"
	"encoding/json"
	"fmt"
)

type ServerCache struct {
	Host string
	Port int
}

type Config struct {
	Cpu    int
	Listen string
	Ns string
	Set string
	Cache []ServerCache
}

// загрузка конфига
func New() *Config{
	c := new(Config)
	file, _ := os.Open("conf.json")
	err := json.NewDecoder(file).Decode(&c)
	if err != nil {
		fmt.Println("Configure failed: ", err)
	}

	return c
}