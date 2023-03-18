package apiserver

import "github.com/Wolfxxxz/Dictionary/store"

//General config for rest api
type Config struct {
	//Port for start api
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

//Should return default config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8081",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
