package utils

import "os"

type Environent struct {
	PORT string
	DSN  string
}

func LoadEnv() *Environent {
	return &Environent{
		PORT: os.Getenv("SERVICE_PORT"),
		DSN:  os.Getenv("DSN"),
	}
}
