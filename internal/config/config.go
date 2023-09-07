package config

import "os"

type Config struct {
	HTTPAddr string
}

func Read() Config {
	config := Config{
		HTTPAddr: ":8888",
	}
	httpAddr, ok := os.LookupEnv("HTTP_ADDR")
	if ok {
		config.HTTPAddr = httpAddr
	}
	return config
}
