package app

import "os"

type Config struct {
	HTTPAddr string
}

func NewConfigFromEnv() Config {
	config := Config{
		HTTPAddr: ":8080",
	}
	if httpAddr, ok := os.LookupEnv("HTTP_ADDR"); ok {
		config.HTTPAddr = httpAddr
	}

	return config
}
