package config

import (
	"os"
	"encoding/json"
	"fmt"
	"path/filepath"
)

type ServerCache struct {
	Host string
	Port int
}

type GrayLog struct {
	Host string
	Port int
}

type Config struct {
	Cpu    int
	Web string
	Rpc string
	Ns string
	Set string
	Cache []ServerCache
	Logger GrayLog

}

// загрузка конфига
func New() *Config{
	c := new(Config)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	file, _ := os.Open(dir + "/conf.json")
	fmt.Println("Config loaded from: " + dir + "/conf.json")
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		fmt.Println("Configure failed: ", err)
	}

	return c
}