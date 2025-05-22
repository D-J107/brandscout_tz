package config

import (
	"fmt"
	"os"
)

type Config struct {
	RestPort string
}

func MustLoad() *Config {
	conf := &Config{
		RestPort: getVariableByEnv("REST_PORT"),
	}
	fmt.Println("brandscout_tt app successfully configured")
	return conf
}

func getVariableByEnv(key string) string {
	if variable := os.Getenv(key); variable != "" {
		return variable
	}
	panic("key " + key + " not provided by environment!")
}
