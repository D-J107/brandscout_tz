package config

import (
	"os"
	"quotes_api/internal/logging"
	"strconv"
	"strings"
)

type Config struct {
	RestPort        string
	ShutdownTimeout int
}

func MustLoad() *Config {
	shutdownTimeout, err := strconv.Atoi(getVariableByEnv("SHUTDOWN_TIMEOUT"))
	if err != nil {
		panic("shutdown timeout variable must be integer")
	}
	conf := &Config{
		RestPort:        getVariableByEnv("REST_PORT"),
		ShutdownTimeout: shutdownTimeout,
	}
	// если забыли указать :
	if !strings.HasPrefix(conf.RestPort, ":") {
		conf.RestPort = ":" + conf.RestPort
	}
	logging.Logger.Info("brandscout_tt app successfully configured")
	return conf
}

func getVariableByEnv(key string) string {
	if variable := os.Getenv(key); variable != "" {
		return variable
	}
	panic("key " + key + " not provided by environment!")
}
