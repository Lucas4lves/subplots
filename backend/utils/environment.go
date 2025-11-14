package utils

import "os"

type Environent struct {
	PORT string
}

func LoadEnv() *Environent {
	return &Environent{
		PORT: os.Getenv("SERVICE_PORT"),
	}
}
